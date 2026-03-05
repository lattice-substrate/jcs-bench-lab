package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type apiBenchSample struct {
	Impl        string
	Operation   string
	Workload    string
	Iters       int
	NSPerOp     float64
	MBPerSec    float64
	BytesPerOp  float64
	AllocsPerOp float64
}

type cliSummarySample struct {
	Track           string
	Impl            string
	Mode            string
	Workload        string
	Class           string
	AvgMS           float64
	P50MS           float64
	AvgMaxRSSKB     float64
	AvgThroughputMB float64
	OraclePasses    int
	N               int
}

type workloadCmp struct {
	Workload string
	Mode     string
	Winner   string
	Speedup  float64
	Schub    float64
	JSON     float64
}

func runBenchstat(apiPath, cliPath, qualityPath string, allowFallback bool) error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	if err := ensureDirs(root); err != nil {
		return err
	}

	apiPath = choosePath(apiPath, filepath.Join(root, "results", "latest-api-bench.txt"))
	cliPath = choosePath(cliPath, filepath.Join(root, "results", "latest-cli-summary.csv"))
	qualityPath = choosePath(qualityPath, filepath.Join(root, "results", "latest-quality.json"))

	apiSamples, err := loadAPIBenchSamples(apiPath)
	if err != nil {
		return fmt.Errorf("load api samples: %w", err)
	}
	cliSamples, err := loadCLISummary(cliPath)
	if err != nil {
		return fmt.Errorf("load cli samples: %w", err)
	}
	quality, err := loadQualityReport(qualityPath)
	if err != nil {
		return fmt.Errorf("load quality report: %w", err)
	}

	bsOut, usedExternal, bsErr := runExternalBenchstat(apiSamples, allowFallback)
	if bsErr != nil {
		return bsErr
	}
	md := renderBenchstatMarkdown(apiPath, cliPath, qualityPath, apiSamples, cliSamples, quality, bsOut, usedExternal)

	stamp := time.Now().UTC().Format("20060102T150405Z")
	outPath := filepath.Join(root, "results", "benchstat-"+stamp+".md")
	latest := filepath.Join(root, "results", "latest-benchstat.md")
	if err := os.WriteFile(outPath, []byte(md), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latest, []byte(md), 0o644); err != nil {
		return err
	}

	fmt.Printf("benchstat complete\n- %s\n- %s\n", outPath, latest)
	return nil
}

