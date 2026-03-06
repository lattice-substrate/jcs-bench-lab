# Benchstat Snapshot

Generated at: 2026-03-05T20:58:34Z

Sources:
- API benchmark: `/home/lenny/jcs-bench-lab/results/latest-api-bench.txt`
- CLI summary: `/home/lenny/jcs-bench-lab/results/latest-cli-summary.csv`
- Quality report: `/home/lenny/jcs-bench-lab/results/latest-quality.json`

## Quality Gate

- determinism_failures: 0
- output_equality_failures: 0
- invalid_failure_parity_issues: 0
- oracle_mismatches: 0

## CLI Canonicalize (valid workloads)

### go:schubfach-vs-bd

_No paired samples found._

### rs:schubfach-vs-bd

_No paired samples found._

### schubfach:go-vs-rs

_No paired samples found._

### bd:go-vs-rs

_No paired samples found._

## CLI Verify (valid workloads)

### go:schubfach-vs-bd

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 4.865 | 6.641 | schubfach | 1.37x |
| array-256 | 1.838 | 2.139 | schubfach | 1.16x |
| canonical-minimal | 1.501 | 1.536 | schubfach | 1.02x |
| control-escapes | 1.515 | 1.854 | schubfach | 1.22x |
| deep-64 | 1.760 | 1.276 | json-canon | 1.38x |
| escaped-key-order | 1.601 | 1.424 | json-canon | 1.12x |
| long-string | 1.871 | 1.755 | json-canon | 1.07x |
| nested-mixed | 1.582 | 1.222 | json-canon | 1.29x |
| numeric-boundary | 1.678 | 1.331 | json-canon | 1.26x |
| rfc-key-sorting | 1.192 | 1.715 | schubfach | 1.44x |
| small | 1.226 | 1.881 | schubfach | 1.53x |
| surrogate-pair | 1.118 | 1.458 | schubfach | 1.30x |
| unicode | 1.437 | 1.491 | schubfach | 1.04x |
| verify-whitespace | 1.674 | 1.815 | schubfach | 1.08x |

### rs:schubfach-vs-bd

| workload | schubfach-rs (avg_ms) | json-canon-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.826 | 7.413 | schubfach-rs | 1.27x |
| array-256 | 2.620 | 2.522 | json-canon-rs | 1.04x |
| canonical-minimal | 1.983 | 2.428 | schubfach-rs | 1.22x |
| control-escapes | 1.775 | 2.049 | schubfach-rs | 1.15x |
| deep-64 | 2.014 | 2.273 | schubfach-rs | 1.13x |
| escaped-key-order | 1.670 | 2.164 | schubfach-rs | 1.30x |
| long-string | 2.391 | 2.744 | schubfach-rs | 1.15x |
| nested-mixed | 1.931 | 2.126 | schubfach-rs | 1.10x |
| numeric-boundary | 2.127 | 2.330 | schubfach-rs | 1.10x |
| rfc-key-sorting | 1.982 | 2.022 | schubfach-rs | 1.02x |
| small | 2.222 | 2.199 | json-canon-rs | 1.01x |
| surrogate-pair | 1.905 | 1.687 | json-canon-rs | 1.13x |
| unicode | 2.281 | 2.162 | json-canon-rs | 1.06x |
| verify-whitespace | 2.144 | 2.000 | json-canon-rs | 1.07x |

### schubfach:go-vs-rs

| workload | schubfach (avg_ms) | schubfach-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 4.865 | 5.826 | schubfach | 1.20x |
| array-256 | 1.838 | 2.620 | schubfach | 1.43x |
| canonical-minimal | 1.501 | 1.983 | schubfach | 1.32x |
| control-escapes | 1.515 | 1.775 | schubfach | 1.17x |
| deep-64 | 1.760 | 2.014 | schubfach | 1.14x |
| escaped-key-order | 1.601 | 1.670 | schubfach | 1.04x |
| long-string | 1.871 | 2.391 | schubfach | 1.28x |
| nested-mixed | 1.582 | 1.931 | schubfach | 1.22x |
| numeric-boundary | 1.678 | 2.127 | schubfach | 1.27x |
| rfc-key-sorting | 1.192 | 1.982 | schubfach | 1.66x |
| small | 1.226 | 2.222 | schubfach | 1.81x |
| surrogate-pair | 1.118 | 1.905 | schubfach | 1.70x |
| unicode | 1.437 | 2.281 | schubfach | 1.59x |
| verify-whitespace | 1.674 | 2.144 | schubfach | 1.28x |

### bd:go-vs-rs

