# Benchmark Report

Generated at: 2026-03-05T06:13:57Z

Sources:
- `/home/lenny/jcs-bench-lab/results/latest-api-bench.txt`
- `/home/lenny/jcs-bench-lab/results/latest-cli-summary.csv`
- `/home/lenny/jcs-bench-lab/results/latest-quality.json`
- `/home/lenny/jcs-bench-lab/results/latest-benchstat.md`
- `/home/lenny/jcs-bench-lab/results/latest-conformance.json`
- `/home/lenny/jcs-bench-lab/results/latest-stats.json`

- `results/latest-fuzz.json`

## Executive Summary

- Conformance failures: 0
- Quality oracle mismatches: 0
- Differential fuzz failures: 0
- Statistically significant practical wins: 2
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
| array-2048 | 19.556 | 17.999 | json-canon | 1.09x |
| array-256 | 6.775 | 7.930 | schubfach | 1.17x |
| canonical-minimal | 4.881 | 6.475 | schubfach | 1.33x |
| control-escapes | 3.742 | 5.387 | schubfach | 1.44x |
| deep-64 | 5.104 | 6.597 | schubfach | 1.29x |
| escaped-key-order | 5.305 | 5.377 | schubfach | 1.01x |
| long-string | 5.491 | 5.917 | schubfach | 1.08x |
| nested-mixed | 6.103 | 4.541 | json-canon | 1.34x |
| numeric-boundary | 4.810 | 4.465 | json-canon | 1.08x |
| rfc-key-sorting | 6.562 | 4.731 | json-canon | 1.39x |
| small | 5.795 | 4.136 | json-canon | 1.40x |
| surrogate-pair | 5.383 | 6.586 | schubfach | 1.22x |
| unicode | 4.707 | 5.706 | schubfach | 1.21x |
| verify-whitespace | 5.221 | 5.230 | schubfach | 1.00x |

### CLI verify

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 23.157 | 28.189 | schubfach | 1.22x |
| array-256 | 7.129 | 8.914 | schubfach | 1.25x |
| canonical-minimal | 4.347 | 5.031 | schubfach | 1.16x |
| control-escapes | 5.195 | 6.435 | schubfach | 1.24x |
| deep-64 | 4.995 | 6.243 | schubfach | 1.25x |
| escaped-key-order | 4.434 | 4.839 | schubfach | 1.09x |
| long-string | 5.339 | 5.513 | schubfach | 1.03x |
| nested-mixed | 5.932 | 5.433 | json-canon | 1.09x |
| numeric-boundary | 4.172 | 5.884 | schubfach | 1.41x |
| rfc-key-sorting | 6.011 | 5.701 | json-canon | 1.05x |
| small | 5.075 | 5.089 | schubfach | 1.00x |
| surrogate-pair | 4.940 | 5.364 | schubfach | 1.09x |
| unicode | 5.757 | 5.159 | json-canon | 1.12x |
| verify-whitespace | 4.462 | 5.084 | schubfach | 1.14x |

