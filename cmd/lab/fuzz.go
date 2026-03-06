package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type fuzzFailure struct {
	CaseID string `json:"case_id"`
	Input  string `json:"input"`
	Reason string `json:"reason"`
}

type fuzzReport struct {
	GeneratedAtUTC string        `json:"generated_at_utc"`
	Seed           int64         `json:"seed"`
	Cases          int           `json:"cases"`
	Failures       []fuzzFailure `json:"failures"`
}

type fuzzNode struct {
	kind    string
	obj     []fuzzMember
	arr     []fuzzNode
	str     string
	num     string
	boolean bool
}

type fuzzMember struct {
	k string
	v fuzzNode
}

func runFuzz(cases int, seed int64, lang, impl string) error {
	if cases < 1 {
		return fmt.Errorf("cases must be >= 1")
	}
	if seed == 0 {
		seed = time.Now().UTC().UnixNano()
	}
	root, err := repoRoot()
	if err != nil {
		return err
	}
	selected, err := selectImplSpecs(root, lang, impl)
	if err != nil {
		return err
	}
	if len(selected) < 2 {
		return fmt.Errorf("fuzz requires at least 2 implementations; got %d", len(selected))
	}
	if err := ensureImplBinsExist(root, selected); err != nil {
		return err
	}
	impls := toImplConfigs(root, selected)

	r := rand.New(rand.NewSource(seed))
	failures := make([]fuzzFailure, 0)
	for i := 0; i < cases; i++ {
		n := genFuzzNode(r, 0)
		input := renderFuzzNode(n, r, false)
		caseID := fmt.Sprintf("FUZZ-%06d", i+1)

		canonByImpl := make(map[string][]byte, len(impls))
		var canonRef []byte
		canonMismatch := false
		for _, impl := range impls {
			canon, err := canonicalizeOutput(impl.Bin, []byte(input))
			if err != nil {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "canonicalize execution error for " + impl.Name})
				canonMismatch = true
				break
			}
			canonByImpl[impl.Name] = canon
			if canonRef == nil {
				canonRef = canon
				continue
			}
			if !bytes.Equal(canonRef, canon) {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "cross-implementation output mismatch"})
				canonMismatch = true
				break
			}
		}
		if canonMismatch {
			continue
		}

		// Idempotence: canonicalize(canonicalize(x)) == canonicalize(x)
		idempotenceFailed := false
		for _, impl := range impls {
			canon := canonByImpl[impl.Name]
			canonAgain, err := canonicalizeOutput(impl.Bin, canon)
			if err != nil {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "idempotence canonicalize second pass failed for " + impl.Name})
				idempotenceFailed = true
				break
			}
			if !bytes.Equal(canonAgain, canon) {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "idempotence mismatch for " + impl.Name})
				idempotenceFailed = true
				break
			}
		}
		if idempotenceFailed {
			continue
		}

		verifyFailed := false
		for _, impl := range impls {
			verifyRes, err := runOne(impl.Bin, "verify", benchInput{Data: canonRef}, -1)
			if err != nil || !verifyRes.OK {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "verify rejected canonicalized output for " + impl.Name})
				verifyFailed = true
				break
			}
		}
		if verifyFailed {
			continue
		}

		// Metamorphic check: canonical output with added leading WS must fail verify.
		mutated := append([]byte(" "), canonRef...)
		mutatedFailed := false
		for _, impl := range impls {
			verifyMut, err := runOne(impl.Bin, "verify", benchInput{Data: mutated}, -1)
			if err != nil {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "metamorphic verify execution error for " + impl.Name})
				mutatedFailed = true
				break
			}
			if verifyMut.OK {
				failures = append(failures, fuzzFailure{CaseID: caseID, Input: input, Reason: "metamorphic verify accepted non-canonical mutation for " + impl.Name})
				mutatedFailed = true
				break
			}
		}
		if mutatedFailed {
			continue
		}
	}

	report := fuzzReport{
		GeneratedAtUTC: time.Now().UTC().Format(time.RFC3339),
		Seed:           seed,
		Cases:          cases,
		Failures:       failures,
	}
	stamp := time.Now().UTC().Format("20060102T150405Z")
	outPath := filepath.Join(root, "results", "fuzz-"+stamp+".json")
	latestPath := filepath.Join(root, "results", "latest-fuzz.json")
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
	fmt.Printf("fuzz complete\n- %s\n- %s\n", outPath, latestPath)
	fmt.Printf("summary: cases=%d failures=%d\n", report.Cases, len(report.Failures))
	if len(report.Failures) > 0 {
		return fmt.Errorf("fuzz found %d failures", len(report.Failures))
	}
	return nil
}