| workload | json-canon (avg_ms) | json-canon-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 6.641 | 7.413 | json-canon | 1.12x |
| array-256 | 2.139 | 2.522 | json-canon | 1.18x |
| canonical-minimal | 1.536 | 2.428 | json-canon | 1.58x |
| control-escapes | 1.854 | 2.049 | json-canon | 1.11x |
| deep-64 | 1.276 | 2.273 | json-canon | 1.78x |
| escaped-key-order | 1.424 | 2.164 | json-canon | 1.52x |
| long-string | 1.755 | 2.744 | json-canon | 1.56x |
| nested-mixed | 1.222 | 2.126 | json-canon | 1.74x |
| numeric-boundary | 1.331 | 2.330 | json-canon | 1.75x |
| rfc-key-sorting | 1.715 | 2.022 | json-canon | 1.18x |
| small | 1.881 | 2.199 | json-canon | 1.17x |
| surrogate-pair | 1.458 | 1.687 | json-canon | 1.16x |
| unicode | 1.491 | 2.162 | json-canon | 1.45x |
| verify-whitespace | 1.815 | 2.000 | json-canon | 1.10x |

## API Benchmarks (ns/op)

### go:schubfach-vs-bd

_No paired samples found._

### rs:schubfach-vs-bd

| workload | schubfach-rs (ns/op) | json-canon-rs (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 5823459.500 | 6526755.000 | schubfach-rs | 1.12x |
| canonicalize/array-256 | 1598979.000 | 2224882.000 | schubfach-rs | 1.39x |
| canonicalize/canonical-minimal | 1302013.000 | 1187670.500 | json-canon-rs | 1.10x |
| canonicalize/control-escapes | 1593835.500 | 1346147.000 | json-canon-rs | 1.18x |
| canonicalize/deep-64 | 1507465.500 | 1672150.000 | schubfach-rs | 1.11x |
| canonicalize/escaped-key-order | 1211664.000 | 1378641.000 | schubfach-rs | 1.14x |
| canonicalize/long-string | 1592899.000 | 1617483.500 | schubfach-rs | 1.02x |
| canonicalize/nested-mixed | 1312620.000 | 1509635.000 | schubfach-rs | 1.15x |
| canonicalize/numeric-boundary | 1173266.000 | 1544288.000 | schubfach-rs | 1.32x |
| canonicalize/rfc-key-sorting | 1012122.000 | 1447956.500 | schubfach-rs | 1.43x |
| canonicalize/small | 1647242.000 | 1750408.000 | schubfach-rs | 1.06x |
| canonicalize/surrogate-pair | 1483949.500 | 1077030.500 | json-canon-rs | 1.38x |
| canonicalize/unicode | 1055125.000 | 1553776.000 | schubfach-rs | 1.47x |
| canonicalize/verify-whitespace | 1362122.500 | 1364327.500 | schubfach-rs | 1.00x |
| verify/canonical/array-2048 | 6698327.500 | 7792264.000 | schubfach-rs | 1.16x |
| verify/canonical/array-256 | 1950924.000 | 2158669.500 | schubfach-rs | 1.11x |
| verify/canonical/canonical-minimal | 1291678.000 | 1495041.000 | schubfach-rs | 1.16x |
| verify/canonical/control-escapes | 1427915.500 | 1689327.000 | schubfach-rs | 1.18x |
| verify/canonical/deep-64 | 1156885.500 | 1255886.000 | schubfach-rs | 1.09x |
| verify/canonical/escaped-key-order | 1320419.000 | 1721116.500 | schubfach-rs | 1.30x |
| verify/canonical/long-string | 1337485.000 | 1610070.000 | schubfach-rs | 1.20x |
| verify/canonical/nested-mixed | 1406312.000 | 1476600.000 | schubfach-rs | 1.05x |
| verify/canonical/numeric-boundary | 1085319.500 | 1524702.500 | schubfach-rs | 1.40x |
| verify/canonical/rfc-key-sorting | 1499163.500 | 1424403.500 | json-canon-rs | 1.05x |
| verify/canonical/small | 1605716.000 | 1343893.500 | json-canon-rs | 1.19x |
| verify/canonical/surrogate-pair | 1458631.500 | 1303438.500 | json-canon-rs | 1.12x |
| verify/canonical/unicode | 1453942.500 | 1826796.000 | schubfach-rs | 1.26x |
| verify/canonical/verify-whitespace | 1580496.000 | 1537442.500 | json-canon-rs | 1.03x |

### schubfach:go-vs-rs

_No paired samples found._

### bd:go-vs-rs

_No paired samples found._

## benchstat Output

benchstat unavailable; fallback summary above is non-inferential and not CI-gating.
