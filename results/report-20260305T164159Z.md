# Benchmark Report

Generated at: 2026-03-05T16:41:59Z

Sources:
- `/home/lenny/jcs-bench-lab/results/latest-api-bench.txt`
- `/home/lenny/jcs-bench-lab/results/latest-cli-summary.csv`
- `/home/lenny/jcs-bench-lab/results/latest-quality.json`
- `/home/lenny/jcs-bench-lab/results/latest-benchstat.md`
- `/home/lenny/jcs-bench-lab/results/latest-conformance.json`
- `/home/lenny/jcs-bench-lab/results/latest-stats.json`

- `/home/lenny/jcs-bench-lab/results/latest-fuzz.json`

## Executive Summary

- Conformance failures: 0
- Quality oracle mismatches: 0
- Differential fuzz failures: 0
- Statistically significant practical wins: 8
- Recommendation status: Recommend `schubfach` based on statistically significant practical wins with conformance/oracle gates passing.

## Conformance Evidence

- cyberphone: total=36 passed=36 failed=0
- lab-workload: total=116 passed=116 failed=0
- rfc8785: total=2 passed=2 failed=0

## Quality Findings

- determinism_failures: none
- output_equality_failures: none
- invalid_failure_parity_issues: none
- oracle_mismatches: none

## Performance Winners by Workload

### CLI canonicalize

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.466 | 7.342 | schubfach | 1.34x |
| array-256 | 2.025 | 2.340 | schubfach | 1.16x |
| canonical-minimal | 1.472 | 1.564 | schubfach | 1.06x |
| control-escapes | 1.479 | 1.550 | schubfach | 1.05x |
| deep-64 | 1.545 | 1.629 | schubfach | 1.05x |
| escaped-key-order | 1.503 | 1.652 | schubfach | 1.10x |
| long-string | 1.561 | 1.740 | schubfach | 1.11x |
| nested-mixed | 1.461 | 1.649 | schubfach | 1.13x |
| numeric-boundary | 1.470 | 1.602 | schubfach | 1.09x |
| rfc-key-sorting | 1.451 | 1.629 | schubfach | 1.12x |
| small | 1.414 | 1.494 | schubfach | 1.06x |
| surrogate-pair | 1.323 | 1.576 | schubfach | 1.19x |
| unicode | 1.429 | 1.616 | schubfach | 1.13x |
| verify-whitespace | 1.481 | 1.475 | json-canon | 1.00x |

### CLI verify

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.707 | 7.335 | schubfach | 1.29x |
| array-256 | 1.916 | 2.213 | schubfach | 1.16x |
| canonical-minimal | 1.511 | 1.618 | schubfach | 1.07x |
| control-escapes | 1.625 | 1.503 | json-canon | 1.08x |
| deep-64 | 1.480 | 1.712 | schubfach | 1.16x |
| escaped-key-order | 1.437 | 1.583 | schubfach | 1.10x |
| long-string | 1.550 | 1.828 | schubfach | 1.18x |
| nested-mixed | 1.481 | 1.529 | schubfach | 1.03x |
| numeric-boundary | 1.507 | 1.589 | schubfach | 1.05x |
| rfc-key-sorting | 1.452 | 1.478 | schubfach | 1.02x |
| small | 1.459 | 1.463 | schubfach | 1.00x |
| surrogate-pair | 1.389 | 1.528 | schubfach | 1.10x |
| unicode | 1.309 | 1.592 | schubfach | 1.22x |
| verify-whitespace | 1.387 | 1.595 | schubfach | 1.15x |

