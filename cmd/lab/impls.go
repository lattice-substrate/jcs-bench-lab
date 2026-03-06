package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type implSpec struct {
	Name         string
	Lang         string
	RepoDir      string
	BinName      string
	BuildKind    string // go | rust
	GoBuildPkg   string
	APIBenchStem string
}

func allImplSpecs(root string) []implSpec {
	return []implSpec{
		{
			Name:         "schubfach",
			Lang:         "go",
			RepoDir:      filepath.Join(root, "impl-schubfach"),
			BinName:      "schubfach-jcs-canon",
			BuildKind:    "go",
			GoBuildPkg:   "./cmd/jcs-canon",
			APIBenchStem: "Schubfach",
		},
		{
			Name:         "json-canon",
			Lang:         "go",
			RepoDir:      filepath.Join(root, "impl-json-canon"),
			BinName:      "json-canon-jcs-canon",
			BuildKind:    "go",
			GoBuildPkg:   "./cmd/jcs-canon",
			APIBenchStem: "JSONCanon",
		},
		{
			Name:         "schubfach-rs",
			Lang:         "rust",
			RepoDir:      filepath.Join(root, "impl-schubfach-rs"),
			BinName:      "schubfach-rs-jcs-canon",
			BuildKind:    "rust",
			APIBenchStem: "SchubfachRs",
		},
		{
			Name:         "json-canon-rs",
			Lang:         "rust",
			RepoDir:      filepath.Join(root, "impl-json-canon-rs"),
			BinName:      "json-canon-rs-jcs-canon",
			BuildKind:    "rust",
			APIBenchStem: "JSONCanonRs",
		},
		{
			Name:         "ryu",
			Lang:         "go",
			RepoDir:      filepath.Join(root, "impl-ryu"),
			BinName:      "ryu-jcs-canon",
			BuildKind:    "go",
			GoBuildPkg:   "./cmd/jcs-canon",
			APIBenchStem: "Ryu",
		},
		{
			Name:         "dragonbox",
			Lang:         "go",
			RepoDir:      filepath.Join(root, "impl-dragonbox"),
			BinName:      "dragonbox-jcs-canon",
			BuildKind:    "go",
			GoBuildPkg:   "./cmd/jcs-canon",
			APIBenchStem: "Dragonbox",
		},
		{
			Name:         "ryu-rs",
			Lang:         "rust",
			RepoDir:      filepath.Join(root, "impl-ryu-rs"),
			BinName:      "ryu-rs-jcs-canon",
			BuildKind:    "rust",
			APIBenchStem: "RyuRs",
		},
		{
			Name:         "dragonbox-rs",
			Lang:         "rust",
			RepoDir:      filepath.Join(root, "impl-dragonbox-rs"),
			BinName:      "dragonbox-rs-jcs-canon",
			BuildKind:    "rust",
			APIBenchStem: "DragonboxRs",
		},
	}
}

func selectImplSpecs(root, langFilter, implFilter string) ([]implSpec, error) {
	lang := strings.TrimSpace(strings.ToLower(langFilter))
	if lang == "" {
		lang = "all"
	}
	switch lang {
	case "all", "go", "rust":
	default:
		return nil, fmt.Errorf("invalid --lang %q (want all|go|rust)", langFilter)
	}

	specs := allImplSpecs(root)
	filtered := make([]implSpec, 0, len(specs))
	for _, spec := range specs {
		if lang != "all" && spec.Lang != lang {
			continue
		}
		filtered = append(filtered, spec)
	}

	implFilter = strings.TrimSpace(implFilter)
	if implFilter == "" {
		if len(filtered) == 0 {
			return nil, fmt.Errorf("no implementations selected (lang=%s)", lang)
		}
		return filtered, nil
	}

	wanted := map[string]bool{}
	for _, part := range strings.Split(implFilter, ",") {
		p := strings.TrimSpace(part)
		if p == "" {
			continue
		}
		wanted[p] = true
	}
	if len(wanted) == 0 {
		return nil, fmt.Errorf("empty --impl filter %q", implFilter)
	}

	selected := make([]implSpec, 0, len(wanted))
	for _, spec := range filtered {
		if wanted[spec.Name] {
			selected = append(selected, spec)
		}
	}
	if len(selected) == 0 {
		known := implNames(filtered)
		if len(known) == 0 {
			known = implNames(specs)
		}
		return nil, fmt.Errorf("--impl %q did not match selected language set; available: %s", implFilter, strings.Join(known, ", "))
	}
	return selected, nil
}

