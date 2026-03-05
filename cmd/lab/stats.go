package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type statsComparison struct {
	Track            string  `json:"track"`
	Mode             string  `json:"mode"`
	Workload         string  `json:"workload"`
	Class            string  `json:"class"`
	ImplA            string  `json:"impl_a"`
	ImplB            string  `json:"impl_b"`
	NImplA           int     `json:"n_impl_a"`
	NImplB           int     `json:"n_impl_b"`
	MeanMSImplA      float64 `json:"mean_ms_impl_a"`
	MeanMSImplB      float64 `json:"mean_ms_impl_b"`
	MedianMSImplA    float64 `json:"median_ms_impl_a"`
	MedianMSImplB    float64 `json:"median_ms_impl_b"`
	CVImplA          float64 `json:"cv_impl_a"`
	CVImplB          float64 `json:"cv_impl_b"`
	Winner           string  `json:"winner"`
	Speedup          float64 `json:"speedup"`
	CI95Low          float64 `json:"ci95_low"`
	CI95High         float64 `json:"ci95_high"`
	PValue           float64 `json:"p_value"`
	PValueAdjusted   float64 `json:"p_value_adjusted"`
	EffectSizeCohenD float64 `json:"effect_size_cohen_d"`
	Significant      bool    `json:"significant"`
	SignificantBH    bool    `json:"significant_bh"`
	NoiseFloorMS          float64 `json:"noise_floor_ms"`
	MinObservableEffPct   float64 `json:"min_observable_effect_pct"`
	OraclePassRateA       float64 `json:"oracle_pass_rate_a"`
	OraclePassRateB       float64 `json:"oracle_pass_rate_b"`
}

type statsReport struct {
	GeneratedAtUTC string            `json:"generated_at_utc"`
	Alpha          float64           `json:"alpha"`
	Resamples      int               `json:"resamples"`
	Comparisons    []statsComparison `json:"comparisons"`
}

func runStats(runsPath, apiBenchPath string, alpha float64, resamples int) error {
	if alpha <= 0 || alpha >= 1 {
		return errors.New("alpha must be in (0,1)")
	}
	if resamples < 200 {
		return errors.New("resamples must be >= 200")
	}
	root, err := repoRoot()
	if err != nil {
		return err
	}
	runsPath = choosePath(runsPath, filepath.Join(root, "results", "latest-cli-runs.csv"))
	apiBenchPath = choosePath(apiBenchPath, filepath.Join(root, "results", "latest-api-bench.txt"))

	runs, err := loadRunsCSV(runsPath)
	if err != nil {
		return err
	}

	comparisons := buildStatsComparisons(runs, alpha, resamples)

	// Parse API bench output and add API comparisons.
	apiRuns, apiErr := loadAPIBenchRuns(apiBenchPath)
	if apiErr == nil && len(apiRuns) > 0 {
		apiComparisons := buildStatsComparisons(apiRuns, alpha, resamples)
		comparisons = append(comparisons, apiComparisons...)
		fmt.Printf("included %d API comparisons from %s\n", len(apiComparisons), apiBenchPath)
	} else if apiErr != nil {
		fmt.Printf("warning: skipping API bench stats (%v)\n", apiErr)
	}

	applyBenjaminiHochberg(comparisons, alpha)

	report := statsReport{
		GeneratedAtUTC: time.Now().UTC().Format(time.RFC3339),
		Alpha:          alpha,
		Resamples:      resamples,
		Comparisons:    comparisons,
	}

	stamp := time.Now().UTC().Format("20060102T150405Z")
	jsonPath := filepath.Join(root, "results", "stats-"+stamp+".json")
	mdPath := filepath.Join(root, "results", "stats-"+stamp+".md")
	latestJSON := filepath.Join(root, "results", "latest-stats.json")
	latestMD := filepath.Join(root, "results", "latest-stats.md")

	b, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(jsonPath, b, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latestJSON, b, 0o644); err != nil {
		return err
	}

	if err := os.WriteFile(mdPath, []byte(renderStatsMarkdown(report, runsPath)), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latestMD, []byte(renderStatsMarkdown(report, runsPath)), 0o644); err != nil {
		return err
	}

	fmt.Printf("stats complete\n- %s\n- %s\n", jsonPath, mdPath)
	fmt.Printf("latest links\n- %s\n- %s\n", latestJSON, latestMD)
	return nil
}

