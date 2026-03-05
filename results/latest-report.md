# Benchmark Report

Generated at: 2026-03-05T19:31:06Z

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
- Statistically significant wins (BH-corrected): 5
- Recommendation status: Recommend `schubfach` based on BH-significant wins with conformance/oracle gates passing.

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
| array-2048 | 5.323 | 7.422 | schubfach | 1.39x |
| array-256 | 2.017 | 2.405 | schubfach | 1.19x |
| canonical-minimal | 1.505 | 1.645 | schubfach | 1.09x |
| control-escapes | 1.426 | 1.498 | schubfach | 1.05x |
| deep-64 | 1.456 | 1.633 | schubfach | 1.12x |
| escaped-key-order | 1.343 | 1.631 | schubfach | 1.21x |
| long-string | 1.571 | 1.746 | schubfach | 1.11x |
| nested-mixed | 1.552 | 1.482 | json-canon | 1.05x |
| numeric-boundary | 1.588 | 1.539 | json-canon | 1.03x |
| rfc-key-sorting | 1.446 | 1.501 | schubfach | 1.04x |
| small | 1.428 | 1.558 | schubfach | 1.09x |
| surrogate-pair | 1.403 | 1.743 | schubfach | 1.24x |
| unicode | 1.505 | 1.551 | schubfach | 1.03x |
| verify-whitespace | 1.533 | 1.587 | schubfach | 1.04x |

### CLI verify

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.767 | 7.231 | schubfach | 1.25x |
| array-256 | 2.020 | 2.257 | schubfach | 1.12x |
| canonical-minimal | 1.567 | 1.477 | json-canon | 1.06x |
| control-escapes | 1.315 | 1.666 | schubfach | 1.27x |
| deep-64 | 1.577 | 1.654 | schubfach | 1.05x |
| escaped-key-order | 1.379 | 1.568 | schubfach | 1.14x |
| long-string | 1.576 | 1.725 | schubfach | 1.09x |
| nested-mixed | 1.551 | 1.643 | schubfach | 1.06x |
| numeric-boundary | 1.316 | 1.588 | schubfach | 1.21x |
| rfc-key-sorting | 1.461 | 1.516 | schubfach | 1.04x |
| small | 1.481 | 1.588 | schubfach | 1.07x |
| surrogate-pair | 1.389 | 1.529 | schubfach | 1.10x |
| unicode | 1.556 | 1.654 | schubfach | 1.06x |
| verify-whitespace | 1.521 | 1.592 | schubfach | 1.05x |