### API

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 10390436.000 | 13400888.000 | schubfach | 1.29x |
| canonicalize/array-256 | 1242055.000 | 1483814.000 | schubfach | 1.19x |
| canonicalize/canonical-minimal | 1034.000 | 1147.000 | schubfach | 1.11x |
| canonicalize/control-escapes | 4396.000 | 4419.000 | schubfach | 1.01x |
| canonicalize/deep | 134957.000 | 141426.000 | schubfach | 1.05x |
| canonicalize/deep-64 | 55526.000 | 64079.000 | schubfach | 1.15x |
| canonicalize/escaped-key-order | 2706.000 | 3069.000 | schubfach | 1.13x |
| canonicalize/large | 34419219.000 | 32711983.000 | json-canon | 1.05x |
| canonicalize/long-string | 109449.000 | 104875.000 | json-canon | 1.04x |
| canonicalize/medium | 1235456.000 | 1610904.000 | schubfach | 1.30x |
| canonicalize/mixed-prod | 478684.000 | 554694.000 | schubfach | 1.16x |
| canonicalize/nested-mixed | 6505.000 | 9405.000 | schubfach | 1.45x |
| canonicalize/number-heavy | 17501.000 | 24905.000 | schubfach | 1.42x |
| canonicalize/numeric-boundary | 14223.000 | 19645.000 | schubfach | 1.38x |
| canonicalize/rfc-key-sorting | 9266.000 | 8312.000 | json-canon | 1.11x |
| canonicalize/small | 4030.000 | 5605.000 | schubfach | 1.39x |
| canonicalize/surrogate-pair | 1976.000 | 2052.000 | schubfach | 1.04x |
| canonicalize/unicode | 4731.000 | 5217.000 | schubfach | 1.10x |
| canonicalize/verify-whitespace | 3699.000 | 4280.000 | schubfach | 1.16x |
| verify/canonical/array-2048 | 11198095.000 | 13106305.000 | schubfach | 1.17x |
| verify/canonical/array-256 | 1227366.000 | 1704143.000 | schubfach | 1.39x |
| verify/canonical/canonical-minimal | 1379.000 | 1624.000 | schubfach | 1.18x |
| verify/canonical/control-escapes | 3780.000 | 4964.000 | schubfach | 1.31x |
| verify/canonical/deep | 112643.000 | 142435.000 | schubfach | 1.26x |
| verify/canonical/deep-64 | 63230.000 | 65230.000 | schubfach | 1.03x |
| verify/canonical/escaped-key-order | 2413.000 | 3476.000 | schubfach | 1.44x |
| verify/canonical/large | 33524820.000 | 32702136.000 | json-canon | 1.03x |
| verify/canonical/long-string | 143180.000 | 127728.000 | json-canon | 1.12x |
| verify/canonical/medium | 1499521.000 | 1922577.000 | schubfach | 1.28x |
| verify/canonical/mixed-prod | 545587.000 | 666333.000 | schubfach | 1.22x |
| verify/canonical/nested-mixed | 8609.000 | 9745.000 | schubfach | 1.13x |
| verify/canonical/number-heavy | 20347.000 | 25870.000 | schubfach | 1.27x |
| verify/canonical/numeric-boundary | 15163.000 | 21140.000 | schubfach | 1.39x |
| verify/canonical/rfc-key-sorting | 10990.000 | 9040.000 | json-canon | 1.22x |
| verify/canonical/small | 4072.000 | 6616.000 | schubfach | 1.62x |
| verify/canonical/surrogate-pair | 2368.000 | 2875.000 | schubfach | 1.21x |
| verify/canonical/unicode | 4569.000 | 5358.000 | schubfach | 1.17x |
| verify/canonical/verify-whitespace | 2761.000 | 4397.000 | schubfach | 1.59x |
| verify/noncanonical/control-escapes | 4941.000 | 5013.000 | schubfach | 1.01x |
| verify/noncanonical/escaped-key-order | 2084.000 | 2792.000 | schubfach | 1.34x |
| verify/noncanonical/large | 28956019.000 | 32457332.000 | schubfach | 1.12x |
| verify/noncanonical/medium | 754185.000 | 1650602.000 | schubfach | 2.19x |
| verify/noncanonical/mixed-prod | 474069.000 | 671622.000 | schubfach | 1.42x |
| verify/noncanonical/nested-mixed | 6415.000 | 9879.000 | schubfach | 1.54x |
| verify/noncanonical/number-heavy | 16509.000 | 23918.000 | schubfach | 1.45x |
| verify/noncanonical/numeric-boundary | 14554.000 | 19888.000 | schubfach | 1.37x |
| verify/noncanonical/rfc-key-sorting | 8347.000 | 7785.000 | json-canon | 1.07x |
| verify/noncanonical/small | 4756.000 | 4251.000 | json-canon | 1.12x |
| verify/noncanonical/surrogate-pair | 2116.000 | 1955.000 | json-canon | 1.08x |
| verify/noncanonical/unicode | 4861.000 | 5250.000 | schubfach | 1.08x |
| verify/noncanonical/verify-whitespace | 2710.000 | 3415.000 | schubfach | 1.26x |

## Statistical Inference