func implNames(specs []implSpec) []string {
	out := make([]string, 0, len(specs))
	for _, spec := range specs {
		out = append(out, spec.Name)
	}
	sort.Strings(out)
	return out
}

func implBinPath(root string, spec implSpec) string {
	return filepath.Join(root, "bin", spec.BinName)
}

func rustBenchPath(spec implSpec) string {
	return filepath.Join(spec.RepoDir, "target", "release", "bench")
}

func toImplConfigs(root string, specs []implSpec) []implConfig {
	out := make([]implConfig, 0, len(specs))
	for _, spec := range specs {
		out = append(out, implConfig{Name: spec.Name, Bin: implBinPath(root, spec)})
	}
	return out
}

func ensureImplBinsExist(root string, specs []implSpec) error {
	for _, spec := range specs {
		if _, err := os.Stat(implBinPath(root, spec)); err != nil {
			return fmt.Errorf("missing binary %s (run setup first)", implBinPath(root, spec))
		}
	}
	return nil
}

type implPair struct {
	A     string
	B     string
	Label string
}

func comparisonPairs() []implPair {
	return []implPair{
		// Within-language algorithm pairs (Go)
		{A: "schubfach", B: "json-canon", Label: "go:schubfach-vs-bd"},
		{A: "schubfach", B: "ryu", Label: "go:schubfach-vs-ryu"},
		{A: "schubfach", B: "dragonbox", Label: "go:schubfach-vs-dragonbox"},
		{A: "ryu", B: "dragonbox", Label: "go:ryu-vs-dragonbox"},
		{A: "ryu", B: "json-canon", Label: "go:ryu-vs-bd"},
		{A: "dragonbox", B: "json-canon", Label: "go:dragonbox-vs-bd"},
		// Within-language algorithm pairs (Rust)
		{A: "schubfach-rs", B: "json-canon-rs", Label: "rs:schubfach-vs-bd"},
		{A: "schubfach-rs", B: "ryu-rs", Label: "rs:schubfach-vs-ryu"},
		{A: "schubfach-rs", B: "dragonbox-rs", Label: "rs:schubfach-vs-dragonbox"},
		{A: "ryu-rs", B: "dragonbox-rs", Label: "rs:ryu-vs-dragonbox"},
		{A: "ryu-rs", B: "json-canon-rs", Label: "rs:ryu-vs-bd"},
		{A: "dragonbox-rs", B: "json-canon-rs", Label: "rs:dragonbox-vs-bd"},
		// Cross-language same-algorithm pairs
		{A: "schubfach", B: "schubfach-rs", Label: "schubfach:go-vs-rs"},
		{A: "json-canon", B: "json-canon-rs", Label: "bd:go-vs-rs"},
		{A: "ryu", B: "ryu-rs", Label: "ryu:go-vs-rs"},
		{A: "dragonbox", B: "dragonbox-rs", Label: "dragonbox:go-vs-rs"},
	}
}

func parseAPIPrefix(prefix string) (impl, mode string, ok bool) {
	mode = ""
	stem := ""
	switch {
	case strings.HasPrefix(prefix, "BenchmarkAPICanonicalize"):
		mode = "canonicalize"
		stem = strings.TrimPrefix(prefix, "BenchmarkAPICanonicalize")
	case strings.HasPrefix(prefix, "BenchmarkAPIVerify"):
		mode = "verify"
		stem = strings.TrimPrefix(prefix, "BenchmarkAPIVerify")
	default:
		return "", "", false
	}

	switch stem {
	case "Schubfach":
		return "schubfach", mode, true
	case "JSONCanon":
		return "json-canon", mode, true
	case "SchubfachRs":
		return "schubfach-rs", mode, true
	case "JSONCanonRs":
		return "json-canon-rs", mode, true
	case "Ryu":
		return "ryu", mode, true
	case "Dragonbox":
		return "dragonbox", mode, true
	case "RyuRs":
		return "ryu-rs", mode, true
	case "DragonboxRs":
		return "dragonbox-rs", mode, true
	default:
		return "", "", false
	}
}

func findImplSpecByName(root, name string) (implSpec, bool) {
	for _, spec := range allImplSpecs(root) {
		if spec.Name == name {
			return spec, true
		}
	}
	return implSpec{}, false
}
