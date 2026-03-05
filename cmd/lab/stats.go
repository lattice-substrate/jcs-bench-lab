package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
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
	EffectSizeCohenD float64 `json:"effect_size_cohen_d"`
	Significant      bool    `json:"significant"`
	PracticalWin          bool    `json:"practical_win"`
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

func runStats(runsPath string, alpha float64, resamples int) error {
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

	runs, err := loadRunsCSV(runsPath)
	if err != nil {
		return err
	}

	comparisons := buildStatsComparisons(runs, alpha, resamples)
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

		ciLo, ciHi := bootstrapSpeedupCI(a.durations, b.durations, winner, resamples)
		pval := permutationPValue(a.durations, b.durations, resamples)
		effect := cohenD(a.durations, b.durations)
		practical := speedup >= 1.03

		// Noise floor = CV × mean of faster impl.
		// Min observable effect ≈ (2 × CV × 100) / sqrt(N) (approximate 80% power threshold).
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
			minObsEff = (2.0 * fasterCV * 100.0) / math.Sqrt(fasterN)
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
			PracticalWin:          practical,
			NoiseFloorMS:          noiseFloor,
			MinObservableEffPct:   minObsEff,
			OraclePassRateA:       safeRate(a.oraclePass, a.oracleTotal),
			OraclePassRateB:       safeRate(b.oraclePass, b.oracleTotal),
		})
	}
	return cmp
}

func bootstrapSpeedupCI(a, b []float64, winner string, resamples int) (float64, float64) {
	r := rand.New(rand.NewSource(int64(len(a)*1000003 + len(b))))
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
	lo := samples[int(math.Floor(float64(len(samples)-1)*0.025))]
	hi := samples[int(math.Floor(float64(len(samples)-1)*0.975))]
	return lo, hi
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

func permutationPValue(a, b []float64, resamples int) float64 {
	obs := math.Abs(avg(a) - avg(b))
	pooled := make([]float64, 0, len(a)+len(b))
	pooled = append(pooled, a...)
	pooled = append(pooled, b...)
	r := rand.New(rand.NewSource(int64(len(a)*131 + len(b)*271 + 17)))
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

func renderStatsMarkdown(report statsReport, runsPath string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Statistical Comparison\n\n")
	fmt.Fprintf(&b, "Generated at: %s\n\n", report.GeneratedAtUTC)
	fmt.Fprintf(&b, "Source: `%s`\n\n", runsPath)
	fmt.Fprintf(&b, "Protocol: permutation test + bootstrap 95%% CI, alpha=%.4f, resamples=%d\n\n", report.Alpha, report.Resamples)
	fmt.Fprintf(&b, "| track | mode | workload | winner | speedup | ci95 | p-value | effect d | significant | practical | noise floor (ms) | min obs effect (%%) |\n")
	fmt.Fprintf(&b, "|---|---|---|---|---:|---|---:|---:|---|---|---:|---:|\n")
	for _, c := range report.Comparisons {
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %.3fx | [%.3f, %.3f] | %.4f | %.3f | %t | %t | %.4f | %.2f |\n",
			c.Track, c.Mode, c.Workload, c.Winner, c.Speedup, c.CI95Low, c.CI95High, c.PValue, c.EffectSizeCohenD, c.Significant, c.PracticalWin, c.NoiseFloorMS, c.MinObservableEffPct)
	}
	return b.String()
}