### API

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2561729.700 | 4100567.700 | schubfach | 1.60x |
| canonicalize/array-256 | 306198.500 | 475343.900 | schubfach | 1.55x |
| canonicalize/canonical-minimal | 351.560 | 531.430 | schubfach | 1.51x |
| canonicalize/control-escapes | 1349.800 | 1264.900 | json-canon | 1.07x |
| canonicalize/deep | 36879.800 | 36765.500 | json-canon | 1.00x |
| canonicalize/deep-64 | 16244.400 | 16336.800 | schubfach | 1.01x |
| canonicalize/escaped-key-order | 705.980 | 1129.900 | schubfach | 1.60x |
| canonicalize/large | 11976455.000 | 20612503.900 | schubfach | 1.72x |
| canonicalize/long-string | 76796.100 | 75367.800 | json-canon | 1.02x |
| canonicalize/medium | 346537.400 | 586299.200 | schubfach | 1.69x |
| canonicalize/mixed-prod | 136379.000 | 180597.100 | schubfach | 1.32x |
| canonicalize/nested-mixed | 2176.000 | 3144.200 | schubfach | 1.44x |
| canonicalize/number-heavy | 14115.500 | 21353.100 | schubfach | 1.51x |
| canonicalize/numeric-boundary | 12712.400 | 18781.000 | schubfach | 1.48x |
| canonicalize/rfc-key-sorting | 3002.900 | 2977.400 | json-canon | 1.01x |
| canonicalize/small | 1242.900 | 1673.900 | schubfach | 1.35x |
| canonicalize/surrogate-pair | 649.910 | 641.800 | json-canon | 1.01x |
| canonicalize/unicode | 1559.500 | 1551.100 | json-canon | 1.01x |
| canonicalize/verify-whitespace | 915.130 | 1561.800 | schubfach | 1.71x |
| verify/canonical/array-2048 | 2483304.200 | 4031763.700 | schubfach | 1.62x |
| verify/canonical/array-256 | 317027.400 | 485794.700 | schubfach | 1.53x |
| verify/canonical/canonical-minimal | 360.160 | 547.070 | schubfach | 1.52x |
| verify/canonical/control-escapes | 1330.600 | 1270.100 | json-canon | 1.05x |
| verify/canonical/deep | 38141.200 | 38114.800 | json-canon | 1.00x |
| verify/canonical/deep-64 | 16803.900 | 16810.200 | schubfach | 1.00x |
| verify/canonical/escaped-key-order | 692.740 | 1124.600 | schubfach | 1.62x |
| verify/canonical/large | 11583371.600 | 20145353.200 | schubfach | 1.74x |
| verify/canonical/long-string | 77610.300 | 77086.700 | json-canon | 1.01x |
| verify/canonical/medium | 344421.200 | 582648.300 | schubfach | 1.69x |
| verify/canonical/mixed-prod | 134723.400 | 177058.100 | schubfach | 1.31x |
| verify/canonical/nested-mixed | 2221.200 | 3114.800 | schubfach | 1.40x |
| verify/canonical/number-heavy | 13950.600 | 21546.000 | schubfach | 1.54x |
| verify/canonical/numeric-boundary | 12875.600 | 18819.500 | schubfach | 1.46x |
| verify/canonical/rfc-key-sorting | 2859.200 | 2856.100 | json-canon | 1.00x |
| verify/canonical/small | 1232.700 | 1655.500 | schubfach | 1.34x |
| verify/canonical/surrogate-pair | 658.750 | 656.930 | json-canon | 1.00x |
| verify/canonical/unicode | 1555.200 | 1555.200 | schubfach | 1.00x |
| verify/canonical/verify-whitespace | 939.430 | 1592.900 | schubfach | 1.70x |
| verify/noncanonical/control-escapes | 1322.300 | 1273.400 | json-canon | 1.04x |
| verify/noncanonical/escaped-key-order | 704.360 | 1136.600 | schubfach | 1.61x |
| verify/noncanonical/large | 12019413.200 | 20762854.000 | schubfach | 1.73x |
| verify/noncanonical/medium | 340768.000 | 576457.000 | schubfach | 1.69x |
| verify/noncanonical/mixed-prod | 135143.100 | 179242.600 | schubfach | 1.33x |
| verify/noncanonical/nested-mixed | 2183.900 | 3076.600 | schubfach | 1.41x |
| verify/noncanonical/number-heavy | 13947.400 | 21264.300 | schubfach | 1.52x |
| verify/noncanonical/numeric-boundary | 12792.700 | 18670.100 | schubfach | 1.46x |
| verify/noncanonical/rfc-key-sorting | 2998.900 | 2979.500 | json-canon | 1.01x |
| verify/noncanonical/small | 1241.100 | 1674.100 | schubfach | 1.35x |
| verify/noncanonical/surrogate-pair | 647.670 | 647.350 | json-canon | 1.00x |
| verify/noncanonical/unicode | 1546.300 | 1557.500 | schubfach | 1.01x |
| verify/noncanonical/verify-whitespace | 908.000 | 1566.600 | schubfach | 1.73x |

