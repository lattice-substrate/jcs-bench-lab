# Benchmark Report

Generated at: 2026-03-05T15:42:51Z

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
- Statistically significant practical wins: 5
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
| array-2048 | 5.221 | 6.871 | schubfach | 1.32x |
| array-256 | 1.920 | 2.016 | schubfach | 1.05x |
| canonical-minimal | 1.391 | 1.567 | schubfach | 1.13x |
| control-escapes | 1.488 | 1.576 | schubfach | 1.06x |
| deep-64 | 1.419 | 1.551 | schubfach | 1.09x |
| escaped-key-order | 1.489 | 1.494 | schubfach | 1.00x |
| long-string | 1.669 | 1.814 | schubfach | 1.09x |
| nested-mixed | 1.230 | 1.667 | schubfach | 1.35x |
| numeric-boundary | 1.417 | 1.730 | schubfach | 1.22x |
| rfc-key-sorting | 1.417 | 1.573 | schubfach | 1.11x |
| small | 1.415 | 1.480 | schubfach | 1.05x |
| surrogate-pair | 1.408 | 1.394 | json-canon | 1.01x |
| unicode | 1.376 | 1.547 | schubfach | 1.12x |
| verify-whitespace | 1.337 | 1.476 | schubfach | 1.10x |

### CLI verify

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.373 | 6.968 | schubfach | 1.30x |
| array-256 | 1.834 | 2.241 | schubfach | 1.22x |
| canonical-minimal | 1.483 | 1.628 | schubfach | 1.10x |
| control-escapes | 1.546 | 1.557 | schubfach | 1.01x |
| deep-64 | 1.510 | 1.391 | json-canon | 1.09x |
| escaped-key-order | 1.384 | 1.575 | schubfach | 1.14x |
| long-string | 1.720 | 1.751 | schubfach | 1.02x |
| nested-mixed | 1.296 | 1.466 | schubfach | 1.13x |
| numeric-boundary | 1.398 | 1.548 | schubfach | 1.11x |
| rfc-key-sorting | 1.373 | 1.543 | schubfach | 1.12x |
| small | 1.453 | 1.550 | schubfach | 1.07x |
| surrogate-pair | 1.412 | 1.517 | schubfach | 1.07x |
| unicode | 1.503 | 1.567 | schubfach | 1.04x |
| verify-whitespace | 1.380 | 1.512 | schubfach | 1.10x |