### API

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2546054.100 | 4070709.400 | schubfach | 1.60x |
| canonicalize/array-256 | 305012.300 | 472576.100 | schubfach | 1.55x |
| canonicalize/canonical-minimal | 351.080 | 530.610 | schubfach | 1.51x |
| canonicalize/control-escapes | 1327.300 | 1258.300 | json-canon | 1.05x |
| canonicalize/deep | 36582.000 | 36457.100 | json-canon | 1.00x |
| canonicalize/deep-64 | 16199.400 | 16189.300 | json-canon | 1.00x |
| canonicalize/escaped-key-order | 706.400 | 1128.100 | schubfach | 1.60x |
| canonicalize/large | 11878471.100 | 20512477.000 | schubfach | 1.73x |
| canonicalize/long-string | 75928.400 | 75314.300 | json-canon | 1.01x |
| canonicalize/medium | 343001.100 | 578836.200 | schubfach | 1.69x |
| canonicalize/mixed-prod | 136231.600 | 178649.500 | schubfach | 1.31x |
| canonicalize/nested-mixed | 2181.200 | 3052.700 | schubfach | 1.40x |
| canonicalize/number-heavy | 13773.900 | 21318.100 | schubfach | 1.55x |
| canonicalize/numeric-boundary | 12540.300 | 18802.300 | schubfach | 1.50x |
| canonicalize/rfc-key-sorting | 2997.900 | 2949.800 | json-canon | 1.02x |
| canonicalize/small | 1239.200 | 1664.800 | schubfach | 1.34x |
| canonicalize/surrogate-pair | 649.320 | 638.010 | json-canon | 1.02x |
| canonicalize/unicode | 1549.700 | 1541.800 | json-canon | 1.01x |
| canonicalize/verify-whitespace | 908.960 | 1552.400 | schubfach | 1.71x |
| verify/canonical/array-2048 | 2485566.900 | 4009213.800 | schubfach | 1.61x |
| verify/canonical/array-256 | 316798.500 | 481006.700 | schubfach | 1.52x |
| verify/canonical/canonical-minimal | 357.640 | 542.100 | schubfach | 1.52x |
| verify/canonical/control-escapes | 1316.500 | 1258.600 | json-canon | 1.05x |
| verify/canonical/deep | 37680.400 | 37698.500 | schubfach | 1.00x |
| verify/canonical/deep-64 | 16495.500 | 16636.000 | schubfach | 1.01x |
| verify/canonical/escaped-key-order | 685.910 | 1116.900 | schubfach | 1.63x |
| verify/canonical/large | 11462071.700 | 19965241.200 | schubfach | 1.74x |
| verify/canonical/long-string | 76711.000 | 76997.500 | schubfach | 1.00x |
| verify/canonical/medium | 341590.700 | 579505.300 | schubfach | 1.70x |
| verify/canonical/mixed-prod | 133901.800 | 176167.700 | schubfach | 1.32x |
| verify/canonical/nested-mixed | 2206.400 | 3099.200 | schubfach | 1.40x |
| verify/canonical/number-heavy | 14056.900 | 21465.300 | schubfach | 1.53x |
| verify/canonical/numeric-boundary | 12751.200 | 18892.500 | schubfach | 1.48x |
| verify/canonical/rfc-key-sorting | 2844.900 | 2823.000 | json-canon | 1.01x |
| verify/canonical/small | 1224.700 | 1647.900 | schubfach | 1.35x |
| verify/canonical/surrogate-pair | 654.420 | 653.760 | json-canon | 1.00x |
| verify/canonical/unicode | 1547.500 | 1546.200 | json-canon | 1.00x |
| verify/canonical/verify-whitespace | 932.960 | 1587.800 | schubfach | 1.70x |
| verify/noncanonical/control-escapes | 1316.300 | 1258.200 | json-canon | 1.05x |
| verify/noncanonical/escaped-key-order | 703.630 | 1125.800 | schubfach | 1.60x |
| verify/noncanonical/large | 11961812.000 | 20563299.000 | schubfach | 1.72x |
| verify/noncanonical/medium | 339688.400 | 579647.300 | schubfach | 1.71x |
| verify/noncanonical/mixed-prod | 136231.100 | 179182.700 | schubfach | 1.32x |
| verify/noncanonical/nested-mixed | 2177.200 | 3087.400 | schubfach | 1.42x |
| verify/noncanonical/number-heavy | 13791.400 | 21478.800 | schubfach | 1.56x |
| verify/noncanonical/numeric-boundary | 12643.900 | 18725.900 | schubfach | 1.48x |
| verify/noncanonical/rfc-key-sorting | 2982.400 | 2985.300 | schubfach | 1.00x |
| verify/noncanonical/small | 1237.000 | 1685.900 | schubfach | 1.36x |
| verify/noncanonical/surrogate-pair | 645.510 | 647.550 | schubfach | 1.00x |
| verify/noncanonical/unicode | 1543.100 | 1563.500 | schubfach | 1.01x |
| verify/noncanonical/verify-whitespace | 908.070 | 1570.800 | schubfach | 1.73x |

## Statistical Inference