// loadAPIBenchRuns parses Go testing.B benchmark output into runRecord entries
// suitable for the same statistical comparison pipeline used by CLI benchmarks.
// Each BenchmarkAPI line becomes one run with track="api".
func loadAPIBenchRuns(path string) ([]runRecord, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var runs []runRecord
	sessionCounters := map[string]int{} // track session IDs per impl+workload

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}
		// Parse: BenchmarkAPICanonicalizeSchubfach/workload-GOMAXPROCS \t iters \t ns/op ...
		// or:   BenchmarkAPICanonicalizeJSONCanon/workload-GOMAXPROCS ...
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		name := fields[0]
		// Find ns/op value
		var nsOp float64
		for i, f := range fields {
			if f == "ns/op" && i > 0 {
				nsOp, _ = strconv.ParseFloat(fields[i-1], 64)
				break
			}
		}
		if nsOp == 0 {
			continue
		}

		// Extract impl and mode/workload from benchmark name.
		var impl, mode, workload string
		slashIdx := strings.Index(name, "/")
		if slashIdx < 0 {
			continue
		}
		prefix := name[:slashIdx]
		rest := name[slashIdx+1:]

		// Remove -GOMAXPROCS suffix (e.g., "-20").
		if dashIdx := strings.LastIndex(rest, "-"); dashIdx > 0 {
			if _, err := strconv.Atoi(rest[dashIdx+1:]); err == nil {
				rest = rest[:dashIdx]
			}
		}

		switch {
		case strings.Contains(prefix, "CanonicalizeSchubfach"):
			impl = "schubfach"
			mode = "canonicalize"
		case strings.Contains(prefix, "CanonicalizeJSONCanon"):
			impl = "json-canon"
			mode = "canonicalize"
		case strings.Contains(prefix, "VerifySchubfach"):
			impl = "schubfach"
			mode = "verify"
		case strings.Contains(prefix, "VerifyJSONCanon"):
			impl = "json-canon"
			mode = "verify"
		default:
			continue
		}
		workload = rest

		key := impl + "|" + mode + "|" + workload
		session := sessionCounters[key]
		sessionCounters[key]++

		runs = append(runs, runRecord{
			SessionID:    session,
			Track:        "api",
			Impl:         impl,
			Mode:         mode,
			Workload:     workload,
			Class:        "valid",
			OK:           true,
			PassesOracle: true,
			DurationNS:   int64(nsOp),
		})
	}
	return runs, nil
}


func loadRunsCSV(path string) ([]runRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, errors.New("runs csv has no records")
	}
	h := indexHeader(rows[0])
	required := []string{"session_id", "track", "impl", "mode", "workload", "class", "ok", "passes_oracle", "duration_ns"}
	for _, k := range required {
		if _, ok := h[k]; !ok {
			return nil, fmt.Errorf("missing %s in runs header", k)
		}
	}

	out := make([]runRecord, 0, len(rows)-1)
	for _, row := range rows[1:] {
		ok, _ := strconv.ParseBool(row[h["ok"]])
		passesOracle, _ := strconv.ParseBool(row[h["passes_oracle"]])
		durNS, _ := strconv.ParseInt(row[h["duration_ns"]], 10, 64)
		sessionID, _ := strconv.Atoi(row[h["session_id"]])
		out = append(out, runRecord{
			SessionID:    sessionID,
			Track:        row[h["track"]],
			Impl:         row[h["impl"]],
			Mode:         row[h["mode"]],
			Workload:     row[h["workload"]],
			Class:        row[h["class"]],
			OK:           ok,
			PassesOracle: passesOracle,
			DurationNS:   durNS,
		})
	}
	return out, nil
}