### API

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2548995.833 | 4047979.833 | schubfach | 1.59x |
| canonicalize/array-256 | 303347.500 | 474368.500 | schubfach | 1.56x |
| canonicalize/canonical-minimal | 352.283 | 529.300 | schubfach | 1.50x |
| canonicalize/control-escapes | 1335.167 | 1251.333 | json-canon | 1.07x |
| canonicalize/deep | 36651.333 | 36688.667 | schubfach | 1.00x |
| canonicalize/deep-64 | 16236.667 | 16239.833 | schubfach | 1.00x |
| canonicalize/escaped-key-order | 707.433 | 1124.667 | schubfach | 1.59x |
| canonicalize/large | 11928619.000 | 20545683.667 | schubfach | 1.72x |
| canonicalize/long-string | 74802.167 | 73587.500 | json-canon | 1.02x |
| canonicalize/medium | 340270.333 | 581178.167 | schubfach | 1.71x |
| canonicalize/mixed-prod | 136311.000 | 178287.333 | schubfach | 1.31x |
| canonicalize/nested-mixed | 2181.833 | 3065.667 | schubfach | 1.41x |
| canonicalize/number-heavy | 13529.500 | 20854.833 | schubfach | 1.54x |
| canonicalize/numeric-boundary | 12395.833 | 18294.167 | schubfach | 1.48x |
| canonicalize/rfc-key-sorting | 3018.333 | 2967.833 | json-canon | 1.02x |
| canonicalize/small | 1243.167 | 1667.000 | schubfach | 1.34x |
| canonicalize/surrogate-pair | 647.250 | 639.950 | json-canon | 1.01x |
| canonicalize/unicode | 1548.667 | 1558.667 | schubfach | 1.01x |
| canonicalize/verify-whitespace | 906.500 | 1563.000 | schubfach | 1.72x |
| verify/canonical/array-2048 | 2469781.333 | 4029019.000 | schubfach | 1.63x |
| verify/canonical/array-256 | 317223.833 | 489779.667 | schubfach | 1.54x |
| verify/canonical/canonical-minimal | 358.083 | 548.567 | schubfach | 1.53x |
| verify/canonical/control-escapes | 1318.833 | 1273.167 | json-canon | 1.04x |
| verify/canonical/deep | 37935.667 | 38057.500 | schubfach | 1.00x |
| verify/canonical/deep-64 | 16756.500 | 16837.667 | schubfach | 1.00x |
| verify/canonical/escaped-key-order | 684.650 | 1130.333 | schubfach | 1.65x |
| verify/canonical/large | 11506130.667 | 20018040.000 | schubfach | 1.74x |
| verify/canonical/long-string | 76126.667 | 75923.000 | json-canon | 1.00x |
| verify/canonical/medium | 341654.667 | 580655.333 | schubfach | 1.70x |
| verify/canonical/mixed-prod | 134312.000 | 176078.833 | schubfach | 1.31x |
| verify/canonical/nested-mixed | 2218.167 | 3117.833 | schubfach | 1.41x |
| verify/canonical/number-heavy | 13608.000 | 21223.833 | schubfach | 1.56x |
| verify/canonical/numeric-boundary | 12525.833 | 18510.833 | schubfach | 1.48x |
| verify/canonical/rfc-key-sorting | 2853.667 | 2834.833 | json-canon | 1.01x |
| verify/canonical/small | 1223.000 | 1658.667 | schubfach | 1.36x |
| verify/canonical/surrogate-pair | 659.250 | 654.400 | json-canon | 1.01x |
| verify/canonical/unicode | 1540.000 | 1546.333 | schubfach | 1.00x |
| verify/canonical/verify-whitespace | 936.467 | 1582.833 | schubfach | 1.69x |
| verify/noncanonical/control-escapes | 1308.500 | 1261.333 | json-canon | 1.04x |
| verify/noncanonical/escaped-key-order | 703.350 | 1124.500 | schubfach | 1.60x |
| verify/noncanonical/large | 11908081.333 | 20607667.167 | schubfach | 1.73x |
| verify/noncanonical/medium | 343706.167 | 574766.500 | schubfach | 1.67x |
| verify/noncanonical/mixed-prod | 136452.667 | 178493.167 | schubfach | 1.31x |
| verify/noncanonical/nested-mixed | 2185.000 | 3063.000 | schubfach | 1.40x |
| verify/noncanonical/number-heavy | 13609.333 | 20847.833 | schubfach | 1.53x |
| verify/noncanonical/numeric-boundary | 12518.833 | 18296.500 | schubfach | 1.46x |
| verify/noncanonical/rfc-key-sorting | 3017.500 | 2968.167 | json-canon | 1.02x |
| verify/noncanonical/small | 1245.833 | 1664.167 | schubfach | 1.34x |
| verify/noncanonical/surrogate-pair | 650.567 | 644.033 | json-canon | 1.01x |
| verify/noncanonical/unicode | 1556.833 | 1552.833 | json-canon | 1.00x |
| verify/noncanonical/verify-whitespace | 913.567 | 1563.667 | schubfach | 1.71x |

## Statistical Inference