| track | mode | workload | winner | speedup | ci95 | p-value | p-adj | sig-BH |
|---|---|---|---|---:|---|---:|---:|---|
| e2e | canonicalize | array-2048 | schubfach | 1.394x | [1.302, 1.504] | 0.0002 | 0.0028 | true |
| e2e | canonicalize | array-256 | schubfach | 1.193x | [1.057, 1.344] | 0.0120 | 0.0560 | false |
| e2e | canonicalize | canonical-minimal | schubfach | 1.093x | [0.967, 1.223] | 0.1606 | 0.3549 | false |
| e2e | canonicalize | control-escapes | schubfach | 1.051x | [0.915, 1.226] | 0.5191 | 0.6241 | false |
| e2e | canonicalize | deep-64 | schubfach | 1.122x | [0.994, 1.269] | 0.0816 | 0.2643 | false |
| e2e | canonicalize | escaped-key-order | schubfach | 1.214x | [1.065, 1.360] | 0.0062 | 0.0347 | true |
| e2e | canonicalize | long-string | schubfach | 1.112x | [0.977, 1.259] | 0.1310 | 0.3334 | false |
| e2e | canonicalize | nested-mixed | json-canon | 1.048x | [0.909, 1.179] | 0.4913 | 0.6241 | false |
| e2e | canonicalize | numeric-boundary | json-canon | 1.032x | [0.904, 1.153] | 0.6193 | 0.6241 | false |
| e2e | canonicalize | rfc-key-sorting | schubfach | 1.038x | [0.919, 1.198] | 0.5915 | 0.6241 | false |
| e2e | canonicalize | small | schubfach | 1.091x | [0.951, 1.259] | 0.2515 | 0.4729 | false |
| e2e | canonicalize | surrogate-pair | schubfach | 1.243x | [1.110, 1.358] | 0.0006 | 0.0056 | true |
| e2e | canonicalize | unicode | schubfach | 1.031x | [0.917, 1.162] | 0.6161 | 0.6241 | false |
| e2e | canonicalize | verify-whitespace | schubfach | 1.035x | [0.914, 1.173] | 0.6069 | 0.6241 | false |
| e2e | verify | array-2048 | schubfach | 1.254x | [1.167, 1.343] | 0.0002 | 0.0028 | true |
| e2e | verify | array-256 | schubfach | 1.118x | [0.984, 1.254] | 0.0944 | 0.2643 | false |
| e2e | verify | canonical-minimal | json-canon | 1.061x | [0.926, 1.200] | 0.3881 | 0.5720 | false |
| e2e | verify | control-escapes | schubfach | 1.267x | [1.104, 1.422] | 0.0026 | 0.0182 | true |
| e2e | verify | deep-64 | schubfach | 1.049x | [0.938, 1.183] | 0.4273 | 0.5982 | false |
| e2e | verify | escaped-key-order | schubfach | 1.137x | [0.983, 1.293] | 0.0890 | 0.2643 | false |
| e2e | verify | long-string | schubfach | 1.095x | [0.938, 1.264] | 0.2533 | 0.4729 | false |
| e2e | verify | nested-mixed | schubfach | 1.059x | [0.940, 1.202] | 0.3775 | 0.5720 | false |
| e2e | verify | numeric-boundary | schubfach | 1.207x | [1.032, 1.394] | 0.0286 | 0.1144 | false |
| e2e | verify | rfc-key-sorting | schubfach | 1.038x | [0.901, 1.201] | 0.6241 | 0.6241 | false |
| e2e | verify | small | schubfach | 1.072x | [0.945, 1.209] | 0.2915 | 0.4881 | false |
| e2e | verify | surrogate-pair | schubfach | 1.101x | [0.970, 1.252] | 0.1648 | 0.3549 | false |
| e2e | verify | unicode | schubfach | 1.063x | [0.953, 1.188] | 0.2963 | 0.4881 | false |
| e2e | verify | verify-whitespace | schubfach | 1.047x | [0.927, 1.198] | 0.5007 | 0.6241 | false |

## benchstat Snippet