func canonicalizeOutput(bin string, input []byte) ([]byte, error) {
	cmd := exec.Command(bin, "canonicalize", "-")
	cmd.Stdin = bytes.NewReader(input)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("canonicalize failed: %w stderr=%s", err, strings.TrimSpace(stderr.String()))
	}
	return out, nil
}

func genFuzzNode(r *rand.Rand, depth int) fuzzNode {
	if depth > 4 {
		return genLeaf(r)
	}
	switch r.Intn(6) {
	case 0:
		return genLeaf(r)
	case 1:
		return genLeaf(r)
	case 2:
		return genArray(r, depth)
	case 3:
		return genObject(r, depth)
	case 4:
		return genLeaf(r)
	default:
		return genLeaf(r)
	}
}

func genLeaf(r *rand.Rand) fuzzNode {
	switch r.Intn(5) {
	case 0:
		return fuzzNode{kind: "null"}
	case 1:
		return fuzzNode{kind: "bool", boolean: r.Intn(2) == 0}
	case 2:
		return fuzzNode{kind: "string", str: randString(r)}
	case 3:
		return fuzzNode{kind: "number", num: randNumber(r)}
	default:
		return fuzzNode{kind: "string", str: randUnicodeString(r)}
	}
}

func genArray(r *rand.Rand, depth int) fuzzNode {
	n := r.Intn(5)
	arr := make([]fuzzNode, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, genFuzzNode(r, depth+1))
	}
	return fuzzNode{kind: "array", arr: arr}
}

func genObject(r *rand.Rand, depth int) fuzzNode {
	n := r.Intn(5)
	obj := make([]fuzzMember, 0, n)
	seen := map[string]bool{}
	for len(obj) < n {
		k := randKey(r)
		if seen[k] {
			continue
		}
		seen[k] = true
		obj = append(obj, fuzzMember{k: k, v: genFuzzNode(r, depth+1)})
	}
	return fuzzNode{kind: "object", obj: obj}
}

func renderFuzzNode(n fuzzNode, r *rand.Rand, canonical bool) string {
	switch n.kind {
	case "null":
		return "null"
	case "bool":
		if n.boolean {
			return "true"
		}
		return "false"
	case "string":
		b, _ := json.Marshal(n.str)
		return string(b)
	case "number":
		return n.num
	case "array":
		parts := make([]string, len(n.arr))
		for i := range n.arr {
			parts[i] = renderFuzzNode(n.arr[i], r, canonical)
		}
		sep := ","
		if !canonical && r.Intn(2) == 0 {
			sep = ", "
		}
		return "[" + strings.Join(parts, sep) + "]"
	case "object":
		members := append([]fuzzMember(nil), n.obj...)
		if canonical {
			sort.Slice(members, func(i, j int) bool { return members[i].k < members[j].k })
		} else {
			r.Shuffle(len(members), func(i, j int) { members[i], members[j] = members[j], members[i] })
		}
		parts := make([]string, len(members))
		for i, m := range members {
			kb, _ := json.Marshal(m.k)
			colon := ":"
			if !canonical && r.Intn(2) == 0 {
				colon = " : "
			}
			parts[i] = string(kb) + colon + renderFuzzNode(m.v, r, canonical)
		}
		sep := ","
		if !canonical && r.Intn(2) == 0 {
			sep = ", "
		}
		return "{" + strings.Join(parts, sep) + "}"
	default:
		return "null"
	}
}

func randString(r *rand.Rand) string {
	n := 1 + r.Intn(16)
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(byte('a' + r.Intn(26)))
	}
	return b.String()
}

func randUnicodeString(r *rand.Rand) string {
	pool := []rune{'Ω', 'é', '日', '😀', 'ß', '中', '£'}
	n := 1 + r.Intn(8)
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(pool[r.Intn(len(pool))])
	}
	return b.String()
}

func randKey(r *rand.Rand) string {
	return "k" + strconv.Itoa(r.Intn(32))
}

func randNumber(r *rand.Rand) string {
	switch r.Intn(5) {
	case 0:
		return strconv.Itoa(r.Intn(1000000) - 500000)
	case 1:
		return fmt.Sprintf("%d.%d", r.Intn(1000), r.Intn(1000000))
	case 2:
		return fmt.Sprintf("%de+%d", 1+r.Intn(9), r.Intn(20))
	case 3:
		return fmt.Sprintf("%de-%d", 1+r.Intn(9), 1+r.Intn(20))
	default:
		return fmt.Sprintf("%d", 1+r.Intn(9007199254740990))
	}
}