func buildStatsComparisons(runs []runRecord, alpha float64, resamples int) []statsComparison {
	type agg struct {
		durations   []float64
		oracleTotal int
		oraclePass  int
	}
	grouped := map[string]map[string]*agg{}
	for _, r := range runs {
		if r.Class != "valid" {
			continue
		}
		key := strings.Join([]string{r.Track, r.Mode, r.Workload, r.Class}, "|")
		if _, ok := grouped[key]; !ok {
			grouped[key] = map[string]*agg{}
		}
		a, ok := grouped[key][r.Impl]
		if !ok {
			a = &agg{}
			grouped[key][r.Impl] = a
		}
		a.oracleTotal++
		if r.PassesOracle {
			a.oraclePass++
		}
		if r.OK {
			a.durations = append(a.durations, float64(r.DurationNS)/1e6)
		}
	}

	keys := make([]string, 0, len(grouped))
	for k := range grouped {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	cmp := make([]statsComparison, 0, len(keys))
	for _, key := range keys {
		parts := strings.Split(key, "|")
		impls := grouped[key]
		a, okA := impls["schubfach"]
		b, okB := impls["json-canon"]
		if !okA || !okB || len(a.durations) < 2 || len(b.durations) < 2 {
			continue
		}

		meanA := avg(a.durations)
		meanB := avg(b.durations)
		medianA := percentile(a.durations, 0.50)
		medianB := percentile(b.durations, 0.50)
		winner := "schubfach"
		speedup := meanB / meanA
		if meanB < meanA {
			winner = "json-canon"
			speedup = meanA / meanB
		}

		ciLo, ciHi := bootstrapSpeedupCI(a.durations, b.durations, winner, resamples, key)
		pval := permutationPValue(a.durations, b.durations, resamples, key)
		effect := cohenD(a.durations, b.durations)

		// Noise floor = standard deviation of the faster impl (equivalently CV × mean).
		// This quantifies measurement noise in the same units as the metric (ms).
		fasterMean := meanA
		fasterCV := coefficientOfVariation(a.durations)
		fasterN := float64(len(a.durations))
		if meanB < meanA {
			fasterMean = meanB
			fasterCV = coefficientOfVariation(b.durations)
			fasterN = float64(len(b.durations))
		}
		noiseFloor := fasterCV * fasterMean
		minObsEff := 0.0
		if fasterN > 0 {
			// Two-sample minimum detectable effect at 80% power (α = 0.05):
			// Δ/μ = (z_{α/2} + z_β) × √(2/n) × CV × 100
			// where z_{α/2} + z_β ≈ 1.96 + 0.84 = 2.8, and √2 ≈ 1.4142.
			// This uses only the faster implementation's CV and n as an approximation.
			minObsEff = (2.8 * math.Sqrt2 * fasterCV * 100.0) / math.Sqrt(fasterN)
		}

		cmp = append(cmp, statsComparison{
			Track:            parts[0],
			Mode:             parts[1],
			Workload:         parts[2],
			Class:            parts[3],
			ImplA:            "schubfach",
			ImplB:            "json-canon",
			NImplA:           len(a.durations),
			NImplB:           len(b.durations),
			MeanMSImplA:      meanA,
			MeanMSImplB:      meanB,
			MedianMSImplA:    medianA,
			MedianMSImplB:    medianB,
			CVImplA:          coefficientOfVariation(a.durations),
			CVImplB:          coefficientOfVariation(b.durations),
			Winner:           winner,
			Speedup:          speedup,
			CI95Low:          ciLo,
			CI95High:         ciHi,
			PValue:           pval,
			EffectSizeCohenD: effect,
			Significant:      pval < alpha,
			NoiseFloorMS:          noiseFloor,
			MinObservableEffPct:   minObsEff,
			OraclePassRateA:       safeRate(a.oraclePass, a.oracleTotal),
			OraclePassRateB:       safeRate(b.oraclePass, b.oracleTotal),
		})
	}
	return cmp
}

// applyBenjaminiHochberg applies the Benjamini-Hochberg FDR correction
// to all comparisons and sets PValueAdjusted and SignificantBH fields.
func applyBenjaminiHochberg(comparisons []statsComparison, alpha float64) {
	m := len(comparisons)
	if m == 0 {
		return
	}

	// Build index sorted by raw p-value ascending.
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return comparisons[idx[i]].PValue < comparisons[idx[j]].PValue
	})

	// Compute adjusted p-values (step-up procedure).
	adjP := make([]float64, m)
	for rank, i := range idx {
		k := rank + 1 // 1-based rank
		adjP[rank] = comparisons[i].PValue * float64(m) / float64(k)
	}

	// Enforce monotonicity from the bottom up: adj_p[k] = min(adj_p[k], adj_p[k+1]).
	for rank := m - 2; rank >= 0; rank-- {
		if adjP[rank] > adjP[rank+1] {
			adjP[rank] = adjP[rank+1]
		}
	}

	// Clamp to [0, 1] and assign.
	for rank, i := range idx {
		p := adjP[rank]
		if p > 1.0 {
			p = 1.0
		}
		comparisons[i].PValueAdjusted = p
		comparisons[i].SignificantBH = p < alpha
	}
}

