package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type conformanceCase struct {
	ID                 string   `json:"id"`
	Source             string   `json:"source"`
	Requirement        string   `json:"requirement"`
	Mode               string   `json:"mode"`
	Class              string   `json:"class"`
	InputPath          string   `json:"input_path"`
	ExpectedOK         bool     `json:"expected_ok"`
	ExpectedSHA256     string   `json:"expected_sha256,omitempty"`
	ExpectedErrorClass string   `json:"expected_error_class,omitempty"`
	Tags               []string `json:"tags,omitempty"`
}

type conformanceCaseResult struct {
	CaseID             string   `json:"case_id"`
	Source             string   `json:"source"`
	Requirement        string   `json:"requirement"`
	Impl               string   `json:"impl"`
	Mode               string   `json:"mode"`
	Class              string   `json:"class"`
	InputPath          string   `json:"input_path"`
	ExpectedOK         bool     `json:"expected_ok"`
	ActualOK           bool     `json:"actual_ok"`
	ExpectedSHA256     string   `json:"expected_sha256,omitempty"`
	ActualSHA256       string   `json:"actual_sha256,omitempty"`
	ExpectedErrorClass string   `json:"expected_error_class,omitempty"`
	ActualErrorClass   string   `json:"actual_error_class,omitempty"`
	ExitCode           int      `json:"exit_code"`
	DurationNS         int64    `json:"duration_ns"`
	Pass               bool     `json:"pass"`
	FailureReason      string   `json:"failure_reason,omitempty"`
	Tags               []string `json:"tags,omitempty"`
}

type conformanceReport struct {
	GeneratedAtUTC string                  `json:"generated_at_utc"`
	CaseCount      int                     `json:"case_count"`
	FailureCount   int                     `json:"failure_count"`
	Cases          []conformanceCaseResult `json:"cases"`
	SummaryByImpl  map[string]caseSummary  `json:"summary_by_impl"`
	SummaryBySrc   map[string]caseSummary  `json:"summary_by_source"`
}

type caseSummary struct {
	Total    int `json:"total"`
	Passed   int `json:"passed"`
	Failed   int `json:"failed"`
	PassRate int `json:"pass_rate_pct"`
}

func runConformance(failOnMismatch bool, lang, impl string) error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	if err := ensureDirs(root); err != nil {
		return err
	}

	selected, err := selectImplSpecs(root, lang, impl)
	if err != nil {
		return err
	}
	if err := ensureImplBinsExist(root, selected); err != nil {
		return err
	}
	impls := toImplConfigs(root, selected)

	cases, err := loadConformanceCases(root)
	if err != nil {
		return err
	}

	results := make([]conformanceCaseResult, 0, len(cases)*len(impls))
	for _, c := range cases {
		for _, impl := range impls {
			run, err := runOne(impl.Bin, c.Mode, benchInput{Path: c.InputPath}, -1)
			if err != nil {
				return err
			}
			actualErrClass := classifyErr(run.Stderr)
			pass := run.OK == c.ExpectedOK
			reason := ""
			if !pass {
				reason = "unexpected success/failure outcome"
			}
			if pass && c.ExpectedSHA256 != "" && run.OK && run.OutputSHA256 != c.ExpectedSHA256 {
				pass = false
				reason = "output hash mismatch"
			}
			if pass && !c.ExpectedOK && c.ExpectedErrorClass != "" && actualErrClass != c.ExpectedErrorClass {
				pass = false
				reason = "error class mismatch"
			}
			results = append(results, conformanceCaseResult{
				CaseID:             c.ID,
				Source:             c.Source,
				Requirement:        c.Requirement,
				Impl:               impl.Name,
				Mode:               c.Mode,
				Class:              c.Class,
				InputPath:          c.InputPath,
				ExpectedOK:         c.ExpectedOK,
				ActualOK:           run.OK,
				ExpectedSHA256:     c.ExpectedSHA256,
				ActualSHA256:       run.OutputSHA256,
				ExpectedErrorClass: c.ExpectedErrorClass,
				ActualErrorClass:   actualErrClass,
				ExitCode:           run.ExitCode,
				DurationNS:         run.DurationNS,
				Pass:               pass,
				FailureReason:      reason,
				Tags:               append([]string(nil), c.Tags...),
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].CaseID == results[j].CaseID {
			return results[i].Impl < results[j].Impl
		}
		return results[i].CaseID < results[j].CaseID
	})

	report := conformanceReport{
		GeneratedAtUTC: time.Now().UTC().Format(time.RFC3339),
		CaseCount:      len(results),
		Cases:          results,
		SummaryByImpl:  map[string]caseSummary{},
		SummaryBySrc:   map[string]caseSummary{},
	}
	for _, r := range results {
		if !r.Pass {
			report.FailureCount++
		}
		accImpl := report.SummaryByImpl[r.Impl]
		accImpl.Total++
		if r.Pass {
			accImpl.Passed++
		} else {
			accImpl.Failed++
		}
		if accImpl.Total > 0 {
			accImpl.PassRate = int(float64(accImpl.Passed) * 100 / float64(accImpl.Total))
		}
		report.SummaryByImpl[r.Impl] = accImpl

		accSrc := report.SummaryBySrc[r.Source]
		accSrc.Total++
		if r.Pass {
			accSrc.Passed++
		} else {
			accSrc.Failed++
		}
		if accSrc.Total > 0 {
			accSrc.PassRate = int(float64(accSrc.Passed) * 100 / float64(accSrc.Total))
		}
		report.SummaryBySrc[r.Source] = accSrc
	}

	stamp := time.Now().UTC().Format("20060102T150405Z")
	outPath := filepath.Join(root, "results", "conformance-"+stamp+".json")
	latestPath := filepath.Join(root, "results", "latest-conformance.json")
	b, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(outPath, b, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latestPath, b, 0o644); err != nil {
		return err
	}

	fmt.Printf("conformance complete\n- %s\n- %s\n", outPath, latestPath)
	fmt.Printf("summary: total=%d failed=%d\n", report.CaseCount, report.FailureCount)
	if failOnMismatch && report.FailureCount > 0 {
		return fmt.Errorf("conformance failed: %d failing cases", report.FailureCount)
	}
	return nil
}

