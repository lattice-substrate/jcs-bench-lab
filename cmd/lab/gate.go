package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type gateBaseline struct {
	ByKey map[string]statsComparison `json:"by_key"`
}

func runGate(conformancePath, statsPath, fuzzPath, baselinePath string, maxRegressionPct, alpha float64) error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	conformancePath = choosePath(conformancePath, filepath.Join(root, "results", "latest-conformance.json"))
	statsPath = choosePath(statsPath, filepath.Join(root, "results", "latest-stats.json"))
	fuzzPath = choosePath(fuzzPath, filepath.Join(root, "results", "latest-fuzz.json"))

	// Verify the conformance gate was active during benchmarking.
	qualityPath := filepath.Join(root, "results", "latest-quality.json")
	if qb, qerr := os.ReadFile(qualityPath); qerr == nil {
		var qr qualityReport
		if json.Unmarshal(qb, &qr) == nil && !qr.ConformanceGateActive {
			return errors.New("gate failed: benchmark was run with --skip-conformance; re-run without it")
		}
	}

	conf, err := loadConformanceReport(conformancePath)
	if err != nil {
		return err
	}
	if conf.FailureCount > 0 {
		return fmt.Errorf("gate failed: conformance failures=%d", conf.FailureCount)
	}

	stats, err := loadStatsReport(statsPath)
	if err != nil {
		return err
	}
	if len(stats.Comparisons) == 0 {
		return errors.New("gate failed: no statistical comparisons available")
	}
	fuzz, err := loadFuzzReport(fuzzPath)
	if err != nil {
		return err
	}
	if len(fuzz.Failures) > 0 {
		return fmt.Errorf("gate failed: fuzz failures=%d", len(fuzz.Failures))
	}

	if strings.TrimSpace(baselinePath) == "" {
		return errors.New("gate failed: baseline stats path is required")
	}
	baseline, err := loadStatsReport(baselinePath)
	if err != nil {
		return fmt.Errorf("load baseline stats: %w", err)
	}
	baseIdx := map[string]statsComparison{}
	for _, c := range baseline.Comparisons {
		baseIdx[statsKey(c)] = c
	}

	maxAllowed := 1.0 + maxRegressionPct/100.0
	violations := make([]string, 0)
	for _, c := range stats.Comparisons {
		base, ok := baseIdx[statsKey(c)]
		if !ok {
			continue
		}
		if !c.Significant || c.PValue >= alpha {
			continue
		}
		// degradation check from baseline perspective for each implementation mean.
		relA := c.MeanMSImplA / maxFloat(base.MeanMSImplA, 1e-9)
		relB := c.MeanMSImplB / maxFloat(base.MeanMSImplB, 1e-9)
		if relA > maxAllowed {
			violations = append(violations, fmt.Sprintf("%s: %s mean regression %.2f%%", statsKey(c), c.ImplA, (relA-1.0)*100))
		}
		if relB > maxAllowed {
			violations = append(violations, fmt.Sprintf("%s: %s mean regression %.2f%%", statsKey(c), c.ImplB, (relB-1.0)*100))
		}
	}

	if len(violations) > 0 {
		return fmt.Errorf("gate failed: %d performance regressions: %s", len(violations), strings.Join(violations, "; "))
	}
	fmt.Printf("gate passed\n- conformance: %s\n- stats: %s\n- fuzz: %s\n- baseline: %s\n", conformancePath, statsPath, fuzzPath, baselinePath)
	return nil
}

func loadConformanceReport(path string) (conformanceReport, error) {
	var c conformanceReport
	b, err := os.ReadFile(path)
	if err != nil {
		return c, err
	}
	if err := json.Unmarshal(b, &c); err != nil {
		return c, err
	}
	if c.Cases == nil {
		c.Cases = []conformanceCaseResult{}
	}
	if c.SummaryByImpl == nil {
		c.SummaryByImpl = map[string]caseSummary{}
	}
	if c.SummaryBySrc == nil {
		c.SummaryBySrc = map[string]caseSummary{}
	}
	return c, nil
}

func loadStatsReport(path string) (statsReport, error) {
	var s statsReport
	b, err := os.ReadFile(path)
	if err != nil {
		return s, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return s, err
	}
	if s.Comparisons == nil {
		s.Comparisons = []statsComparison{}
	}
	return s, nil
}

func loadFuzzReport(path string) (fuzzReport, error) {
	var f fuzzReport
	b, err := os.ReadFile(path)
	if err != nil {
		return f, err
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return f, err
	}
	if f.Failures == nil {
		f.Failures = []fuzzFailure{}
	}
	return f, nil
}

func statsKey(c statsComparison) string {
	return strings.Join([]string{c.PairLabel, c.Track, c.Mode, c.Workload, c.Class, c.ImplA, c.ImplB}, "|")
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