## Statistical Inference

| track | mode | workload | winner | speedup | ci95 | p-value | practical |
|---|---|---|---|---:|---|---:|---|
| e2e | canonicalize | array-2048 | schubfach | 1.343x | [1.259, 1.432] | 0.0002 | true |
| e2e | canonicalize | array-256 | schubfach | 1.156x | [1.030, 1.305] | 0.0296 | true |
| e2e | canonicalize | canonical-minimal | schubfach | 1.063x | [0.939, 1.207] | 0.3595 | true |
| e2e | canonicalize | control-escapes | schubfach | 1.048x | [0.940, 1.168] | 0.4201 | true |
| e2e | canonicalize | deep-64 | schubfach | 1.055x | [0.913, 1.221] | 0.4797 | true |
| e2e | canonicalize | escaped-key-order | schubfach | 1.099x | [0.995, 1.218] | 0.0886 | true |
| e2e | canonicalize | long-string | schubfach | 1.114x | [0.986, 1.263] | 0.1062 | true |
| e2e | canonicalize | nested-mixed | schubfach | 1.129x | [0.999, 1.275] | 0.0692 | true |
| e2e | canonicalize | numeric-boundary | schubfach | 1.090x | [0.959, 1.236] | 0.2030 | true |
| e2e | canonicalize | rfc-key-sorting | schubfach | 1.123x | [1.005, 1.265] | 0.0620 | true |
| e2e | canonicalize | small | schubfach | 1.056x | [0.940, 1.196] | 0.3907 | true |
| e2e | canonicalize | surrogate-pair | schubfach | 1.191x | [1.043, 1.355] | 0.0168 | true |
| e2e | canonicalize | unicode | schubfach | 1.131x | [0.993, 1.289] | 0.0810 | true |
| e2e | canonicalize | verify-whitespace | json-canon | 1.004x | [0.878, 1.133] | 0.9488 | false |
| e2e | verify | array-2048 | schubfach | 1.285x | [1.195, 1.378] | 0.0002 | true |
| e2e | verify | array-256 | schubfach | 1.155x | [1.027, 1.308] | 0.0298 | true |
| e2e | verify | canonical-minimal | schubfach | 1.071x | [0.948, 1.202] | 0.2849 | true |
| e2e | verify | control-escapes | json-canon | 1.081x | [0.953, 1.238] | 0.2480 | true |
| e2e | verify | deep-64 | schubfach | 1.156x | [1.022, 1.314] | 0.0310 | true |
| e2e | verify | escaped-key-order | schubfach | 1.101x | [0.988, 1.235] | 0.1020 | true |
| e2e | verify | long-string | schubfach | 1.179x | [1.022, 1.353] | 0.0344 | true |
| e2e | verify | nested-mixed | schubfach | 1.033x | [0.896, 1.188] | 0.6779 | true |
| e2e | verify | numeric-boundary | schubfach | 1.055x | [0.918, 1.214] | 0.4797 | true |
| e2e | verify | rfc-key-sorting | schubfach | 1.018x | [0.894, 1.159] | 0.7858 | false |
| e2e | verify | small | schubfach | 1.003x | [0.878, 1.156] | 0.9676 | false |
| e2e | verify | surrogate-pair | schubfach | 1.100x | [0.955, 1.263] | 0.1938 | true |
| e2e | verify | unicode | schubfach | 1.216x | [1.065, 1.390] | 0.0120 | true |
| e2e | verify | verify-whitespace | schubfach | 1.150x | [1.004, 1.327] | 0.0618 | true |

## benchstat Snippet