// comparisonSeed returns a deterministic seed unique to the comparison key and sample sizes.
func comparisonSeed(comparisonKey string, lenA, lenB int) int64 {
	h := fnv.New64a()
	h.Write([]byte(comparisonKey))
	return int64(h.Sum64()) ^ int64(lenA*1000003+lenB)
}

func bootstrapSpeedupCI(a, b []float64, winner string, resamples int, comparisonKey string) (float64, float64) {
	r := rand.New(rand.NewSource(comparisonSeed(comparisonKey, len(a), len(b))))

	// Compute observed statistic.
	obsA := avg(a)
	obsB := avg(b)
	var observed float64
	if winner == "json-canon" {
		observed = obsA / obsB
	} else {
		observed = obsB / obsA
	}

	// Generate bootstrap samples.
	samples := make([]float64, 0, resamples)
	for i := 0; i < resamples; i++ {
		ma := bootstrapMean(a, r)
		mb := bootstrapMean(b, r)
		if ma <= 0 || mb <= 0 {
			continue
		}
		ratio := mb / ma
		if winner == "json-canon" {
			ratio = ma / mb
		}
		samples = append(samples, ratio)
	}
	if len(samples) == 0 {
		return 1, 1
	}
	sort.Float64s(samples)

	// BCa bias-correction constant z0: proportion of bootstrap samples below observed.
	belowCount := 0
	for _, s := range samples {
		if s < observed {
			belowCount++
		}
	}
	z0 := normInv(float64(belowCount) / float64(len(samples)))

	// BCa acceleration constant a: jackknife estimate of skewness.
	acc := bcaAcceleration(a, b, winner)

	// Adjusted percentile indices.
	zAlpha := normInv(0.025)
	adjLo := normCDF(z0 + (z0+zAlpha)/(1-acc*(z0+zAlpha)))
	adjHi := normCDF(z0 + (z0-zAlpha)/(1-acc*(z0-zAlpha)))

	// Clamp to valid range.
	if adjLo < 0 {
		adjLo = 0
	}
	if adjHi > 1 {
		adjHi = 1
	}
	if adjLo >= adjHi {
		adjLo = 0.025
		adjHi = 0.975
	}

	lo := samples[int(math.Floor(float64(len(samples)-1)*adjLo))]
	hi := samples[int(math.Floor(float64(len(samples)-1)*adjHi))]
	return lo, hi
}

// bcaAcceleration computes the jackknife acceleration constant for the speedup ratio.
func bcaAcceleration(a, b []float64, winner string) float64 {
	n := len(a) + len(b)
	if n < 3 {
		return 0
	}

	// Compute leave-one-out speedup estimates.
	jackknife := make([]float64, n)
	for i := 0; i < len(a); i++ {
		aLoo := make([]float64, 0, len(a)-1)
		aLoo = append(aLoo, a[:i]...)
		aLoo = append(aLoo, a[i+1:]...)
		ma := avg(aLoo)
		mb := avg(b)
		if ma <= 0 || mb <= 0 {
			jackknife[i] = 1
			continue
		}
		if winner == "json-canon" {
			jackknife[i] = ma / mb
		} else {
			jackknife[i] = mb / ma
		}
	}
	for i := 0; i < len(b); i++ {
		bLoo := make([]float64, 0, len(b)-1)
		bLoo = append(bLoo, b[:i]...)
		bLoo = append(bLoo, b[i+1:]...)
		ma := avg(a)
		mb := avg(bLoo)
		if ma <= 0 || mb <= 0 {
			jackknife[len(a)+i] = 1
			continue
		}
		if winner == "json-canon" {
			jackknife[len(a)+i] = ma / mb
		} else {
			jackknife[len(a)+i] = mb / ma
		}
	}

	jMean := avg(jackknife)
	var num, den float64
	for _, j := range jackknife {
		d := jMean - j
		num += d * d * d
		den += d * d
	}
	if den == 0 {
		return 0
	}
	return num / (6.0 * math.Pow(den, 1.5))
}

func bootstrapMean(xs []float64, r *rand.Rand) float64 {
	if len(xs) == 0 {
		return 0
	}
	var sum float64
	for i := 0; i < len(xs); i++ {
		sum += xs[r.Intn(len(xs))]
	}
	return sum / float64(len(xs))
}