| track | mode | workload | winner | speedup | ci95 | p-value | practical |
|---|---|---|---|---:|---|---:|---|
| e2e | canonicalize | array-2048 | schubfach | 1.316x | [1.222, 1.419] | 0.0007 | true |
| e2e | canonicalize | array-256 | schubfach | 1.050x | [0.942, 1.173] | 0.4115 | true |
| e2e | canonicalize | canonical-minimal | schubfach | 1.127x | [0.940, 1.357] | 0.2316 | true |
| e2e | canonicalize | control-escapes | schubfach | 1.059x | [0.920, 1.240] | 0.4778 | true |
| e2e | canonicalize | deep-64 | schubfach | 1.093x | [0.909, 1.288] | 0.3486 | true |
| e2e | canonicalize | escaped-key-order | schubfach | 1.003x | [0.877, 1.150] | 0.9630 | false |
| e2e | canonicalize | long-string | schubfach | 1.087x | [0.937, 1.249] | 0.2892 | true |
| e2e | canonicalize | nested-mixed | schubfach | 1.355x | [1.181, 1.541] | 0.0017 | true |
| e2e | canonicalize | numeric-boundary | schubfach | 1.221x | [1.038, 1.447] | 0.0360 | true |
| e2e | canonicalize | rfc-key-sorting | schubfach | 1.110x | [0.924, 1.336] | 0.3296 | true |
| e2e | canonicalize | small | schubfach | 1.046x | [0.883, 1.262] | 0.6258 | true |
| e2e | canonicalize | surrogate-pair | json-canon | 1.010x | [0.848, 1.203] | 0.9134 | false |
| e2e | canonicalize | unicode | schubfach | 1.124x | [0.949, 1.306] | 0.1833 | true |
| e2e | canonicalize | verify-whitespace | schubfach | 1.104x | [0.926, 1.310] | 0.3126 | true |
| e2e | verify | array-2048 | schubfach | 1.297x | [1.223, 1.374] | 0.0007 | true |
| e2e | verify | array-256 | schubfach | 1.222x | [1.045, 1.449] | 0.0383 | true |
| e2e | verify | canonical-minimal | schubfach | 1.098x | [0.945, 1.277] | 0.2566 | true |
| e2e | verify | control-escapes | schubfach | 1.007x | [0.845, 1.192] | 0.9320 | false |
| e2e | verify | deep-64 | json-canon | 1.085x | [0.924, 1.247] | 0.3269 | true |
| e2e | verify | escaped-key-order | schubfach | 1.138x | [0.961, 1.348] | 0.1709 | true |
| e2e | verify | long-string | schubfach | 1.018x | [0.882, 1.183] | 0.8197 | false |
| e2e | verify | nested-mixed | schubfach | 1.132x | [0.912, 1.381] | 0.2822 | true |
| e2e | verify | numeric-boundary | schubfach | 1.107x | [0.891, 1.380] | 0.3759 | true |
| e2e | verify | rfc-key-sorting | schubfach | 1.124x | [0.948, 1.328] | 0.2046 | true |
| e2e | verify | small | schubfach | 1.067x | [0.874, 1.301] | 0.5245 | true |
| e2e | verify | surrogate-pair | schubfach | 1.074x | [0.914, 1.253] | 0.4169 | true |
| e2e | verify | unicode | schubfach | 1.043x | [0.908, 1.191] | 0.5675 | true |
| e2e | verify | verify-whitespace | schubfach | 1.095x | [0.922, 1.320] | 0.3519 | true |

## benchstat Snippet