func runProfileAPI(count int, benchtime string) error {
	if count < 1 {
		return errors.New("count must be >= 1")
	}
	root, err := repoRoot()
	if err != nil {
		return err
	}
	if err := ensureDirs(root); err != nil {
		return err
	}

	stamp := time.Now().UTC().Format("20060102T150405Z")
	resultsDir := filepath.Join(root, "results")
	patterns := []struct {
		impl    string
		pattern string
	}{
		{impl: "schubfach", pattern: "^BenchmarkAPI(Canonicalize|Verify)Schubfach/"},
		{impl: "json-canon", pattern: "^BenchmarkAPI(Canonicalize|Verify)JSONCanon/"},
	}

	tops := map[string]map[string]float64{}
	for _, p := range patterns {
		cpuPath := filepath.Join(resultsDir, "api-prof-"+p.impl+"-cpu-"+stamp+".pprof")
		memPath := filepath.Join(resultsDir, "api-prof-"+p.impl+"-mem-"+stamp+".pprof")
		outPath := filepath.Join(resultsDir, "api-prof-"+p.impl+"-"+stamp+".txt")
		latestOut := filepath.Join(resultsDir, "latest-api-prof-"+p.impl+".txt")
		latestCPU := filepath.Join(resultsDir, "latest-api-prof-"+p.impl+"-cpu.pprof")
		latestMem := filepath.Join(resultsDir, "latest-api-prof-"+p.impl+"-mem.pprof")
		cpuTop := filepath.Join(resultsDir, "api-prof-"+p.impl+"-cpu-top-"+stamp+".txt")
		memTop := filepath.Join(resultsDir, "api-prof-"+p.impl+"-mem-top-"+stamp+".txt")
		latestCPUTop := filepath.Join(resultsDir, "latest-api-prof-"+p.impl+"-cpu-top.txt")
		latestMemTop := filepath.Join(resultsDir, "latest-api-prof-"+p.impl+"-mem-top.txt")

		args := []string{
			"test", "./internal/apibench",
			"-run", "^$",
			"-bench", p.pattern,
			"-benchmem",
			"-count", strconv.Itoa(count),
			"-benchtime", benchtime,
			"-cpuprofile", cpuPath,
			"-memprofile", memPath,
		}
		cmd := exec.Command("go", args...)
		cmd.Dir = root
		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("profile %s failed: %w\n%s", p.impl, err, buf.String())
		}
		if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
			return err
		}
		if err := os.WriteFile(latestOut, buf.Bytes(), 0o644); err != nil {
			return err
		}
		if err := copyFile(cpuPath, latestCPU); err != nil {
			return err
		}
		if err := copyFile(memPath, latestMem); err != nil {
			return err
		}
		if err := writePprofTop(root, cpuPath, cpuTop); err != nil {
			return err
		}
		if err := writePprofTop(root, memPath, memTop); err != nil {
			return err
		}
		if err := copyFile(cpuTop, latestCPUTop); err != nil {
			return err
		}
		if err := copyFile(memTop, latestMemTop); err != nil {
			return err
		}

		topEntries, err := parsePprofTop(cpuTop)
		if err == nil {
			tops[p.impl] = topEntries
		}
	}

	analysis := renderProfileAnalysisMarkdown(tops)
	analysisPath := filepath.Join(resultsDir, "api-profile-analysis-"+stamp+".md")
	latestAnalysisPath := filepath.Join(resultsDir, "latest-api-profile-analysis.md")
	if err := os.WriteFile(analysisPath, []byte(analysis), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latestAnalysisPath, []byte(analysis), 0o644); err != nil {
		return err
	}

	fmt.Printf("profile-api complete\n- %s\n", filepath.Join(resultsDir, "latest-api-prof-schubfach.txt"))
	fmt.Printf("- %s\n", filepath.Join(resultsDir, "latest-api-prof-json-canon.txt"))
	fmt.Printf("- %s\n", latestAnalysisPath)
	return nil
}

func writePprofTop(root, profilePath, outPath string) error {
	cmd := exec.Command("go", "tool", "pprof", "-top", "-nodecount", "30", profilePath)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pprof top failed for %s: %w\n%s", profilePath, err, string(out))
	}
	return os.WriteFile(outPath, out, 0o644)
}

func parsePprofTop(path string) (map[string]float64, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	re := regexp.MustCompile(`^\s*([0-9.]+)%\s+\S+\s+\S+\s+\S+\s+(.+)$`)
	out := map[string]float64{}
	for _, ln := range lines {
		m := re.FindStringSubmatch(strings.TrimSpace(ln))
		if len(m) != 3 {
			continue
		}
		pct, err := strconv.ParseFloat(m[1], 64)
		if err != nil {
			continue
		}
		fn := strings.TrimSpace(m[2])
		if fn == "" {
			continue
		}
		out[fn] = pct
	}
	return out, nil
}