```text
# Benchstat Snapshot

Generated at: 2026-03-05T19:28:39Z

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
| array-2048 | 5.323 | 7.422 | schubfach | 1.39x |
| array-256 | 2.017 | 2.405 | schubfach | 1.19x |
| canonical-minimal | 1.505 | 1.645 | schubfach | 1.09x |
| control-escapes | 1.426 | 1.498 | schubfach | 1.05x |
| deep-64 | 1.456 | 1.633 | schubfach | 1.12x |
| escaped-key-order | 1.343 | 1.631 | schubfach | 1.21x |
| long-string | 1.571 | 1.746 | schubfach | 1.11x |
| nested-mixed | 1.552 | 1.482 | json-canon | 1.05x |
| numeric-boundary | 1.588 | 1.539 | json-canon | 1.03x |
| rfc-key-sorting | 1.446 | 1.501 | schubfach | 1.04x |
| small | 1.428 | 1.558 | schubfach | 1.09x |
| surrogate-pair | 1.403 | 1.743 | schubfach | 1.24x |
| unicode | 1.505 | 1.551 | schubfach | 1.03x |
| verify-whitespace | 1.533 | 1.587 | schubfach | 1.04x |

## CLI Verify (valid workloads)

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.767 | 7.231 | schubfach | 1.25x |
| array-256 | 2.020 | 2.257 | schubfach | 1.12x |
| canonical-minimal | 1.567 | 1.477 | json-canon | 1.06x |
| control-escapes | 1.315 | 1.666 | schubfach | 1.27x |
| deep-64 | 1.577 | 1.654 | schubfach | 1.05x |
| escaped-key-order | 1.379 | 1.568 | schubfach | 1.14x |
| long-string | 1.576 | 1.725 | schubfach | 1.09x |
| nested-mixed | 1.551 | 1.643 | schubfach | 1.06x |
| numeric-boundary | 1.316 | 1.588 | schubfach | 1.21x |
| rfc-key-sorting | 1.461 | 1.516 | schubfach | 1.04x |
| small | 1.481 | 1.588 | schubfach | 1.07x |
| surrogate-pair | 1.389 | 1.529 | schubfach | 1.10x |
| unicode | 1.556 | 1.654 | schubfach | 1.06x |
| verify-whitespace | 1.521 | 1.592 | schubfach | 1.05x |

## API Benchmarks (ns/op)

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2546054.100 | 4070709.400 | schubfach | 1.60x |
| canonicalize/array-256 | 305012.300 | 472576.100 | schubfach | 1.55x |
| canonicalize/canonical-minimal | 351.080 | 530.610 | schubfach | 1.51x |
| canonicalize/control-escapes | 1327.300 | 1258.300 | json-canon | 1.05x |
| canonicalize/deep | 36582.000 | 36457.100 | json-canon | 1.00x |
| canonicalize/deep-64 | 16199.400 | 16189.300 | json-canon | 1.00x |
| canonicalize/escaped-key-order | 706.400 | 1128.100 | schubfach | 1.60x |
| canonicalize/large | 11878471.100 | 20512477.000 | schubfach | 1.73x |
| canonicalize/long-string | 75928.400 | 75314.300 | json-canon | 1.01x |
| canonicalize/medium | 343001.100 | 578836.200 | schubfach | 1.69x |
| canonicalize/mixed-prod | 136231.600 | 178649.500 | schubfach | 1.31x |
| canonicalize/nested-mixed | 2181.200 | 3052.700 | schubfach | 1.40x |
| canonicalize/number-heavy | 13773.900 | 21318.100 | schubfach | 1.55x |
| canonicalize/numeric-boundary | 12540.300 | 18802.300 | schubfach | 1.50x |
| canonicalize/rfc-key-sorting | 2997.900 | 2949.800 | json-canon | 1.02x |
| canonicalize/small | 1239.200 | 1664.800 | schubfach | 1.34x |
| canonicalize/surrogate-pair | 649.320 | 638.010 | json-canon | 1.02x |
| canonicalize/unicode | 1549.700 | 1541.800 | json-canon | 1.01x |
| canonicalize/verify-whitespace | 908.960 | 1552.400 | schubfach | 1.71x |
| verify/canonical/array-2048 | 2485566.900 | 4009213.800 | schubfach | 1.61x |
| verify/canonical/array-256 | 316798.500 | 481006.700 | schubfach | 1.52x |
| verify/canonical/canonical-minimal | 357.640 | 542.100 | schubfach | 1.52x |
```

## Production Recommendation

Recommend `schubfach` based on BH-significant wins with conformance/oracle gates passing.
