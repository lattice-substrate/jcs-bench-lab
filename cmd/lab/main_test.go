package main

import (
	"strings"
	"testing"
)

func TestClassifyErrExtractsStructuredClass(t *testing.T) {
	got := classifyErr("error: DUPLICATE_KEY at byte=12")
	if got != "duplicate_key" {
		t.Fatalf("classifyErr structured = %q, want duplicate_key", got)
	}

	got = classifyErr("line 1 column 4: invalid token")
	if got != "parse" {
		t.Fatalf("classifyErr parse = %q, want parse", got)
	}
}

func TestApplyQualityChecksInvalidParityAllRuns(t *testing.T) {
	recs := []runRecord{
		{Track: "e2e", Impl: "schubfach", Mode: "verify", Workload: "x", Class: "invalid", SessionID: 0, Run: 0, OK: false, ErrorClass: "parse", PassesOracle: true, CaseID: "A"},
		{Track: "e2e", Impl: "json-canon", Mode: "verify", Workload: "x", Class: "invalid", SessionID: 0, Run: 0, OK: false, ErrorClass: "parse", PassesOracle: true, CaseID: "A"},
		{Track: "e2e", Impl: "schubfach", Mode: "verify", Workload: "x", Class: "invalid", SessionID: 1, Run: 1, OK: false, ErrorClass: "parse", PassesOracle: true, CaseID: "A"},
		{Track: "e2e", Impl: "json-canon", Mode: "verify", Workload: "x", Class: "invalid", SessionID: 1, Run: 1, OK: true, ErrorClass: "none", PassesOracle: false, CaseID: "A"},
	}
	q := qualityReport{}
	applyQualityChecks(recs, &q)
	if len(q.InvalidFailureParityIssues) == 0 {
		t.Fatal("expected invalid parity issue from run/session mismatch")
	}
	if len(q.OracleMismatches) == 0 {
		t.Fatal("expected oracle mismatch to be reported")
	}
}

func TestDefaultWorkloadsExpanded(t *testing.T) {
	valid := defaultValidFixtures()
	invalid := defaultInvalidFixtures()
	if len(valid) < 12 {
		t.Fatalf("valid fixtures too small: got %d", len(valid))
	}
	if len(invalid) < 12 {
		t.Fatalf("invalid fixtures too small: got %d", len(invalid))
	}
	for name, v := range valid {
		if len(v.Canonical) == 0 || len(v.Input) == 0 {
			t.Fatalf("fixture %s missing input/canonical", name)
		}
	}
}

func TestParseAPIBenchmarkLineSupportsVerify(t *testing.T) {
	line := "BenchmarkAPIVerifySchubfach/canonical/small-16 1000 1234 ns/op 42.00 MB/s 50 B/op 2 allocs/op"
	s, ok := parseAPIBenchmarkLine(line)
	if !ok {
		t.Fatal("expected parse success")
	}
	if s.Operation != "verify" {
		t.Fatalf("operation=%q want verify", s.Operation)
	}
	if s.Impl != "schubfach" {
		t.Fatalf("impl=%q want schubfach", s.Impl)
	}
	if !strings.Contains(s.Workload, "canonical/small") {
		t.Fatalf("workload=%q missing expected segments", s.Workload)
	}
}

func TestRecommendationPolicy(t *testing.T) {
	q := qualityReport{}
	conf := conformanceReport{FailureCount: 1}
	stats := statsReport{}
	fuzz := fuzzReport{}
	msg := recommendation(q, conf, stats, fuzz)
	if !strings.Contains(strings.ToLower(msg), "conformance") {
		t.Fatalf("unexpected message for conformance failure: %q", msg)
	}

	conf.FailureCount = 0
	q.OracleMismatches = []string{"x"}
	msg = recommendation(q, conf, stats, fuzz)
	if !strings.Contains(strings.ToLower(msg), "quality") {
		t.Fatalf("unexpected message for quality failure: %q", msg)
	}

	q.OracleMismatches = nil
	stats.Comparisons = []statsComparison{{Winner: "schubfach", Significant: true, SignificantBH: true}}
	msg = recommendation(q, conf, stats, fuzz)
	if !strings.Contains(msg, "schubfach") {
		t.Fatalf("expected schubfach recommendation, got %q", msg)
	}
}