```text
# Benchstat Snapshot

Generated at: 2026-03-05T15:37:32Z

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
| array-2048 | 5.221 | 6.871 | schubfach | 1.32x |
| array-256 | 1.920 | 2.016 | schubfach | 1.05x |
| canonical-minimal | 1.391 | 1.567 | schubfach | 1.13x |
| control-escapes | 1.488 | 1.576 | schubfach | 1.06x |
| deep-64 | 1.419 | 1.551 | schubfach | 1.09x |
| escaped-key-order | 1.489 | 1.494 | schubfach | 1.00x |
| long-string | 1.669 | 1.814 | schubfach | 1.09x |
| nested-mixed | 1.230 | 1.667 | schubfach | 1.35x |
| numeric-boundary | 1.417 | 1.730 | schubfach | 1.22x |
| rfc-key-sorting | 1.417 | 1.573 | schubfach | 1.11x |
| small | 1.415 | 1.480 | schubfach | 1.05x |
| surrogate-pair | 1.408 | 1.394 | json-canon | 1.01x |
| unicode | 1.376 | 1.547 | schubfach | 1.12x |
| verify-whitespace | 1.337 | 1.476 | schubfach | 1.10x |

## CLI Verify (valid workloads)

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.373 | 6.968 | schubfach | 1.30x |
| array-256 | 1.834 | 2.241 | schubfach | 1.22x |
| canonical-minimal | 1.483 | 1.628 | schubfach | 1.10x |
| control-escapes | 1.546 | 1.557 | schubfach | 1.01x |
| deep-64 | 1.510 | 1.391 | json-canon | 1.09x |
| escaped-key-order | 1.384 | 1.575 | schubfach | 1.14x |
| long-string | 1.720 | 1.751 | schubfach | 1.02x |
| nested-mixed | 1.296 | 1.466 | schubfach | 1.13x |
| numeric-boundary | 1.398 | 1.548 | schubfach | 1.11x |
| rfc-key-sorting | 1.373 | 1.543 | schubfach | 1.12x |
| small | 1.453 | 1.550 | schubfach | 1.07x |
| surrogate-pair | 1.412 | 1.517 | schubfach | 1.07x |
| unicode | 1.503 | 1.567 | schubfach | 1.04x |
| verify-whitespace | 1.380 | 1.512 | schubfach | 1.10x |

## API Benchmarks (ns/op)

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2548995.833 | 4047979.833 | schubfach | 1.59x |
| canonicalize/array-256 | 303347.500 | 474368.500 | schubfach | 1.56x |
| canonicalize/canonical-minimal | 352.283 | 529.300 | schubfach | 1.50x |
| canonicalize/control-escapes | 1335.167 | 1251.333 | json-canon | 1.07x |
| canonicalize/deep | 36651.333 | 36688.667 | schubfach | 1.00x |
| canonicalize/deep-64 | 16236.667 | 16239.833 | schubfach | 1.00x |
| canonicalize/escaped-key-order | 707.433 | 1124.667 | schubfach | 1.59x |
| canonicalize/large | 11928619.000 | 20545683.667 | schubfach | 1.72x |
| canonicalize/long-string | 74802.167 | 73587.500 | json-canon | 1.02x |
| canonicalize/medium | 340270.333 | 581178.167 | schubfach | 1.71x |
| canonicalize/mixed-prod | 136311.000 | 178287.333 | schubfach | 1.31x |
| canonicalize/nested-mixed | 2181.833 | 3065.667 | schubfach | 1.41x |
| canonicalize/number-heavy | 13529.500 | 20854.833 | schubfach | 1.54x |
| canonicalize/numeric-boundary | 12395.833 | 18294.167 | schubfach | 1.48x |
| canonicalize/rfc-key-sorting | 3018.333 | 2967.833 | json-canon | 1.02x |
| canonicalize/small | 1243.167 | 1667.000 | schubfach | 1.34x |
| canonicalize/surrogate-pair | 647.250 | 639.950 | json-canon | 1.01x |
| canonicalize/unicode | 1548.667 | 1558.667 | schubfach | 1.01x |
| canonicalize/verify-whitespace | 906.500 | 1563.000 | schubfach | 1.72x |
| verify/canonical/array-2048 | 2469781.333 | 4029019.000 | schubfach | 1.63x |
| verify/canonical/array-256 | 317223.833 | 489779.667 | schubfach | 1.54x |
| verify/canonical/canonical-minimal | 358.083 | 548.567 | schubfach | 1.53x |
```

## Production Recommendation

Recommend `schubfach` based on statistically significant practical wins with conformance/oracle gates passing.
