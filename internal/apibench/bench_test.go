package apibench

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	schubfach "github.com/lattice-substrate/jcs-schubfach/jcs"
	jsoncanon "github.com/lattice-substrate/json-canon/jcs"
)

type workload struct {
	Name string
	Data []byte
}

func loadWorkloadsFromDir(b *testing.B, rel string) []workload {
	b.Helper()
	root := findRepoRoot(b)
	dir := filepath.Join(root, rel)
	entries, err := os.ReadDir(dir)
	if err != nil {
		b.Fatalf("read workloads: %v", err)
	}
	out := make([]workload, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		p := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(p)
		if err != nil {
			b.Fatalf("read %s: %v", p, err)
		}
		out = append(out, workload{Name: strings.TrimSuffix(e.Name(), ".json"), Data: data})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	if len(out) == 0 {
		b.Fatalf("no workloads found in %s; run `go run ./cmd/lab gen-workloads` first", rel)
	}
	return out
}

func loadValidWorkloads(b *testing.B) []workload {
	return loadWorkloadsFromDir(b, filepath.Join("workloads", "valid"))
}

func loadCanonicalWorkloads(b *testing.B) []workload {
	return loadWorkloadsFromDir(b, filepath.Join("workloads", "canonical"))
}

func loadNonCanonicalWorkloads(b *testing.B) []workload {
	valid := loadValidWorkloads(b)
	canonical := loadCanonicalWorkloads(b)
	byName := make(map[string][]byte, len(canonical))
	for _, c := range canonical {
		byName[c.Name] = c.Data
	}
	out := make([]workload, 0, len(valid))
	for _, v := range valid {
		c, ok := byName[v.Name]
		if !ok {
			continue
		}
		if !bytes.Equal(v.Data, c) {
			out = append(out, v)
		}
	}
	if len(out) == 0 {
		b.Fatal("no noncanonical workloads found")
	}
	return out
}

func findRepoRoot(b *testing.B) string {
	b.Helper()
	wd, err := os.Getwd()
	if err != nil {
		b.Fatalf("getwd: %v", err)
	}
	cur := wd
	for i := 0; i < 8; i++ {
		if _, err := os.Stat(filepath.Join(cur, "go.mod")); err == nil {
			if _, err := os.Stat(filepath.Join(cur, "cmd", "lab", "main.go")); err == nil {
				return cur
			}
		}
		next := filepath.Dir(cur)
		if next == cur {
			break
		}
		cur = next
	}
	b.Fatalf("could not locate lab root from %s", wd)
	return ""
}

func BenchmarkAPICanonicalizeSchubfach(b *testing.B) {
	workloads := loadValidWorkloads(b)
	for _, w := range workloads {
		w := w
		b.Run(w.Name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(w.Data)))
			for i := 0; i < b.N; i++ {
				out, err := schubfach.Canonicalize(w.Data)
				if err != nil {
					b.Fatalf("canonicalize: %v", err)
				}
				if len(out) == 0 {
					b.Fatal("empty output")
				}
			}
		})
	}
}

func BenchmarkAPICanonicalizeJSONCanon(b *testing.B) {
	workloads := loadValidWorkloads(b)
	for _, w := range workloads {
		w := w
		b.Run(w.Name, func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(w.Data)))
			for i := 0; i < b.N; i++ {
				out, err := jsoncanon.Canonicalize(w.Data)
				if err != nil {
					b.Fatalf("canonicalize: %v", err)
				}
				if len(out) == 0 {
					b.Fatal("empty output")
				}
			}
		})
	}
}

func verifyViaCanonicalize(fn func([]byte) ([]byte, error), input []byte) (bool, error) {
	out, err := fn(input)
	if err != nil {
		return false, err
	}
	return bytes.Equal(out, input), nil
}

func BenchmarkAPIVerifySchubfach(b *testing.B) {
	canonical := loadCanonicalWorkloads(b)
	noncanonical := loadNonCanonicalWorkloads(b)

	b.Run("canonical", func(b *testing.B) {
		for _, w := range canonical {
			w := w
			b.Run(w.Name, func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(int64(len(w.Data)))
				for i := 0; i < b.N; i++ {
					ok, err := verifyViaCanonicalize(schubfach.Canonicalize, w.Data)
					if err != nil {
						b.Fatalf("verify canonicalize: %v", err)
					}
					if !ok {
						b.Fatal("canonical input unexpectedly failed verify")
					}
				}
			})
		}
	})

	b.Run("noncanonical", func(b *testing.B) {
		for _, w := range noncanonical {
			w := w
			b.Run(w.Name, func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(int64(len(w.Data)))
				for i := 0; i < b.N; i++ {
					ok, err := verifyViaCanonicalize(schubfach.Canonicalize, w.Data)
					if err != nil {
						b.Fatalf("verify canonicalize: %v", err)
					}
					if ok {
						b.Fatal("noncanonical input unexpectedly passed verify")
					}
				}
			})
		}
	})
}

func BenchmarkAPIVerifyJSONCanon(b *testing.B) {
	canonical := loadCanonicalWorkloads(b)
	noncanonical := loadNonCanonicalWorkloads(b)

	b.Run("canonical", func(b *testing.B) {
		for _, w := range canonical {
			w := w
			b.Run(w.Name, func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(int64(len(w.Data)))
				for i := 0; i < b.N; i++ {
					ok, err := verifyViaCanonicalize(jsoncanon.Canonicalize, w.Data)
					if err != nil {
						b.Fatalf("verify canonicalize: %v", err)
					}
					if !ok {
						b.Fatal("canonical input unexpectedly failed verify")
					}
				}
			})
		}
	})

	b.Run("noncanonical", func(b *testing.B) {
		for _, w := range noncanonical {
			w := w
			b.Run(w.Name, func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(int64(len(w.Data)))
				for i := 0; i < b.N; i++ {
					ok, err := verifyViaCanonicalize(jsoncanon.Canonicalize, w.Data)
					if err != nil {
						b.Fatalf("verify canonicalize: %v", err)
					}
					if ok {
						b.Fatal("noncanonical input unexpectedly passed verify")
					}
				}
			})
		}
	})
}