func renderProfileAnalysisMarkdown(tops map[string]map[string]float64) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# API Profile Differential Analysis\n\n")
	fmt.Fprintf(&b, "Generated at: %s\n\n", time.Now().UTC().Format(time.RFC3339))
	if len(tops) == 0 {
		fmt.Fprintf(&b, "No profile top data available.\n")
		return b.String()
	}
	schub := tops["schubfach"]
	jsonCanon := tops["json-canon"]
	if len(schub) == 0 || len(jsonCanon) == 0 {
		fmt.Fprintf(&b, "Insufficient profile data for comparative hotspot attribution.\n")
		return b.String()
	}
	fmt.Fprintf(&b, "| function | schubfach flat%% | json-canon flat%% | delta (schub - json) |\n")
	fmt.Fprintf(&b, "|---|---:|---:|---:|\n")
	seen := map[string]bool{}
	for fn := range schub {
		seen[fn] = true
	}
	for fn := range jsonCanon {
		seen[fn] = true
	}
	fns := make([]string, 0, len(seen))
	for fn := range seen {
		fns = append(fns, fn)
	}
	sort.Strings(fns)
	for _, fn := range fns {
		s := schub[fn]
		j := jsonCanon[fn]
		if s == 0 && j == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %.2f | %.2f | %.2f |\n", fn, s, j, s-j)
	}
	return b.String()
}

func copyFile(src, dst string) error {
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, b, 0o644)
}

func runReport(apiPath, cliPath, qualityPath, benchstatPath, conformancePath, statsPath, fuzzPath string) error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	if err := ensureDirs(root); err != nil {
		return err
	}

	apiPath = choosePath(apiPath, filepath.Join(root, "results", "latest-api-bench.txt"))
	cliPath = choosePath(cliPath, filepath.Join(root, "results", "latest-cli-summary.csv"))
	qualityPath = choosePath(qualityPath, filepath.Join(root, "results", "latest-quality.json"))
	benchstatPath = choosePath(benchstatPath, filepath.Join(root, "results", "latest-benchstat.md"))
	conformancePath = choosePath(conformancePath, filepath.Join(root, "results", "latest-conformance.json"))
	statsPath = choosePath(statsPath, filepath.Join(root, "results", "latest-stats.json"))
	fuzzPath = choosePath(fuzzPath, filepath.Join(root, "results", "latest-fuzz.json"))

	apiSamples, err := loadAPIBenchSamples(apiPath)
	if err != nil {
		return err
	}
	cliSamples, err := loadCLISummary(cliPath)
	if err != nil {
		return err
	}
	quality, err := loadQualityReport(qualityPath)
	if err != nil {
		return err
	}
	conf, err := loadConformanceReport(conformancePath)
	if err != nil {
		return err
	}
	stats, err := loadStatsReport(statsPath)
	if err != nil {
		return err
	}
	fuzz, err := loadFuzzReport(fuzzPath)
	if err != nil {
		return err
	}
	benchstatSnippet := loadBenchstatSnippet(benchstatPath)

	report := renderReportMarkdown(apiSamples, cliSamples, quality, conf, stats, fuzz, benchstatSnippet, apiPath, cliPath, qualityPath, benchstatPath, conformancePath, statsPath, fuzzPath)

	stamp := time.Now().UTC().Format("20060102T150405Z")
	reportPath := filepath.Join(root, "REPORT.md")
	archivePath := filepath.Join(root, "results", "report-"+stamp+".md")
	latestPath := filepath.Join(root, "results", "latest-report.md")
	if err := os.WriteFile(reportPath, []byte(report), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(archivePath, []byte(report), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latestPath, []byte(report), 0o644); err != nil {
		return err
	}

	fmt.Printf("report complete\n- %s\n- %s\n- %s\n", reportPath, archivePath, latestPath)
	return nil
}

func choosePath(flagPath, fallback string) string {
	if strings.TrimSpace(flagPath) != "" {
		return flagPath
	}
	return fallback
}

func loadQualityReport(path string) (qualityReport, error) {
	var q qualityReport
	b, err := os.ReadFile(path)
	if err != nil {
		return q, err
	}
	if err := json.Unmarshal(b, &q); err != nil {
		return q, err
	}
	if q.DeterminismFailures == nil {
		q.DeterminismFailures = []string{}
	}
	if q.OutputEqualityFailures == nil {
		q.OutputEqualityFailures = []string{}
	}
	if q.InvalidFailureParityIssues == nil {
		q.InvalidFailureParityIssues = []string{}
	}
	if q.OracleMismatches == nil {
		q.OracleMismatches = []string{}
	}
	if q.CaseFailures == nil {
		q.CaseFailures = []string{}
	}
	if q.Notes == nil {
		q.Notes = map[string]string{}
	}
	return q, nil
}