| track | mode | workload | winner | speedup | ci95 | p-value | practical |
|---|---|---|---|---:|---|---:|---|
| e2e | canonicalize | array-2048 | json-canon | 1.086x | [0.676, 1.927] | 0.7591 | true |
| e2e | canonicalize | array-256 | schubfach | 1.171x | [0.750, 1.858] | 0.5165 | true |
| e2e | canonicalize | canonical-minimal | schubfach | 1.326x | [1.124, 1.828] | 0.0003 | true |
| e2e | canonicalize | control-escapes | schubfach | 1.439x | [0.865, 2.809] | 0.2403 | true |
| e2e | canonicalize | deep-64 | schubfach | 1.293x | [1.045, 1.757] | 0.0497 | true |
| e2e | canonicalize | escaped-key-order | schubfach | 1.014x | [0.686, 1.472] | 0.9600 | false |
| e2e | canonicalize | long-string | schubfach | 1.077x | [0.735, 1.594] | 0.6811 | true |
| e2e | canonicalize | nested-mixed | json-canon | 1.344x | [0.853, 2.280] | 0.2329 | true |
| e2e | canonicalize | numeric-boundary | json-canon | 1.077x | [0.657, 1.817] | 0.8870 | true |
| e2e | canonicalize | rfc-key-sorting | json-canon | 1.387x | [1.053, 2.075] | 0.0610 | true |
| e2e | canonicalize | small | json-canon | 1.401x | [1.008, 2.201] | 0.0760 | true |
| e2e | canonicalize | surrogate-pair | schubfach | 1.223x | [0.850, 1.831] | 0.3242 | true |
| e2e | canonicalize | unicode | schubfach | 1.212x | [0.785, 2.059] | 0.4209 | true |
| e2e | canonicalize | verify-whitespace | schubfach | 1.002x | [0.615, 1.519] | 0.9960 | false |
| e2e | verify | array-2048 | schubfach | 1.217x | [0.858, 1.641] | 0.2692 | true |
| e2e | verify | array-256 | schubfach | 1.250x | [0.869, 1.829] | 0.2912 | true |
| e2e | verify | canonical-minimal | schubfach | 1.157x | [0.711, 1.785] | 0.5402 | true |
| e2e | verify | control-escapes | schubfach | 1.238x | [0.996, 1.568] | 0.0906 | true |
| e2e | verify | deep-64 | schubfach | 1.250x | [0.919, 1.908] | 0.2289 | true |
| e2e | verify | escaped-key-order | schubfach | 1.091x | [0.691, 1.645] | 0.7178 | true |
| e2e | verify | long-string | schubfach | 1.033x | [0.648, 1.601] | 0.8944 | true |
| e2e | verify | nested-mixed | json-canon | 1.092x | [0.854, 1.548] | 0.6101 | true |
| e2e | verify | numeric-boundary | schubfach | 1.411x | [0.903, 2.543] | 0.1866 | true |
| e2e | verify | rfc-key-sorting | json-canon | 1.054x | [0.892, 1.338] | 0.7304 | true |
| e2e | verify | small | schubfach | 1.003x | [0.657, 1.440] | 0.9937 | false |
| e2e | verify | surrogate-pair | schubfach | 1.086x | [0.787, 1.482] | 0.6315 | true |
| e2e | verify | unicode | json-canon | 1.116x | [0.880, 1.556] | 0.4585 | true |
| e2e | verify | verify-whitespace | schubfach | 1.139x | [0.806, 1.725] | 0.5202 | true |

## benchstat Snippet