```text
# Benchstat Snapshot

Generated at: 2026-03-05T16:39:29Z

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

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.466 | 7.342 | schubfach | 1.34x |
| array-256 | 2.025 | 2.340 | schubfach | 1.16x |
| canonical-minimal | 1.472 | 1.564 | schubfach | 1.06x |
| control-escapes | 1.479 | 1.550 | schubfach | 1.05x |
| deep-64 | 1.545 | 1.629 | schubfach | 1.05x |
| escaped-key-order | 1.503 | 1.652 | schubfach | 1.10x |
| long-string | 1.561 | 1.740 | schubfach | 1.11x |
| nested-mixed | 1.461 | 1.649 | schubfach | 1.13x |
| numeric-boundary | 1.470 | 1.602 | schubfach | 1.09x |
| rfc-key-sorting | 1.451 | 1.629 | schubfach | 1.12x |
| small | 1.414 | 1.494 | schubfach | 1.06x |
| surrogate-pair | 1.323 | 1.576 | schubfach | 1.19x |
| unicode | 1.429 | 1.616 | schubfach | 1.13x |
| verify-whitespace | 1.481 | 1.475 | json-canon | 1.00x |

## CLI Verify (valid workloads)

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.707 | 7.335 | schubfach | 1.29x |
| array-256 | 1.916 | 2.213 | schubfach | 1.16x |
| canonical-minimal | 1.511 | 1.618 | schubfach | 1.07x |
| control-escapes | 1.625 | 1.503 | json-canon | 1.08x |
| deep-64 | 1.480 | 1.712 | schubfach | 1.16x |
| escaped-key-order | 1.437 | 1.583 | schubfach | 1.10x |
| long-string | 1.550 | 1.828 | schubfach | 1.18x |
| nested-mixed | 1.481 | 1.529 | schubfach | 1.03x |
| numeric-boundary | 1.507 | 1.589 | schubfach | 1.05x |
| rfc-key-sorting | 1.452 | 1.478 | schubfach | 1.02x |
| small | 1.459 | 1.463 | schubfach | 1.00x |
| surrogate-pair | 1.389 | 1.528 | schubfach | 1.10x |
| unicode | 1.309 | 1.592 | schubfach | 1.22x |
| verify-whitespace | 1.387 | 1.595 | schubfach | 1.15x |

## API Benchmarks (ns/op)

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2561729.700 | 4100567.700 | schubfach | 1.60x |
| canonicalize/array-256 | 306198.500 | 475343.900 | schubfach | 1.55x |
| canonicalize/canonical-minimal | 351.560 | 531.430 | schubfach | 1.51x |
| canonicalize/control-escapes | 1349.800 | 1264.900 | json-canon | 1.07x |
| canonicalize/deep | 36879.800 | 36765.500 | json-canon | 1.00x |
| canonicalize/deep-64 | 16244.400 | 16336.800 | schubfach | 1.01x |
| canonicalize/escaped-key-order | 705.980 | 1129.900 | schubfach | 1.60x |
| canonicalize/large | 11976455.000 | 20612503.900 | schubfach | 1.72x |
| canonicalize/long-string | 76796.100 | 75367.800 | json-canon | 1.02x |
| canonicalize/medium | 346537.400 | 586299.200 | schubfach | 1.69x |
| canonicalize/mixed-prod | 136379.000 | 180597.100 | schubfach | 1.32x |
| canonicalize/nested-mixed | 2176.000 | 3144.200 | schubfach | 1.44x |
| canonicalize/number-heavy | 14115.500 | 21353.100 | schubfach | 1.51x |
| canonicalize/numeric-boundary | 12712.400 | 18781.000 | schubfach | 1.48x |
| canonicalize/rfc-key-sorting | 3002.900 | 2977.400 | json-canon | 1.01x |
| canonicalize/small | 1242.900 | 1673.900 | schubfach | 1.35x |
| canonicalize/surrogate-pair | 649.910 | 641.800 | json-canon | 1.01x |
| canonicalize/unicode | 1559.500 | 1551.100 | json-canon | 1.01x |
| canonicalize/verify-whitespace | 915.130 | 1561.800 | schubfach | 1.71x |
| verify/canonical/array-2048 | 2483304.200 | 4031763.700 | schubfach | 1.62x |
| verify/canonical/array-256 | 317027.400 | 485794.700 | schubfach | 1.53x |
| verify/canonical/canonical-minimal | 360.160 | 547.070 | schubfach | 1.52x |
```

## Production Recommendation

Recommend `schubfach` based on statistically significant practical wins with conformance/oracle gates passing.