func loadAPIBenchSamples(path string) ([]apiBenchSample, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	out := make([]apiBenchSample, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "BenchmarkAPI") {
			continue
		}
		s, ok := parseAPIBenchmarkLine(line)
		if ok {
			out = append(out, s)
		}
	}
	return out, nil
}

func parseAPIBenchmarkLine(line string) (apiBenchSample, bool) {
	fields := strings.Fields(line)
	if len(fields) < 4 {
		return apiBenchSample{}, false
	}
	name := fields[0]
	if !strings.HasPrefix(name, "BenchmarkAPI") {
		return apiBenchSample{}, false
	}

	rest := strings.TrimPrefix(name, "BenchmarkAPI")
	parts := strings.SplitN(rest, "/", 2)
	if len(parts) != 2 {
		return apiBenchSample{}, false
	}
	prefix := parts[0]
	impl := ""
	op := ""
	switch {
	case strings.HasSuffix(prefix, "Schubfach"):
		impl = "schubfach"
		op = strings.TrimSuffix(prefix, "Schubfach")
	case strings.HasSuffix(prefix, "JSONCanon"):
		impl = "json-canon"
		op = strings.TrimSuffix(prefix, "JSONCanon")
	default:
		return apiBenchSample{}, false
	}
	workload := trimBenchCPUSuffix(parts[1])
	iters, _ := strconv.Atoi(fields[1])

	s := apiBenchSample{Impl: impl, Operation: strings.ToLower(op), Workload: workload, Iters: iters}
	for i := 2; i+1 < len(fields); i += 2 {
		val, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			continue
		}
		switch fields[i+1] {
		case "ns/op":
			s.NSPerOp = val
		case "MB/s":
			s.MBPerSec = val
		case "B/op":
			s.BytesPerOp = val
		case "allocs/op":
			s.AllocsPerOp = val
		}
	}
	return s, true
}

func trimBenchCPUSuffix(s string) string {
	i := strings.LastIndexByte(s, '-')
	if i <= 0 || i+1 >= len(s) {
		return s
	}
	for _, r := range s[i+1:] {
		if r < '0' || r > '9' {
			return s
		}
	}
	return s[:i]
}

func loadCLISummary(path string) ([]cliSummarySample, error) {
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
	if len(rows) < 1 {
		return nil, errors.New("empty cli summary")
	}
	h := indexHeader(rows[0])
	required := []string{"track", "impl", "mode", "workload", "class", "avg_ms", "p50_ms", "avg_max_rss_kb", "avg_throughput_mb_s", "oracle_passes", "n"}
	for _, k := range required {
		if _, ok := h[k]; !ok {
			return nil, fmt.Errorf("missing %s in cli summary header", k)
		}
	}

	out := make([]cliSummarySample, 0, len(rows)-1)
	for _, row := range rows[1:] {
		avgMS, _ := strconv.ParseFloat(row[h["avg_ms"]], 64)
		p50MS, _ := strconv.ParseFloat(row[h["p50_ms"]], 64)
		avgRSS, _ := strconv.ParseFloat(row[h["avg_max_rss_kb"]], 64)
		avgTP, _ := strconv.ParseFloat(row[h["avg_throughput_mb_s"]], 64)
		oraclePasses, _ := strconv.Atoi(row[h["oracle_passes"]])
		n, _ := strconv.Atoi(row[h["n"]])
		out = append(out, cliSummarySample{
			Track:           row[h["track"]],
			Impl:            row[h["impl"]],
			Mode:            row[h["mode"]],
			Workload:        row[h["workload"]],
			Class:           row[h["class"]],
			AvgMS:           avgMS,
			P50MS:           p50MS,
			AvgMaxRSSKB:     avgRSS,
			AvgThroughputMB: avgTP,
			OraclePasses:    oraclePasses,
			N:               n,
		})
	}
	return out, nil
}

