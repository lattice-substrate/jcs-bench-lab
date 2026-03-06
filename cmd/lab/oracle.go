package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//go:embed oracle_scripts/*.js
var oracleScripts embed.FS

type oracleEngine struct {
	Name    string // v8, spidermonkey, jsc
	Binary  string // resolved path
	Script  string // embed filename (e.g. "oracle_scripts/validate_v8.js")
	Args    []string
	NeedEOF bool // append __ORACLE_EOF__ sentinel to stdin
}

type oracleDivergence struct {
	Line     int    `json:"line"`
	Hex      string `json:"hex"`
	Expected string `json:"expected"`
	Got      string `json:"got"`
}

type oracleEngineResult struct {
	Engine      string             `json:"engine"`
	Version     string             `json:"version"`
	Total       int                `json:"total"`
	Passed      int                `json:"passed"`
	Failed      int                `json:"failed"`
	Divergences []oracleDivergence `json:"divergences"`
}

type oracleFileResult struct {
	File    string             `json:"file"`
	Engines []oracleEngineResult `json:"engines"`
}

type oracleReport struct {
	GeneratedAtUTC      string                        `json:"generated_at_utc"`
	VectorFiles         []string                      `json:"vector_files"`
	TotalVectors        int                           `json:"total_vectors"`
	Engines             map[string]oracleEngineResult `json:"engines"`
	CrossEngineAgreement bool                         `json:"cross_engine_agreement"`
}

func discoverEngines(v8Bin, smBin, jscBin string) []oracleEngine {
	var engines []oracleEngine

	// V8 via Node.js
	if bin := resolveEngine(v8Bin, []string{"node"}); bin != "" {
		engines = append(engines, oracleEngine{
			Name:   "v8",
			Binary: bin,
			Script: "oracle_scripts/validate_v8.js",
		})
	}

	// SpiderMonkey
	if bin := resolveEngine(smBin, []string{"js", "js128", "js115", "js102", "spidermonkey"}); bin != "" {
		engines = append(engines, oracleEngine{
			Name:   "spidermonkey",
			Binary: bin,
			Script: "oracle_scripts/validate_spidermonkey.js",
		})
	}

	// JavaScriptCore
	if bin := resolveJSC(jscBin); bin != "" {
		engines = append(engines, oracleEngine{
			Name:    "jsc",
			Binary:  bin,
			Script:  "oracle_scripts/validate_jsc.js",
			NeedEOF: true,
		})
	}

	return engines
}

func resolveEngine(explicit string, candidates []string) string {
	if explicit != "" {
		if p, err := exec.LookPath(explicit); err == nil {
			return p
		}
		return ""
	}
	for _, name := range candidates {
		if p, err := exec.LookPath(name); err == nil {
			return p
		}
	}
	return ""
}

func resolveJSC(explicit string) string {
	if explicit != "" {
		if p, err := exec.LookPath(explicit); err == nil {
			return p
		}
		return ""
	}
	if p, err := exec.LookPath("jsc"); err == nil {
		return p
	}
	// Search common webkit paths
	matches, _ := filepath.Glob("/usr/lib/*/webkit*/jsc")
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func getEngineVersion(engine oracleEngine) string {
	switch engine.Name {
	case "v8":
		out, err := exec.Command(engine.Binary, "-e", "console.log(process.version + ' V8=' + process.versions.v8)").Output()
		if err != nil {
			return "unknown"
		}
		return strings.TrimSpace(string(out))
	case "spidermonkey":
		out, err := exec.Command(engine.Binary, "--version").CombinedOutput()
		if err != nil {
			return "unknown"
		}
		return strings.TrimSpace(string(out))
	case "jsc":
		// JSC doesn't have --version; use package version
		out, _ := exec.Command("dpkg-query", "-W", "-f=${Version}", "libjavascriptcoregtk-bin").Output()
		if len(out) > 0 {
			return "JavaScriptCore " + strings.TrimSpace(string(out))
		}
		return "JavaScriptCore"
	}
	return "unknown"
}

func validateVectorsWithEngine(engine oracleEngine, csvPath string, version string) (oracleEngineResult, error) {
	scriptData, err := oracleScripts.ReadFile(engine.Script)
	if err != nil {
		return oracleEngineResult{}, fmt.Errorf("read embedded script %s: %w", engine.Script, err)
	}

	tmpScript, err := os.CreateTemp("", "oracle-validate-*.js")
	if err != nil {
		return oracleEngineResult{}, err
	}
	defer os.Remove(tmpScript.Name())
	if _, err := tmpScript.Write(scriptData); err != nil {
		tmpScript.Close()
		return oracleEngineResult{}, err
	}
	tmpScript.Close()

	csvFile, err := os.Open(csvPath)
	if err != nil {
		return oracleEngineResult{}, fmt.Errorf("open vectors %s: %w", csvPath, err)
	}
	defer csvFile.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	args := append(engine.Args, tmpScript.Name())
	cmd := exec.CommandContext(ctx, engine.Binary, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return oracleEngineResult{}, err
	}

	if err := cmd.Start(); err != nil {
		return oracleEngineResult{}, fmt.Errorf("start %s: %w", engine.Binary, err)
	}

	if _, err := io.Copy(stdin, csvFile); err != nil {
		stdin.Close()
		cmd.Wait()
		return oracleEngineResult{}, fmt.Errorf("pipe to %s: %w", engine.Name, err)
	}
	if engine.NeedEOF {
		io.WriteString(stdin, "__ORACLE_EOF__\n")
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return oracleEngineResult{}, fmt.Errorf("%s exited with error: %w\nstderr: %s", engine.Name, err, stderr.String())
	}

	var result oracleEngineResult
	if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
		return oracleEngineResult{}, fmt.Errorf("parse %s output: %w\nraw: %s", engine.Name, err, stdout.String())
	}
	if version != "" {
		result.Version = version
	}
	return result, nil
}