func loadConformanceCases(root string) ([]conformanceCase, error) {
	cases := make([]conformanceCase, 0, 256)

	cyberInDir := filepath.Join(root, "impl-schubfach", "conformance", "official", "cyberphone", "input")
	entries, err := os.ReadDir(cyberInDir)
	if err != nil {
		return nil, fmt.Errorf("read cyberphone fixtures: %w", err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ".json")
		inPath := filepath.Join(cyberInDir, e.Name())
		outPath := filepath.Join(root, "impl-schubfach", "conformance", "official", "cyberphone", "output", e.Name())
		h, err := fileSHA256(outPath)
		if err != nil {
			return nil, err
		}
		cases = append(cases, conformanceCase{
			ID:             "OFFICIAL-CYBERPHONE-CANON-" + strings.ToUpper(strings.ReplaceAll(name, "-", "_")),
			Source:         "cyberphone",
			Requirement:    "RFC8785-CANON",
			Mode:           "canonicalize",
			Class:          "valid",
			InputPath:      inPath,
			ExpectedOK:     true,
			ExpectedSHA256: h,
			Tags:           []string{"official", "canonicalize"},
		})
		cases = append(cases, conformanceCase{
			ID:             "OFFICIAL-CYBERPHONE-VERIFY-PASS-" + strings.ToUpper(strings.ReplaceAll(name, "-", "_")),
			Source:         "cyberphone",
			Requirement:    "RFC8785-VERIFY",
			Mode:           "verify",
			Class:          "valid",
			InputPath:      outPath,
			ExpectedOK:     true,
			ExpectedSHA256: hashEmptyOutput(),
			Tags:           []string{"official", "verify"},
		})
		ib, err := os.ReadFile(inPath)
		if err != nil {
			return nil, err
		}
		ob, err := os.ReadFile(outPath)
		if err != nil {
			return nil, err
		}
		if !bytesEqual(ib, ob) {
			cases = append(cases, conformanceCase{
				ID:                 "OFFICIAL-CYBERPHONE-VERIFY-REJECT-" + strings.ToUpper(strings.ReplaceAll(name, "-", "_")),
				Source:             "cyberphone",
				Requirement:        "RFC8785-VERIFY",
				Mode:               "verify",
				Class:              "invalid",
				InputPath:          inPath,
				ExpectedOK:         false,
				ExpectedErrorClass: "not_canonical",
				Tags:               []string{"official", "verify", "reject"},
			})
		}
	}

	rfcIn := filepath.Join(root, "impl-schubfach", "conformance", "official", "rfc8785", "key_sorting_input.json")
	rfcCanonical := []byte("{\"\\r\":\"Carriage Return\",\"1\":\"One\",\"\u0080\":\"Control\",\"ö\":\"Latin Small Letter O With Diaeresis\",\"€\":\"Euro Sign\",\"😀\":\"Emoji: Grinning Face\",\"דּ\":\"Hebrew Letter Dalet With Dagesh\"}")
	h := sha256.Sum256(rfcCanonical)
	cases = append(cases, conformanceCase{
		ID:             "OFFICIAL-RFC8785-KEY-SORTING",
		Source:         "rfc8785",
		Requirement:    "RFC8785-SORT",
		Mode:           "canonicalize",
		Class:          "valid",
		InputPath:      rfcIn,
		ExpectedOK:     true,
		ExpectedSHA256: fmt.Sprintf("%x", h[:]),
		Tags:           []string{"official", "ordering"},
	})

	manifest, err := loadManifest(filepath.Join(root, "workloads", "manifest.json"))
	if err != nil {
		return nil, err
	}
	for _, w := range manifest {
		p := filepath.Join(root, filepath.FromSlash(w.Path))
		caseBase := strings.ToUpper(strings.ReplaceAll(w.Name, "-", "_"))
		if w.Class == "valid" {
			canonPath := filepath.Join(root, filepath.FromSlash(w.CanonicalPath))
			expectedSHA, err := fileSHA256(canonPath)
			if err != nil {
				return nil, err
			}
			cases = append(cases, conformanceCase{
				ID:             "LAB-WORKLOAD-CANON-" + caseBase,
				Source:         "lab-workload",
				Requirement:    "RFC8785-CANON",
				Mode:           "canonicalize",
				Class:          "valid",
				InputPath:      p,
				ExpectedOK:     true,
				ExpectedSHA256: expectedSHA,
				Tags:           append([]string{"lab", "canonicalize"}, w.Tags...),
			})
			cases = append(cases, conformanceCase{
				ID:                 "LAB-WORKLOAD-VERIFY-PASS-" + caseBase,
				Source:             "lab-workload",
				Requirement:        "RFC8785-VERIFY",
				Mode:               "verify",
				Class:              "valid",
				InputPath:          canonPath,
				ExpectedOK:         true,
				ExpectedSHA256:     hashEmptyOutput(),
				ExpectedErrorClass: "none",
				Tags:               append([]string{"lab", "verify"}, w.Tags...),
			})
		} else {
			cases = append(cases, conformanceCase{
				ID:                 "LAB-WORKLOAD-CANON-REJECT-" + caseBase,
				Source:             "lab-workload",
				Requirement:        "RFC8785-PARSE",
				Mode:               "canonicalize",
				Class:              "invalid",
				InputPath:          p,
				ExpectedOK:         false,
				ExpectedErrorClass: expectedErrorClassForTags(w.Tags),
				Tags:               append([]string{"lab", "canonicalize", "reject"}, w.Tags...),
			})
			cases = append(cases, conformanceCase{
				ID:                 "LAB-WORKLOAD-VERIFY-REJECT-" + caseBase,
				Source:             "lab-workload",
				Requirement:        "RFC8785-VERIFY",
				Mode:               "verify",
				Class:              "invalid",
				InputPath:          p,
				ExpectedOK:         false,
				ExpectedErrorClass: expectedErrorClassForTags(w.Tags),
				Tags:               append([]string{"lab", "verify", "reject"}, w.Tags...),
			})
		}
	}

	sort.Slice(cases, func(i, j int) bool { return cases[i].ID < cases[j].ID })
	return cases, nil
}

func hashEmptyOutput() string {
	h := sha256.Sum256(nil)
	return fmt.Sprintf("%x", h[:])
}

func expectedErrorClassForTags(tags []string) string {
	set := map[string]bool{}
	for _, t := range tags {
		set[t] = true
	}
	switch {
	case set["duplicate"]:
		return "duplicate_key"
	case set["surrogate"]:
		return "lone_surrogate"
	default:
		// Do not over-constrain generic parse categories; implementations may use
		// different but still valid failure classes for the same rejection.
		return ""
	}
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