func indexHeader(header []string) map[string]int {
	out := make(map[string]int, len(header))
	for i, k := range header {
		out[strings.TrimSpace(k)] = i
	}
	return out
}

func runExternalBenchstat(samples []apiBenchSample, allowFallback bool) (string, bool, error) {
	if len(samples) == 0 {
		return "", false, errors.New("no api samples available for benchstat")
	}
	if _, err := exec.LookPath("benchstat"); err != nil {
		if allowFallback {
			return "", false, nil
		}
		return "", false, errors.New("benchstat binary not found in PATH (use -allow-fallback for non-gating exploratory output)")
	}

	tmpDir, err := os.MkdirTemp("", "jcs-benchstat-*")
	if err != nil {
		return "", false, err
	}
	defer os.RemoveAll(tmpDir)

	aPath := filepath.Join(tmpDir, "schubfach.txt")
	bPath := filepath.Join(tmpDir, "json-canon.txt")
	aLines, bLines := synthBenchstatInputs(samples)
	if len(aLines) == 0 || len(bLines) == 0 {
		return "", false, errors.New("insufficient paired api benchmark samples for benchstat")
	}
	if err := os.WriteFile(aPath, []byte(strings.Join(aLines, "\n")+"\n"), 0o644); err != nil {
		return "", false, err
	}
	if err := os.WriteFile(bPath, []byte(strings.Join(bLines, "\n")+"\n"), 0o644); err != nil {
		return "", false, err
	}
	cmd := exec.Command("benchstat", aPath, bPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", true, fmt.Errorf("benchstat failed: %w\n%s", err, string(out))
	}
	return string(out), true, nil
}