```text
# Benchstat Snapshot

Generated at: 2026-03-05T06:13:32Z

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
| array-2048 | 19.556 | 17.999 | json-canon | 1.09x |
| array-256 | 6.775 | 7.930 | schubfach | 1.17x |
| canonical-minimal | 4.881 | 6.475 | schubfach | 1.33x |
| control-escapes | 3.742 | 5.387 | schubfach | 1.44x |
| deep-64 | 5.104 | 6.597 | schubfach | 1.29x |
| escaped-key-order | 5.305 | 5.377 | schubfach | 1.01x |
| long-string | 5.491 | 5.917 | schubfach | 1.08x |
| nested-mixed | 6.103 | 4.541 | json-canon | 1.34x |
| numeric-boundary | 4.810 | 4.465 | json-canon | 1.08x |
| rfc-key-sorting | 6.562 | 4.731 | json-canon | 1.39x |
| small | 5.795 | 4.136 | json-canon | 1.40x |
| surrogate-pair | 5.383 | 6.586 | schubfach | 1.22x |
| unicode | 4.707 | 5.706 | schubfach | 1.21x |
| verify-whitespace | 5.221 | 5.230 | schubfach | 1.00x |

## CLI Verify (valid workloads)

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 23.157 | 28.189 | schubfach | 1.22x |
| array-256 | 7.129 | 8.914 | schubfach | 1.25x |
| canonical-minimal | 4.347 | 5.031 | schubfach | 1.16x |
| control-escapes | 5.195 | 6.435 | schubfach | 1.24x |
| deep-64 | 4.995 | 6.243 | schubfach | 1.25x |
| escaped-key-order | 4.434 | 4.839 | schubfach | 1.09x |
| long-string | 5.339 | 5.513 | schubfach | 1.03x |
| nested-mixed | 5.932 | 5.433 | json-canon | 1.09x |
| numeric-boundary | 4.172 | 5.884 | schubfach | 1.41x |
| rfc-key-sorting | 6.011 | 5.701 | json-canon | 1.05x |
| small | 5.075 | 5.089 | schubfach | 1.00x |
| surrogate-pair | 4.940 | 5.364 | schubfach | 1.09x |
| unicode | 5.757 | 5.159 | json-canon | 1.12x |
| verify-whitespace | 4.462 | 5.084 | schubfach | 1.14x |

## API Benchmarks (ns/op)

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 10390436.000 | 13400888.000 | schubfach | 1.29x |
| canonicalize/array-256 | 1242055.000 | 1483814.000 | schubfach | 1.19x |
| canonicalize/canonical-minimal | 1034.000 | 1147.000 | schubfach | 1.11x |
| canonicalize/control-escapes | 4396.000 | 4419.000 | schubfach | 1.01x |
| canonicalize/deep | 134957.000 | 141426.000 | schubfach | 1.05x |
| canonicalize/deep-64 | 55526.000 | 64079.000 | schubfach | 1.15x |
| canonicalize/escaped-key-order | 2706.000 | 3069.000 | schubfach | 1.13x |
| canonicalize/large | 34419219.000 | 32711983.000 | json-canon | 1.05x |
| canonicalize/long-string | 109449.000 | 104875.000 | json-canon | 1.04x |
| canonicalize/medium | 1235456.000 | 1610904.000 | schubfach | 1.30x |
| canonicalize/mixed-prod | 478684.000 | 554694.000 | schubfach | 1.16x |
| canonicalize/nested-mixed | 6505.000 | 9405.000 | schubfach | 1.45x |
| canonicalize/number-heavy | 17501.000 | 24905.000 | schubfach | 1.42x |
| canonicalize/numeric-boundary | 14223.000 | 19645.000 | schubfach | 1.38x |
| canonicalize/rfc-key-sorting | 9266.000 | 8312.000 | json-canon | 1.11x |
| canonicalize/small | 4030.000 | 5605.000 | schubfach | 1.39x |
| canonicalize/surrogate-pair | 1976.000 | 2052.000 | schubfach | 1.04x |
| canonicalize/unicode | 4731.000 | 5217.000 | schubfach | 1.10x |
| canonicalize/verify-whitespace | 3699.000 | 4280.000 | schubfach | 1.16x |
| verify/canonical/array-2048 | 11198095.000 | 13106305.000 | schubfach | 1.17x |
| verify/canonical/array-256 | 1227366.000 | 1704143.000 | schubfach | 1.39x |
| verify/canonical/canonical-minimal | 1379.000 | 1624.000 | schubfach | 1.18x |
```

## Production Recommendation

Recommend `schubfach` based on statistically significant practical wins with conformance/oracle gates passing.