func runOracleValidate(v8Bin, smBin, jscBin string, requireAll bool) error {
	root, err := repoRoot()
	if err != nil {
		return err
	}
	if err := ensureDirs(root); err != nil {
		return err
	}

	engines := discoverEngines(v8Bin, smBin, jscBin)
	fmt.Printf("discovered %d engine(s):", len(engines))
	for _, e := range engines {
		fmt.Printf(" %s(%s)", e.Name, e.Binary)
	}
	fmt.Println()

	if requireAll && len(engines) < 3 {
		found := make([]string, len(engines))
		for i, e := range engines {
			found[i] = e.Name
		}
		return fmt.Errorf("--require-all: need v8+spidermonkey+jsc, found only: %s", strings.Join(found, ", "))
	}
	if len(engines) == 0 {
		return fmt.Errorf("no JavaScript engines found; install node, js115, and/or jsc")
	}

	vectorFiles := []string{
		filepath.Join(root, "impl-schubfach", "jcsfloat", "testdata", "golden_vectors.csv"),
		filepath.Join(root, "impl-schubfach", "jcsfloat", "testdata", "golden_stress_vectors.csv"),
	}
	for _, vf := range vectorFiles {
		if _, err := os.Stat(vf); err != nil {
			return fmt.Errorf("vector file not found: %s", vf)
		}
	}

	// Pre-fetch engine versions
	versions := map[string]string{}
	for _, e := range engines {
		versions[e.Name] = getEngineVersion(e)
		fmt.Printf("  %s version: %s\n", e.Name, versions[e.Name])
	}

	// Aggregate results per engine across all vector files
	aggregated := map[string]oracleEngineResult{}
	for _, e := range engines {
		aggregated[e.Name] = oracleEngineResult{
			Engine:      e.Name,
			Version:     versions[e.Name],
			Divergences: []oracleDivergence{},
		}
	}

	for _, csvPath := range vectorFiles {
		baseName := filepath.Base(csvPath)
		for _, engine := range engines {
			fmt.Printf("validating %s with %s ... ", baseName, engine.Name)
			result, err := validateVectorsWithEngine(engine, csvPath, versions[engine.Name])
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
				if requireAll {
					return fmt.Errorf("engine %s failed on %s: %w", engine.Name, baseName, err)
				}
				continue
			}
			fmt.Printf("%d passed, %d failed\n", result.Passed, result.Failed)

			agg := aggregated[engine.Name]
			agg.Total += result.Total
			agg.Passed += result.Passed
			agg.Failed += result.Failed
			remaining := 100 - len(agg.Divergences)
			if remaining > 0 && len(result.Divergences) > 0 {
				if len(result.Divergences) > remaining {
					agg.Divergences = append(agg.Divergences, result.Divergences[:remaining]...)
				} else {
					agg.Divergences = append(agg.Divergences, result.Divergences...)
				}
			}
			aggregated[engine.Name] = agg
		}
	}

	// Build report
	totalVectors := 0
	allPass := true
	engineResults := map[string]oracleEngineResult{}
	for name, agg := range aggregated {
		if agg.Failed > 0 {
			allPass = false
		}
		if totalVectors == 0 {
			totalVectors = agg.Total
		}
		engineResults[name] = agg
	}

	fileNames := make([]string, len(vectorFiles))
	for i, vf := range vectorFiles {
		fileNames[i] = filepath.Base(vf)
	}

	report := oracleReport{
		GeneratedAtUTC:       time.Now().UTC().Format(time.RFC3339),
		VectorFiles:          fileNames,
		TotalVectors:         totalVectors,
		Engines:              engineResults,
		CrossEngineAgreement: allPass,
	}

	reportJSON, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}

	ts := time.Now().Format("20060102-150405")
	timestamped := filepath.Join(root, "results", fmt.Sprintf("oracle-validate-%s.json", ts))
	latest := filepath.Join(root, "results", "latest-oracle-validate.json")

	if err := os.WriteFile(timestamped, reportJSON, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(latest, reportJSON, 0o644); err != nil {
		return err
	}

	fmt.Printf("\nresults written to:\n  %s\n  %s\n", timestamped, latest)
	if allPass {
		fmt.Printf("\ncross-engine agreement: PASS (%d engines, %d vectors each)\n", len(engineResults), totalVectors)
	} else {
		fmt.Printf("\ncross-engine agreement: FAIL (divergences detected)\n")
		return fmt.Errorf("oracle validation failed: cross-engine divergences detected")
	}
	return nil
}