func synthBenchstatInputs(samples []apiBenchSample) (schubLines []string, jsonLines []string) {
	schubLines = make([]string, 0, len(samples))
	jsonLines = make([]string, 0, len(samples))
	for _, s := range samples {
		if s.NSPerOp <= 0 {
			continue
		}
		benchName := fmt.Sprintf("Benchmark%s/%s", strings.Title(s.Operation), s.Workload)
		line := fmt.Sprintf("%s %d %.4f ns/op %.4f MB/s %.4f B/op %.4f allocs/op",
			benchName, maxInt(s.Iters, 1), s.NSPerOp, s.MBPerSec, s.BytesPerOp, s.AllocsPerOp)
		switch s.Impl {
		case "schubfach":
			schubLines = append(schubLines, line)
		case "json-canon":
			jsonLines = append(jsonLines, line)
		}
	}
	sort.Strings(schubLines)
	sort.Strings(jsonLines)
	return schubLines, jsonLines
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func renderBenchstatMarkdown(apiPath, cliPath, qualityPath string, apiSamples []apiBenchSample, cliSamples []cliSummarySample, quality qualityReport, benchstatOut string, usedExternal bool) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Benchstat Snapshot\n\n")
	fmt.Fprintf(&b, "Generated at: %s\n\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "Sources:\n")
	fmt.Fprintf(&b, "- API benchmark: `%s`\n", apiPath)
	fmt.Fprintf(&b, "- CLI summary: `%s`\n", cliPath)
	fmt.Fprintf(&b, "- Quality report: `%s`\n\n", qualityPath)

	apiCmp := compareAPIWorkloads(apiSamples)
	cliCanonCmp := compareCLIWorkloads(cliSamples, "canonicalize")
	cliVerifyCmp := compareCLIWorkloads(cliSamples, "verify")

	fmt.Fprintf(&b, "## Quality Gate\n\n")
	fmt.Fprintf(&b, "- determinism_failures: %d\n", len(quality.DeterminismFailures))
	fmt.Fprintf(&b, "- output_equality_failures: %d\n", len(quality.OutputEqualityFailures))
	fmt.Fprintf(&b, "- invalid_failure_parity_issues: %d\n", len(quality.InvalidFailureParityIssues))
	fmt.Fprintf(&b, "- oracle_mismatches: %d\n\n", len(quality.OracleMismatches))

	fmt.Fprintf(&b, "## CLI Canonicalize (valid workloads)\n\n")
	b.WriteString(renderCmpTable(cliCanonCmp, "avg_ms"))
	b.WriteString("\n")

	fmt.Fprintf(&b, "## CLI Verify (valid workloads)\n\n")
	b.WriteString(renderCmpTable(cliVerifyCmp, "avg_ms"))
	b.WriteString("\n")

	fmt.Fprintf(&b, "## API Benchmarks (ns/op)\n\n")
	b.WriteString(renderCmpTable(apiCmp, "ns/op"))
	b.WriteString("\n")

	fmt.Fprintf(&b, "## benchstat Output\n\n")
	if usedExternal {
		fmt.Fprintf(&b, "```text\n%s\n```\n", strings.TrimSpace(benchstatOut))
	} else {
		fmt.Fprintf(&b, "benchstat unavailable; fallback summary above is non-inferential and not CI-gating.\n")
	}

	return b.String()
}

func compareAPIWorkloads(samples []apiBenchSample) []workloadCmp {
	type agg struct {
		workload string
		impl     string
		ns       []float64
	}
	m := map[string]*agg{}
	for _, s := range samples {
		if s.NSPerOp <= 0 {
			continue
		}
		wl := s.Operation + "/" + s.Workload
		key := wl + "|" + s.Impl
		a, ok := m[key]
		if !ok {
			a = &agg{workload: wl, impl: s.Impl}
			m[key] = a
		}
		a.ns = append(a.ns, s.NSPerOp)
	}
	byWorkload := map[string]map[string]float64{}
	for _, a := range m {
		if _, ok := byWorkload[a.workload]; !ok {
			byWorkload[a.workload] = map[string]float64{}
		}
		byWorkload[a.workload][a.impl] = avg(a.ns)
	}

	out := make([]workloadCmp, 0, len(byWorkload))
	for workload, impls := range byWorkload {
		schub, okS := impls["schubfach"]
		jsonCanon, okJ := impls["json-canon"]
		if !okS || !okJ || schub <= 0 || jsonCanon <= 0 {
			continue
		}
		winner := "schubfach"
		speedup := jsonCanon / schub
		if jsonCanon < schub {
			winner = "json-canon"
			speedup = schub / jsonCanon
		}
		out = append(out, workloadCmp{Workload: workload, Mode: "api", Winner: winner, Speedup: speedup, Schub: schub, JSON: jsonCanon})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Workload < out[j].Workload })
	return out
}

func compareCLIWorkloads(samples []cliSummarySample, mode string) []workloadCmp {
	byWorkload := map[string]map[string]float64{}
	for _, s := range samples {
		if s.Class != "valid" || s.Mode != mode || s.AvgMS <= 0 {
			continue
		}
		if _, ok := byWorkload[s.Workload]; !ok {
			byWorkload[s.Workload] = map[string]float64{}
		}
		byWorkload[s.Workload][s.Impl] = s.AvgMS
	}
	out := make([]workloadCmp, 0, len(byWorkload))
	for workload, impls := range byWorkload {
		schub, okS := impls["schubfach"]
		jsonCanon, okJ := impls["json-canon"]
		if !okS || !okJ {
			continue
		}
		winner := "schubfach"
		speedup := jsonCanon / schub
		if jsonCanon < schub {
			winner = "json-canon"
			speedup = schub / jsonCanon
		}
		out = append(out, workloadCmp{Workload: workload, Mode: mode, Winner: winner, Speedup: speedup, Schub: schub, JSON: jsonCanon})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Workload < out[j].Workload })
	return out
}

