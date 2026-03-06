package main

import "testing"

func TestSelectImplSpecs(t *testing.T) {
	root := "/tmp/x"

	goImpls, err := selectImplSpecs(root, "go", "")
	if err != nil {
		t.Fatalf("select go: %v", err)
	}
	if len(goImpls) != 2 {
		t.Fatalf("go impl count=%d want 2", len(goImpls))
	}

	rustImpls, err := selectImplSpecs(root, "rust", "")
	if err != nil {
		t.Fatalf("select rust: %v", err)
	}
	if len(rustImpls) != 2 {
		t.Fatalf("rust impl count=%d want 2", len(rustImpls))
	}

	only, err := selectImplSpecs(root, "all", "schubfach-rs")
	if err != nil {
		t.Fatalf("select specific: %v", err)
	}
	if len(only) != 1 || only[0].Name != "schubfach-rs" {
		t.Fatalf("specific selection=%v want schubfach-rs", implNames(only))
	}
}

func TestParseAPIPrefix(t *testing.T) {
	impl, mode, ok := parseAPIPrefix("BenchmarkAPICanonicalizeSchubfachRs")
	if !ok {
		t.Fatal("expected parse ok for rust canonicalize")
	}
	if impl != "schubfach-rs" || mode != "canonicalize" {
		t.Fatalf("got impl=%q mode=%q", impl, mode)
	}

	impl, mode, ok = parseAPIPrefix("BenchmarkAPIVerifyJSONCanonRs")
	if !ok {
		t.Fatal("expected parse ok for rust verify")
	}
	if impl != "json-canon-rs" || mode != "verify" {
		t.Fatalf("got impl=%q mode=%q", impl, mode)
	}
}