func permutationPValue(a, b []float64, resamples int, comparisonKey string) float64 {
	obs := math.Abs(avg(a) - avg(b))
	pooled := make([]float64, 0, len(a)+len(b))
	pooled = append(pooled, a...)
	pooled = append(pooled, b...)
	r := rand.New(rand.NewSource(comparisonSeed(comparisonKey, len(a), len(b)) + 1))
	count := 0
	for i := 0; i < resamples; i++ {
		r.Shuffle(len(pooled), func(i, j int) {
			pooled[i], pooled[j] = pooled[j], pooled[i]
		})
		ma := avg(pooled[:len(a)])
		mb := avg(pooled[len(a):])
		if math.Abs(ma-mb) >= obs {
			count++
		}
	}
	return float64(count+1) / float64(resamples+1)
}

func cohenD(a, b []float64) float64 {
	if len(a) < 2 || len(b) < 2 {
		return 0
	}
	meanA := avg(a)
	meanB := avg(b)
	varA := variance(a, meanA)
	varB := variance(b, meanB)
	pooled := math.Sqrt(((float64(len(a)-1) * varA) + (float64(len(b)-1) * varB)) / float64(len(a)+len(b)-2))
	if pooled == 0 {
		return 0
	}
	return (meanA - meanB) / pooled
}

func variance(xs []float64, mean float64) float64 {
	if len(xs) < 2 {
		return 0
	}
	var ss float64
	for _, x := range xs {
		d := x - mean
		ss += d * d
	}
	return ss / float64(len(xs)-1)
}

func coefficientOfVariation(xs []float64) float64 {
	if len(xs) < 2 {
		return 0
	}
	m := avg(xs)
	if m == 0 {
		return 0
	}
	return math.Sqrt(variance(xs, m)) / m
}

func safeRate(pass, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(pass) / float64(total)
}

// normCDF returns the cumulative distribution function of the standard normal distribution.
func normCDF(x float64) float64 {
	return 0.5 * (1.0 + math.Erf(x/math.Sqrt2))
}

// normInv returns the inverse CDF (quantile function) of the standard normal distribution.
// Uses the rational approximation by Abramowitz and Stegun.
func normInv(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}
	if p >= 1 {
		return math.Inf(1)
	}
	if p == 0.5 {
		return 0
	}
	if p > 0.5 {
		return -normInv(1 - p)
	}
	// Rational approximation for 0 < p <= 0.5.
	t := math.Sqrt(-2.0 * math.Log(p))
	// Coefficients from Abramowitz and Stegun formula 26.2.23.
	c0 := 2.515517
	c1 := 0.802853
	c2 := 0.010328
	d1 := 1.432788
	d2 := 0.189269
	d3 := 0.001308
	return -(t - (c0+c1*t+c2*t*t)/(1.0+d1*t+d2*t*t+d3*t*t*t))
}

func renderStatsMarkdown(report statsReport, runsPath string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Statistical Comparison\n\n")
	fmt.Fprintf(&b, "Generated at: %s\n\n", report.GeneratedAtUTC)
	fmt.Fprintf(&b, "Source: `%s`\n\n", runsPath)
	fmt.Fprintf(&b, "Protocol: permutation test + BCa bootstrap 95%% CI, alpha=%.4f, resamples=%d\n\n", report.Alpha, report.Resamples)
	fmt.Fprintf(&b, "Benjamini-Hochberg FDR correction applied across %d comparisons.\n\n", len(report.Comparisons))
	fmt.Fprintf(&b, "| track | mode | workload | winner | speedup | ci95 | p-value | p-adj | effect d | sig | sig-BH | noise floor (ms) | min obs effect (%%) |\n")
	fmt.Fprintf(&b, "|---|---|---|---|---:|---|---:|---:|---:|---|---|---:|---:|\n")
	for _, c := range report.Comparisons {
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %.3fx | [%.3f, %.3f] | %.4f | %.4f | %.3f | %t | %t | %.4f | %.2f |\n",
			c.Track, c.Mode, c.Workload, c.Winner, c.Speedup, c.CI95Low, c.CI95High,
			c.PValue, c.PValueAdjusted, c.EffectSizeCohenD, c.Significant, c.SignificantBH,
			c.NoiseFloorMS, c.MinObservableEffPct)
	}
	return b.String()
}
