# Reproducibility Artifact

## Prerequisites

- Go 1.24+ (`go version`)
- Node.js 23+ (`node --version`) — for oracle vector regeneration only
- Linux x86-64 (arm64 for cross-architecture verification)
- `benchstat` (optional): `go install golang.org/x/perf/cmd/benchstat@latest`
- LaTeX (optional): for paper compilation (`pdflatex`, `bibtex`)

## Quick Verification

```bash
# Build both implementations
go run ./cmd/lab setup

# Generate workloads
go run ./cmd/lab gen-workloads

# Run conformance gate (must pass before benchmarks)
go run ./cmd/lab conformance

# Verify oracle vectors match Node.js output
cd impl-schubfach/jcsfloat/testdata
node generate_golden.js 2>/dev/null | sha256sum
# Expected: 593bdecbe0dccbc182bc3baf570b716887db25739fc61b7808764ecb966d5636
node generate_stress_golden.js 2>/dev/null | sha256sum
# Expected: 287d21ac87e5665550f1baf86038302a0afc67a74a020dffb872f1a93b26d410
cd ../../..
```

## Full Benchmark Pipeline

Expected runtime: ~30-60 minutes depending on hardware.

```bash
go run ./cmd/lab setup
go run ./cmd/lab gen-workloads
go run ./cmd/lab conformance
go run ./cmd/lab bench-api -count 10
go run ./cmd/lab bench-cli -repeats 15 -warmup 3
go run ./cmd/lab fuzz -cases 2000
go run ./cmd/lab stats -resamples 5000
go run ./cmd/lab benchstat
go run ./cmd/lab profile-api
go run ./cmd/lab report
```

## Output Files

All results are written to `results/` with timestamped filenames and
`latest-*` symlinks:

| File | Description |
|------|-------------|
| `latest-env.txt` | Hardware/software environment capture |
| `latest-api-bench.txt` | Go API benchmark output (benchstat-compatible) |
| `latest-cli-runs.csv` | Raw CLI benchmark measurements |
| `latest-cli-summary.csv` | Aggregated CLI benchmark statistics |
| `latest-quality.json` | Quality checks (determinism, oracle parity) |
| `latest-conformance.json` | Conformance gate results |
| `latest-stats.json` | Statistical comparisons with p-values and CIs |
| `latest-stats.md` | Human-readable statistical summary |
| `latest-fuzz.json` | Differential fuzz results |
| `latest-benchstat.md` | benchstat comparative analysis |

## Oracle Vector Provenance

See `impl-schubfach/jcsfloat/testdata/PROVENANCE.json` for full details.

- **Oracle function**: ECMA-262 `String(x)` (§7.1.12.1) via V8
- **Verified Node.js**: v23.3.0 (V8 12.9.202.28-node.11)
- **Total vectors**: 286,362 (54,445 primary + 231,917 stress)
- **Coverage**: subnormals, powers of two, RFC 8785 Appendix B, boundary neighborhoods, pseudorandom

## Checksum Verification

```bash
sha256sum impl-schubfach/jcsfloat/testdata/golden_vectors.csv
# 593bdecbe0dccbc182bc3baf570b716887db25739fc61b7808764ecb966d5636

sha256sum impl-schubfach/jcsfloat/testdata/golden_stress_vectors.csv
# 287d21ac87e5665550f1baf86038302a0afc67a74a020dffb872f1a93b26d410
```

## Packaging

```bash
./scripts/package-artifact.sh
# Creates jcs-bench-lab-artifact-YYYYMMDD.tar.gz
```
