# jcs-bench-lab

[![DOI](https://zenodo.org/badge/doi/10.5281/zenodo.18890836.svg)](https://doi.org/10.5281/zenodo.18890836)

Go-native benchmarking lab for comparing:
- `impl-schubfach` (`lattice-substrate/jcs-schubfach`)
- `impl-json-canon` (`lattice-substrate/json-canon`)
- `impl-schubfach-rs` (`lattice-substrate/jcs-schubfach-rs`)
- `impl-json-canon-rs` (`lattice-substrate/json-canon-rs`)

All implementations are treated as black-box competitors.

## First runnable benchmark command

```bash
go run ./cmd/lab smoke
```

This single command will:
1. Build selected CLI binaries into `bin/` (default: all Go+Rust impls)
2. Capture environment metadata into `results/latest-env.txt`
3. Generate expanded workload + canonical fixture corpus into `workloads/`
4. Run authoritative conformance matrix with case IDs
5. Run repeated CLI benchmarks with randomized order
6. Emit raw/summarized/stats/report outputs into `results/`

## Main commands

```bash
go run ./cmd/lab setup
go run ./cmd/lab setup -lang go
go run ./cmd/lab setup -lang rust
go run ./cmd/lab gen-workloads
go run ./cmd/lab conformance
go run ./cmd/lab conformance -lang rust
go run ./cmd/lab conformance -impl schubfach-rs
go run ./cmd/lab bench-cli -track e2e -repeats 9 -warmup 2 -seed 42
go run ./cmd/lab bench-cli -lang go
go run ./cmd/lab bench-cli -lang rust
go run ./cmd/lab bench-cli -track cli-algorithmic -repeats 9 -warmup 2
go run ./cmd/lab bench-cli -track verify-only -repeats 9 -warmup 2
go run ./cmd/lab bench-api -count 7
go run ./cmd/lab fuzz -cases 1000 -seed 123
go run ./cmd/lab stats -alpha 0.05 -resamples 2000
go run ./cmd/lab benchstat
go run ./cmd/lab profile-api -count 1 -benchtime 1s
go run ./cmd/lab gate -fuzz results/latest-fuzz.json -baseline results/baseline-stats.json
go run ./cmd/lab report
```

Results are always written under `results/` with timestamped files and `latest-*` pointers.

## Statistical and Gate Pipeline

```bash
# assumes setup + workload generation already complete
go run ./cmd/lab conformance
go run ./cmd/lab bench-cli -track verify-only -repeats 7 -warmup 1
go run ./cmd/lab fuzz -cases 500
go run ./cmd/lab stats
go run ./cmd/lab bench-api -count 3
go run ./cmd/lab benchstat -allow-fallback=false
go run ./cmd/lab gate -fuzz results/latest-fuzz.json -baseline results/baseline-stats.json
go run ./cmd/lab report
```

This emits:
- `results/latest-conformance.json` (authoritative oracle pass/fail by case ID)
- `results/latest-cli-runs.csv`, `results/latest-cli-summary.csv`, `results/latest-quality.json`
- `results/latest-stats.json`, `results/latest-stats.md`
- `results/latest-benchstat.md`
- `results/latest-api-prof-*.txt`, `results/latest-api-prof-*-cpu.pprof`, `results/latest-api-prof-*-mem.pprof`
- `results/latest-api-profile-analysis.md`
- `results/latest-report.md` and updates `REPORT.md`