func renderCmpTable(rows []workloadCmp, unit string) string {
	if len(rows) == 0 {
		return "_No paired samples found._\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "| workload | schubfach (%s) | json-canon (%s) | winner | speedup |\n", unit, unit)
	fmt.Fprintf(&b, "|---|---:|---:|---|---:|\n")
	for _, r := range rows {
		fmt.Fprintf(&b, "| %s | %.3f | %.3f | %s | %.2fx |\n", r.Workload, r.Schub, r.JSON, r.Winner, r.Speedup)
	}
	return b.String()
}

func loadBenchstatSnippet(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	lines := strings.Split(string(b), "\n")
	if len(lines) > 80 {
		lines = lines[:80]
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func renderReportMarkdown(apiSamples []apiBenchSample, cliSamples []cliSummarySample, quality qualityReport, conf conformanceReport, stats statsReport, fuzz fuzzReport, benchstatSnippet, apiPath, cliPath, qualityPath, benchstatPath, conformancePath, statsPath, fuzzPath string) string {
	cliCanon := compareCLIWorkloads(cliSamples, "canonicalize")
	cliVerify := compareCLIWorkloads(cliSamples, "verify")
	apiCmp := compareAPIWorkloads(apiSamples)

	status := recommendation(quality, conf, stats, fuzz)

	var b strings.Builder
	fmt.Fprintf(&b, "# Benchmark Report\n\n")
	fmt.Fprintf(&b, "Generated at: %s\n\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "Sources:\n")
	fmt.Fprintf(&b, "- `%s`\n", apiPath)
	fmt.Fprintf(&b, "- `%s`\n", cliPath)
	fmt.Fprintf(&b, "- `%s`\n", qualityPath)
	fmt.Fprintf(&b, "- `%s`\n", benchstatPath)
	fmt.Fprintf(&b, "- `%s`\n", conformancePath)
	fmt.Fprintf(&b, "- `%s`\n\n", statsPath)
	fmt.Fprintf(&b, "- `%s`\n\n", fuzzPath)

	fmt.Fprintf(&b, "## Executive Summary\n\n")
	fmt.Fprintf(&b, "- Conformance failures: %d\n", conf.FailureCount)
	fmt.Fprintf(&b, "- Quality oracle mismatches: %d\n", len(quality.OracleMismatches))
	fmt.Fprintf(&b, "- Differential fuzz failures: %d\n", len(fuzz.Failures))
	fmt.Fprintf(&b, "- Statistically significant practical wins: %d\n", countSignificantPractical(stats.Comparisons))
	fmt.Fprintf(&b, "- Recommendation status: %s\n\n", status)

	fmt.Fprintf(&b, "## Conformance Evidence\n\n")
	for src, s := range conf.SummaryBySrc {
		fmt.Fprintf(&b, "- %s: total=%d passed=%d failed=%d\n", src, s.Total, s.Passed, s.Failed)
	}
	if conf.FailureCount > 0 {
		fmt.Fprintf(&b, "\nFailing case IDs:\n")
		fail := collectFailingCaseIDs(conf.Cases, 40)
		for _, id := range fail {
			fmt.Fprintf(&b, "- %s\n", id)
		}
	}
	b.WriteString("\n")

	fmt.Fprintf(&b, "## Quality Findings\n\n")
	writeIssueList(&b, "determinism_failures", quality.DeterminismFailures)
	writeIssueList(&b, "output_equality_failures", quality.OutputEqualityFailures)
	writeIssueList(&b, "invalid_failure_parity_issues", quality.InvalidFailureParityIssues)
	writeIssueList(&b, "oracle_mismatches", quality.OracleMismatches)
	b.WriteString("\n")

	fmt.Fprintf(&b, "## Performance Winners by Workload\n\n")
	fmt.Fprintf(&b, "### CLI canonicalize\n\n")
	b.WriteString(renderCmpTable(cliCanon, "avg_ms"))
	b.WriteString("\n")
	fmt.Fprintf(&b, "### CLI verify\n\n")
	b.WriteString(renderCmpTable(cliVerify, "avg_ms"))
	b.WriteString("\n")
	fmt.Fprintf(&b, "### API\n\n")
	b.WriteString(renderCmpTable(apiCmp, "ns/op"))
	b.WriteString("\n")

	fmt.Fprintf(&b, "## Statistical Inference\n\n")
	fmt.Fprintf(&b, "| track | mode | workload | winner | speedup | ci95 | p-value | practical |\n")
	fmt.Fprintf(&b, "|---|---|---|---|---:|---|---:|---|\n")
	for _, c := range stats.Comparisons {
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %.3fx | [%.3f, %.3f] | %.4f | %t |\n", c.Track, c.Mode, c.Workload, c.Winner, c.Speedup, c.CI95Low, c.CI95High, c.PValue, c.PracticalWin)
	}
	b.WriteString("\n")

	fmt.Fprintf(&b, "## benchstat Snippet\n\n")
	if benchstatSnippet == "" {
		fmt.Fprintf(&b, "_benchstat snippet unavailable_\n")
	} else {
		fmt.Fprintf(&b, "```text\n%s\n```\n", benchstatSnippet)
	}

	fmt.Fprintf(&b, "\n## Production Recommendation\n\n")
	fmt.Fprintf(&b, "%s\n", status)
	return b.String()
}

func writeIssueList(b *strings.Builder, label string, issues []string) {
	if len(issues) == 0 {
		fmt.Fprintf(b, "- %s: none\n", label)
		return
	}
	fmt.Fprintf(b, "- %s:\n", label)
	for _, v := range issues {
		fmt.Fprintf(b, "  - %s\n", v)
	}
}

func countSignificantPractical(rows []statsComparison) int {
	n := 0
	for _, r := range rows {
		if r.Significant && r.PracticalWin {
			n++
		}
	}
	return n
}

func collectFailingCaseIDs(rows []conformanceCaseResult, limit int) []string {
	ids := make([]string, 0, limit)
	seen := map[string]bool{}
	for _, r := range rows {
		if r.Pass {
			continue
		}
		if seen[r.CaseID] {
			continue
		}
		seen[r.CaseID] = true
		ids = append(ids, r.CaseID)
		if len(ids) >= limit {
			break
		}
	}
	sort.Strings(ids)
	return ids
}

func recommendation(q qualityReport, conf conformanceReport, stats statsReport, fuzz fuzzReport) string {
	if conf.FailureCount > 0 {
		return "Do not promote any implementation: authoritative conformance failures are present."
	}
	if len(fuzz.Failures) > 0 {
		return "Do not promote any implementation: differential/property fuzzing failures are present."
	}
	if len(q.DeterminismFailures) > 0 || len(q.OutputEqualityFailures) > 0 || len(q.InvalidFailureParityIssues) > 0 || len(q.OracleMismatches) > 0 {
		return "Do not promote any implementation: quality parity/oracle issues are present."
	}
	score := map[string]int{"schubfach": 0, "json-canon": 0}
	for _, c := range stats.Comparisons {
		if c.Significant && c.PracticalWin {
			score[c.Winner]++
		}
	}
	if score["schubfach"] == 0 && score["json-canon"] == 0 {
		return "No statistically significant practical winner. Keep dual-track evaluation and expand workload representativeness."
	}
	if score["schubfach"] > score["json-canon"] {
		return "Recommend `schubfach` based on statistically significant practical wins with conformance/oracle gates passing."
	}
	if score["json-canon"] > score["schubfach"] {
		return "Recommend `json-canon` based on statistically significant practical wins with conformance/oracle gates passing."
	}
	return "Tie under statistical decision policy. Keep dual-track deployment and collect additional workload evidence."
}

func geomean(xs []float64) float64 {
	if len(xs) == 0 {
		return 1
	}
	var sum float64
	n := 0
	for _, x := range xs {
		if x <= 0 || math.IsNaN(x) || math.IsInf(x, 0) {
			continue
		}
		sum += math.Log(x)
		n++
	}
	if n == 0 {
		return 1
	}
	return math.Exp(sum / float64(n))
}
