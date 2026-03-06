# Benchmark Report

Generated at: 2026-03-06T00:19:59Z

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
- Statistically significant wins (BH-corrected): 204
- Recommendation status: Recommend `schubfach-rs` based on BH-significant wins with conformance/oracle gates passing.

## Conformance Evidence

- cyberphone: total=72 passed=72 failed=0
- lab-workload: total=232 passed=232 failed=0
- rfc8785: total=4 passed=4 failed=0

## Quality Findings

- determinism_failures: none
- output_equality_failures: none
- invalid_failure_parity_issues: none
- oracle_mismatches: none

## Performance Winners by Workload

### CLI canonicalize

#### go:schubfach-vs-bd

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.466 | 7.063 | schubfach | 1.29x |
| array-256 | 1.923 | 2.167 | schubfach | 1.13x |
| canonical-minimal | 1.510 | 1.565 | schubfach | 1.04x |
| control-escapes | 1.380 | 1.673 | schubfach | 1.21x |
| deep-64 | 1.467 | 1.638 | schubfach | 1.12x |
| escaped-key-order | 1.549 | 1.531 | json-canon | 1.01x |
| long-string | 1.409 | 1.692 | schubfach | 1.20x |
| nested-mixed | 1.511 | 1.497 | json-canon | 1.01x |
| numeric-boundary | 1.537 | 1.706 | schubfach | 1.11x |
| rfc-key-sorting | 1.561 | 1.608 | schubfach | 1.03x |
| small | 1.450 | 1.449 | json-canon | 1.00x |
| surrogate-pair | 1.510 | 1.524 | schubfach | 1.01x |
| unicode | 1.597 | 1.470 | json-canon | 1.09x |
| verify-whitespace | 1.455 | 1.614 | schubfach | 1.11x |

#### rs:schubfach-vs-bd

| workload | schubfach-rs (avg_ms) | json-canon-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 2.606 | 4.584 | schubfach-rs | 1.76x |
| array-256 | 0.828 | 1.152 | schubfach-rs | 1.39x |
| canonical-minimal | 0.546 | 0.582 | schubfach-rs | 1.07x |
| control-escapes | 0.531 | 0.594 | schubfach-rs | 1.12x |
| deep-64 | 0.564 | 0.647 | schubfach-rs | 1.15x |
| escaped-key-order | 0.573 | 0.712 | schubfach-rs | 1.24x |
| long-string | 0.605 | 0.636 | schubfach-rs | 1.05x |
| nested-mixed | 0.555 | 0.644 | schubfach-rs | 1.16x |
| numeric-boundary | 0.543 | 0.738 | schubfach-rs | 1.36x |
| rfc-key-sorting | 0.535 | 0.588 | schubfach-rs | 1.10x |
| small | 0.530 | 0.642 | schubfach-rs | 1.21x |
| surrogate-pair | 0.536 | 0.603 | schubfach-rs | 1.12x |
| unicode | 0.534 | 0.612 | schubfach-rs | 1.15x |
| verify-whitespace | 0.559 | 0.691 | schubfach-rs | 1.24x |

#### schubfach:go-vs-rs

| workload | schubfach (avg_ms) | schubfach-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.466 | 2.606 | schubfach-rs | 2.10x |
| array-256 | 1.923 | 0.828 | schubfach-rs | 2.32x |
| canonical-minimal | 1.510 | 0.546 | schubfach-rs | 2.77x |
| control-escapes | 1.380 | 0.531 | schubfach-rs | 2.60x |
| deep-64 | 1.467 | 0.564 | schubfach-rs | 2.60x |
| escaped-key-order | 1.549 | 0.573 | schubfach-rs | 2.70x |
| long-string | 1.409 | 0.605 | schubfach-rs | 2.33x |
| nested-mixed | 1.511 | 0.555 | schubfach-rs | 2.72x |
| numeric-boundary | 1.537 | 0.543 | schubfach-rs | 2.83x |
| rfc-key-sorting | 1.561 | 0.535 | schubfach-rs | 2.92x |
| small | 1.450 | 0.530 | schubfach-rs | 2.73x |
| surrogate-pair | 1.510 | 0.536 | schubfach-rs | 2.82x |
| unicode | 1.597 | 0.534 | schubfach-rs | 2.99x |
| verify-whitespace | 1.455 | 0.559 | schubfach-rs | 2.61x |

#### bd:go-vs-rs

| workload | json-canon (avg_ms) | json-canon-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 7.063 | 4.584 | json-canon-rs | 1.54x |
| array-256 | 2.167 | 1.152 | json-canon-rs | 1.88x |
| canonical-minimal | 1.565 | 0.582 | json-canon-rs | 2.69x |
| control-escapes | 1.673 | 0.594 | json-canon-rs | 2.82x |
| deep-64 | 1.638 | 0.647 | json-canon-rs | 2.53x |
| escaped-key-order | 1.531 | 0.712 | json-canon-rs | 2.15x |
| long-string | 1.692 | 0.636 | json-canon-rs | 2.66x |
| nested-mixed | 1.497 | 0.644 | json-canon-rs | 2.32x |
| numeric-boundary | 1.706 | 0.738 | json-canon-rs | 2.31x |
| rfc-key-sorting | 1.608 | 0.588 | json-canon-rs | 2.73x |
| small | 1.449 | 0.642 | json-canon-rs | 2.26x |
| surrogate-pair | 1.524 | 0.603 | json-canon-rs | 2.53x |
| unicode | 1.470 | 0.612 | json-canon-rs | 2.40x |
| verify-whitespace | 1.614 | 0.691 | json-canon-rs | 2.34x |

### CLI verify

#### go:schubfach-vs-bd

| workload | schubfach (avg_ms) | json-canon (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.623 | 7.128 | schubfach | 1.27x |
| array-256 | 1.834 | 2.355 | schubfach | 1.28x |
| canonical-minimal | 1.386 | 1.389 | schubfach | 1.00x |
| control-escapes | 1.397 | 1.588 | schubfach | 1.14x |
| deep-64 | 1.484 | 1.798 | schubfach | 1.21x |
| escaped-key-order | 1.436 | 1.551 | schubfach | 1.08x |
| long-string | 1.659 | 1.718 | schubfach | 1.04x |
| nested-mixed | 1.453 | 1.676 | schubfach | 1.15x |
| numeric-boundary | 1.432 | 1.599 | schubfach | 1.12x |
| rfc-key-sorting | 1.493 | 1.698 | schubfach | 1.14x |
| small | 1.402 | 1.702 | schubfach | 1.21x |
| surrogate-pair | 1.448 | 1.625 | schubfach | 1.12x |
| unicode | 1.400 | 1.663 | schubfach | 1.19x |
| verify-whitespace | 1.448 | 1.570 | schubfach | 1.08x |

#### rs:schubfach-vs-bd

| workload | schubfach-rs (avg_ms) | json-canon-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 2.419 | 4.146 | schubfach-rs | 1.71x |
| array-256 | 0.821 | 1.217 | schubfach-rs | 1.48x |
| canonical-minimal | 0.557 | 0.581 | schubfach-rs | 1.04x |
| control-escapes | 0.537 | 0.617 | schubfach-rs | 1.15x |
| deep-64 | 0.561 | 0.609 | schubfach-rs | 1.08x |
| escaped-key-order | 0.553 | 0.645 | schubfach-rs | 1.17x |
| long-string | 0.612 | 0.647 | schubfach-rs | 1.06x |
| nested-mixed | 0.546 | 0.639 | schubfach-rs | 1.17x |
| numeric-boundary | 0.545 | 0.715 | schubfach-rs | 1.31x |
| rfc-key-sorting | 0.534 | 0.601 | schubfach-rs | 1.12x |
| small | 0.546 | 0.708 | schubfach-rs | 1.30x |
| surrogate-pair | 0.533 | 0.578 | schubfach-rs | 1.08x |
| unicode | 0.546 | 0.629 | schubfach-rs | 1.15x |
| verify-whitespace | 0.554 | 0.688 | schubfach-rs | 1.24x |

#### schubfach:go-vs-rs

| workload | schubfach (avg_ms) | schubfach-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 5.623 | 2.419 | schubfach-rs | 2.32x |
| array-256 | 1.834 | 0.821 | schubfach-rs | 2.23x |
| canonical-minimal | 1.386 | 0.557 | schubfach-rs | 2.49x |
| control-escapes | 1.397 | 0.537 | schubfach-rs | 2.60x |
| deep-64 | 1.484 | 0.561 | schubfach-rs | 2.64x |
| escaped-key-order | 1.436 | 0.553 | schubfach-rs | 2.60x |
| long-string | 1.659 | 0.612 | schubfach-rs | 2.71x |
| nested-mixed | 1.453 | 0.546 | schubfach-rs | 2.66x |
| numeric-boundary | 1.432 | 0.545 | schubfach-rs | 2.63x |
| rfc-key-sorting | 1.493 | 0.534 | schubfach-rs | 2.79x |
| small | 1.402 | 0.546 | schubfach-rs | 2.56x |
| surrogate-pair | 1.448 | 0.533 | schubfach-rs | 2.72x |
| unicode | 1.400 | 0.546 | schubfach-rs | 2.56x |
| verify-whitespace | 1.448 | 0.554 | schubfach-rs | 2.62x |

#### bd:go-vs-rs

| workload | json-canon (avg_ms) | json-canon-rs (avg_ms) | winner | speedup |
|---|---:|---:|---|---:|
| array-2048 | 7.128 | 4.146 | json-canon-rs | 1.72x |
| array-256 | 2.355 | 1.217 | json-canon-rs | 1.93x |
| canonical-minimal | 1.389 | 0.581 | json-canon-rs | 2.39x |
| control-escapes | 1.588 | 0.617 | json-canon-rs | 2.57x |
| deep-64 | 1.798 | 0.609 | json-canon-rs | 2.95x |
| escaped-key-order | 1.551 | 0.645 | json-canon-rs | 2.40x |
| long-string | 1.718 | 0.647 | json-canon-rs | 2.66x |
| nested-mixed | 1.676 | 0.639 | json-canon-rs | 2.62x |
| numeric-boundary | 1.599 | 0.715 | json-canon-rs | 2.23x |
| rfc-key-sorting | 1.698 | 0.601 | json-canon-rs | 2.83x |
| small | 1.702 | 0.708 | json-canon-rs | 2.40x |
| surrogate-pair | 1.625 | 0.578 | json-canon-rs | 2.81x |
| unicode | 1.663 | 0.629 | json-canon-rs | 2.65x |
| verify-whitespace | 1.570 | 0.688 | json-canon-rs | 2.28x |

### API

#### go:schubfach-vs-bd

| workload | schubfach (ns/op) | json-canon (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2541751.700 | 4077595.500 | schubfach | 1.60x |
| canonicalize/array-256 | 305466.600 | 473477.800 | schubfach | 1.55x |
| canonicalize/canonical-minimal | 351.300 | 531.540 | schubfach | 1.51x |
| canonicalize/control-escapes | 1339.000 | 1257.500 | json-canon | 1.06x |
| canonicalize/deep | 36725.400 | 36748.900 | schubfach | 1.00x |
| canonicalize/deep-64 | 16230.500 | 16230.700 | schubfach | 1.00x |
| canonicalize/escaped-key-order | 707.640 | 1129.100 | schubfach | 1.60x |
| canonicalize/large | 11923742.900 | 20570122.400 | schubfach | 1.73x |
| canonicalize/long-string | 75830.600 | 75638.600 | json-canon | 1.00x |
| canonicalize/medium | 341106.000 | 582942.000 | schubfach | 1.71x |
| canonicalize/mixed-prod | 136351.900 | 180492.400 | schubfach | 1.32x |
| canonicalize/nested-mixed | 2184.200 | 3088.800 | schubfach | 1.41x |
| canonicalize/number-heavy | 13807.300 | 21411.800 | schubfach | 1.55x |
| canonicalize/numeric-boundary | 12692.700 | 18815.000 | schubfach | 1.48x |
| canonicalize/rfc-key-sorting | 3003.600 | 2972.700 | json-canon | 1.01x |
| canonicalize/small | 1251.400 | 1674.800 | schubfach | 1.34x |
| canonicalize/surrogate-pair | 648.030 | 657.330 | schubfach | 1.01x |
| canonicalize/unicode | 1554.100 | 1550.700 | json-canon | 1.00x |
| canonicalize/verify-whitespace | 910.410 | 1568.600 | schubfach | 1.72x |
| verify/canonical/array-2048 | 2480515.600 | 4028940.800 | schubfach | 1.62x |
| verify/canonical/array-256 | 316572.100 | 485478.800 | schubfach | 1.53x |
| verify/canonical/canonical-minimal | 359.820 | 547.650 | schubfach | 1.52x |
| verify/canonical/control-escapes | 1327.400 | 1271.300 | json-canon | 1.04x |
| verify/canonical/deep | 37855.400 | 38276.600 | schubfach | 1.01x |
| verify/canonical/deep-64 | 16778.500 | 16830.100 | schubfach | 1.00x |
| verify/canonical/escaped-key-order | 690.010 | 1130.100 | schubfach | 1.64x |
| verify/canonical/large | 11539381.200 | 20021439.300 | schubfach | 1.74x |
| verify/canonical/long-string | 77271.600 | 77030.400 | json-canon | 1.00x |
| verify/canonical/medium | 342086.900 | 578289.700 | schubfach | 1.69x |
| verify/canonical/mixed-prod | 133360.500 | 176062.500 | schubfach | 1.32x |
| verify/canonical/nested-mixed | 2219.100 | 3107.200 | schubfach | 1.40x |
| verify/canonical/number-heavy | 13890.000 | 21550.700 | schubfach | 1.55x |
| verify/canonical/numeric-boundary | 12793.500 | 18861.100 | schubfach | 1.47x |
| verify/canonical/rfc-key-sorting | 2856.600 | 2843.200 | json-canon | 1.00x |
| verify/canonical/small | 1223.400 | 1655.800 | schubfach | 1.35x |
| verify/canonical/surrogate-pair | 657.530 | 652.600 | json-canon | 1.01x |
| verify/canonical/unicode | 1548.400 | 1546.800 | json-canon | 1.00x |
| verify/canonical/verify-whitespace | 933.770 | 1588.400 | schubfach | 1.70x |
| verify/noncanonical/control-escapes | 1320.500 | 1260.500 | json-canon | 1.05x |
| verify/noncanonical/escaped-key-order | 704.590 | 1129.600 | schubfach | 1.60x |
| verify/noncanonical/large | 11963796.900 | 20568032.000 | schubfach | 1.72x |
| verify/noncanonical/medium | 344752.500 | 575932.800 | schubfach | 1.67x |
| verify/noncanonical/mixed-prod | 136719.000 | 177614.400 | schubfach | 1.30x |
| verify/noncanonical/nested-mixed | 2200.300 | 3067.300 | schubfach | 1.39x |
| verify/noncanonical/number-heavy | 13876.400 | 21344.300 | schubfach | 1.54x |
| verify/noncanonical/numeric-boundary | 12722.300 | 18694.100 | schubfach | 1.47x |
| verify/noncanonical/rfc-key-sorting | 3023.600 | 2958.200 | json-canon | 1.02x |
| verify/noncanonical/small | 1250.000 | 1673.900 | schubfach | 1.34x |
| verify/noncanonical/surrogate-pair | 656.450 | 644.390 | json-canon | 1.02x |
| verify/noncanonical/unicode | 1561.000 | 1547.000 | json-canon | 1.01x |
| verify/noncanonical/verify-whitespace | 918.810 | 1562.800 | schubfach | 1.70x |

#### rs:schubfach-vs-bd

| workload | schubfach-rs (ns/op) | json-canon-rs (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 1185767.500 | 2882257.500 | schubfach-rs | 2.43x |
| canonicalize/array-256 | 157030.100 | 321393.600 | schubfach-rs | 2.05x |
| canonicalize/canonical-minimal | 190.500 | 414.600 | schubfach-rs | 2.18x |
| canonicalize/control-escapes | 606.900 | 607.000 | schubfach-rs | 1.00x |
| canonicalize/deep-64 | 9380.100 | 9184.700 | json-canon-rs | 1.02x |
| canonicalize/escaped-key-order | 298.200 | 746.400 | schubfach-rs | 2.50x |
| canonicalize/long-string | 17259.300 | 17278.000 | schubfach-rs | 1.00x |
| canonicalize/nested-mixed | 842.800 | 1684.300 | schubfach-rs | 2.00x |
| canonicalize/numeric-boundary | 976.500 | 8974.700 | schubfach-rs | 9.19x |
| canonicalize/rfc-key-sorting | 1294.400 | 1250.200 | json-canon-rs | 1.04x |
| canonicalize/small | 583.200 | 1000.100 | schubfach-rs | 1.71x |
| canonicalize/surrogate-pair | 267.100 | 280.600 | schubfach-rs | 1.05x |
| canonicalize/unicode | 1676.400 | 1964.500 | schubfach-rs | 1.17x |
| canonicalize/verify-whitespace | 347.000 | 980.500 | schubfach-rs | 2.83x |
| verify/canonical/array-2048 | 1208080.100 | 2906957.700 | schubfach-rs | 2.41x |
| verify/canonical/array-256 | 141609.300 | 309810.800 | schubfach-rs | 2.19x |
| verify/canonical/canonical-minimal | 174.800 | 413.100 | schubfach-rs | 2.36x |
| verify/canonical/control-escapes | 612.500 | 606.500 | json-canon-rs | 1.01x |
| verify/canonical/deep-64 | 9153.500 | 9103.900 | json-canon-rs | 1.01x |
| verify/canonical/escaped-key-order | 285.900 | 724.500 | schubfach-rs | 2.53x |
| verify/canonical/long-string | 17384.500 | 17325.000 | json-canon-rs | 1.00x |
| verify/canonical/nested-mixed | 848.200 | 1641.000 | schubfach-rs | 1.93x |
| verify/canonical/numeric-boundary | 961.400 | 8740.100 | schubfach-rs | 9.09x |
| verify/canonical/rfc-key-sorting | 1312.100 | 1321.200 | schubfach-rs | 1.01x |
| verify/canonical/small | 602.000 | 1013.000 | schubfach-rs | 1.68x |
| verify/canonical/surrogate-pair | 282.600 | 297.000 | schubfach-rs | 1.05x |
| verify/canonical/unicode | 1696.500 | 1631.400 | json-canon-rs | 1.04x |
| verify/canonical/verify-whitespace | 363.500 | 978.200 | schubfach-rs | 2.69x |

#### schubfach:go-vs-rs

| workload | schubfach (ns/op) | schubfach-rs (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 2541751.700 | 1185767.500 | schubfach-rs | 2.14x |
| canonicalize/array-256 | 305466.600 | 157030.100 | schubfach-rs | 1.95x |
| canonicalize/canonical-minimal | 351.300 | 190.500 | schubfach-rs | 1.84x |
| canonicalize/control-escapes | 1339.000 | 606.900 | schubfach-rs | 2.21x |
| canonicalize/deep-64 | 16230.500 | 9380.100 | schubfach-rs | 1.73x |
| canonicalize/escaped-key-order | 707.640 | 298.200 | schubfach-rs | 2.37x |
| canonicalize/long-string | 75830.600 | 17259.300 | schubfach-rs | 4.39x |
| canonicalize/nested-mixed | 2184.200 | 842.800 | schubfach-rs | 2.59x |
| canonicalize/numeric-boundary | 12692.700 | 976.500 | schubfach-rs | 13.00x |
| canonicalize/rfc-key-sorting | 3003.600 | 1294.400 | schubfach-rs | 2.32x |
| canonicalize/small | 1251.400 | 583.200 | schubfach-rs | 2.15x |
| canonicalize/surrogate-pair | 648.030 | 267.100 | schubfach-rs | 2.43x |
| canonicalize/unicode | 1554.100 | 1676.400 | schubfach | 1.08x |
| canonicalize/verify-whitespace | 910.410 | 347.000 | schubfach-rs | 2.62x |
| verify/canonical/array-2048 | 2480515.600 | 1208080.100 | schubfach-rs | 2.05x |
| verify/canonical/array-256 | 316572.100 | 141609.300 | schubfach-rs | 2.24x |
| verify/canonical/canonical-minimal | 359.820 | 174.800 | schubfach-rs | 2.06x |
| verify/canonical/control-escapes | 1327.400 | 612.500 | schubfach-rs | 2.17x |
| verify/canonical/deep-64 | 16778.500 | 9153.500 | schubfach-rs | 1.83x |
| verify/canonical/escaped-key-order | 690.010 | 285.900 | schubfach-rs | 2.41x |
| verify/canonical/long-string | 77271.600 | 17384.500 | schubfach-rs | 4.44x |
| verify/canonical/nested-mixed | 2219.100 | 848.200 | schubfach-rs | 2.62x |
| verify/canonical/numeric-boundary | 12793.500 | 961.400 | schubfach-rs | 13.31x |
| verify/canonical/rfc-key-sorting | 2856.600 | 1312.100 | schubfach-rs | 2.18x |
| verify/canonical/small | 1223.400 | 602.000 | schubfach-rs | 2.03x |
| verify/canonical/surrogate-pair | 657.530 | 282.600 | schubfach-rs | 2.33x |
| verify/canonical/unicode | 1548.400 | 1696.500 | schubfach | 1.10x |
| verify/canonical/verify-whitespace | 933.770 | 363.500 | schubfach-rs | 2.57x |

#### bd:go-vs-rs

| workload | json-canon (ns/op) | json-canon-rs (ns/op) | winner | speedup |
|---|---:|---:|---|---:|
| canonicalize/array-2048 | 4077595.500 | 2882257.500 | json-canon-rs | 1.41x |
| canonicalize/array-256 | 473477.800 | 321393.600 | json-canon-rs | 1.47x |
| canonicalize/canonical-minimal | 531.540 | 414.600 | json-canon-rs | 1.28x |
| canonicalize/control-escapes | 1257.500 | 607.000 | json-canon-rs | 2.07x |
| canonicalize/deep-64 | 16230.700 | 9184.700 | json-canon-rs | 1.77x |
| canonicalize/escaped-key-order | 1129.100 | 746.400 | json-canon-rs | 1.51x |
| canonicalize/long-string | 75638.600 | 17278.000 | json-canon-rs | 4.38x |
| canonicalize/nested-mixed | 3088.800 | 1684.300 | json-canon-rs | 1.83x |
| canonicalize/numeric-boundary | 18815.000 | 8974.700 | json-canon-rs | 2.10x |
| canonicalize/rfc-key-sorting | 2972.700 | 1250.200 | json-canon-rs | 2.38x |
| canonicalize/small | 1674.800 | 1000.100 | json-canon-rs | 1.67x |
| canonicalize/surrogate-pair | 657.330 | 280.600 | json-canon-rs | 2.34x |
| canonicalize/unicode | 1550.700 | 1964.500 | json-canon | 1.27x |
| canonicalize/verify-whitespace | 1568.600 | 980.500 | json-canon-rs | 1.60x |
| verify/canonical/array-2048 | 4028940.800 | 2906957.700 | json-canon-rs | 1.39x |
| verify/canonical/array-256 | 485478.800 | 309810.800 | json-canon-rs | 1.57x |
| verify/canonical/canonical-minimal | 547.650 | 413.100 | json-canon-rs | 1.33x |
| verify/canonical/control-escapes | 1271.300 | 606.500 | json-canon-rs | 2.10x |
| verify/canonical/deep-64 | 16830.100 | 9103.900 | json-canon-rs | 1.85x |
| verify/canonical/escaped-key-order | 1130.100 | 724.500 | json-canon-rs | 1.56x |
| verify/canonical/long-string | 77030.400 | 17325.000 | json-canon-rs | 4.45x |
| verify/canonical/nested-mixed | 3107.200 | 1641.000 | json-canon-rs | 1.89x |
| verify/canonical/numeric-boundary | 18861.100 | 8740.100 | json-canon-rs | 2.16x |
| verify/canonical/rfc-key-sorting | 2843.200 | 1321.200 | json-canon-rs | 2.15x |
| verify/canonical/small | 1655.800 | 1013.000 | json-canon-rs | 1.63x |
| verify/canonical/surrogate-pair | 652.600 | 297.000 | json-canon-rs | 2.20x |
| verify/canonical/unicode | 1546.800 | 1631.400 | json-canon | 1.05x |
| verify/canonical/verify-whitespace | 1588.400 | 978.200 | json-canon-rs | 1.62x |

## Statistical Inference

| track | mode | workload | winner | speedup | ci95 | p-value | p-adj | sig-BH |
|---|---|---|---|---:|---|---:|---:|---|
| e2e | canonicalize | array-2048 | schubfach | 1.292x | [1.210, 1.384] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | array-2048 | schubfach-rs | 1.759x | [1.578, 1.992] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | array-2048 | schubfach-rs | 2.098x | [1.910, 2.329] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | array-2048 | json-canon-rs | 1.541x | [1.394, 1.667] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | array-256 | schubfach | 1.127x | [0.991, 1.282] | 0.0912 | 0.1072 | false |
| e2e | canonicalize | array-256 | schubfach-rs | 1.391x | [1.262, 1.515] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | array-256 | schubfach-rs | 2.322x | [2.079, 2.589] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | array-256 | json-canon-rs | 1.882x | [1.693, 2.122] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | canonical-minimal | schubfach | 1.036x | [0.917, 1.163] | 0.5691 | 0.6046 | false |
| e2e | canonicalize | canonical-minimal | schubfach-rs | 1.067x | [0.987, 1.159] | 0.1492 | 0.1698 | false |
| e2e | canonicalize | canonical-minimal | schubfach-rs | 2.767x | [2.517, 3.003] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | canonical-minimal | json-canon-rs | 2.688x | [2.394, 2.998] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | control-escapes | schubfach | 1.213x | [1.083, 1.369] | 0.0046 | 0.0059 | true |
| e2e | canonicalize | control-escapes | schubfach-rs | 1.117x | [1.030, 1.199] | 0.0110 | 0.0139 | true |
| e2e | canonicalize | control-escapes | schubfach-rs | 2.597x | [2.316, 2.883] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | control-escapes | json-canon-rs | 2.819x | [2.598, 3.076] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | deep-64 | schubfach | 1.117x | [0.997, 1.259] | 0.0864 | 0.1021 | false |
| e2e | canonicalize | deep-64 | schubfach-rs | 1.147x | [1.067, 1.258] | 0.0012 | 0.0016 | true |
| e2e | canonicalize | deep-64 | schubfach-rs | 2.600x | [2.339, 2.906] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | deep-64 | json-canon-rs | 2.532x | [2.291, 2.759] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | escaped-key-order | json-canon | 1.011x | [0.899, 1.138] | 0.8612 | 0.8863 | false |
| e2e | canonicalize | escaped-key-order | schubfach-rs | 1.242x | [1.178, 1.302] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | escaped-key-order | schubfach-rs | 2.702x | [2.423, 2.893] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | escaped-key-order | json-canon-rs | 2.150x | [1.939, 2.354] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | long-string | schubfach | 1.201x | [1.031, 1.394] | 0.0292 | 0.0359 | true |
| e2e | canonicalize | long-string | schubfach-rs | 1.051x | [0.958, 1.158] | 0.3217 | 0.3516 | false |
| e2e | canonicalize | long-string | schubfach-rs | 2.328x | [2.067, 2.658] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | long-string | json-canon-rs | 2.660x | [2.333, 3.005] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | nested-mixed | json-canon | 1.010x | [0.894, 1.147] | 0.8732 | 0.8943 | false |
| e2e | canonicalize | nested-mixed | schubfach-rs | 1.161x | [1.057, 1.253] | 0.0034 | 0.0044 | true |
| e2e | canonicalize | nested-mixed | schubfach-rs | 2.724x | [2.509, 2.934] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | nested-mixed | json-canon-rs | 2.324x | [2.028, 2.634] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | numeric-boundary | schubfach | 1.110x | [0.985, 1.236] | 0.0934 | 0.1093 | false |
| e2e | canonicalize | numeric-boundary | schubfach-rs | 1.358x | [1.295, 1.427] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | numeric-boundary | schubfach-rs | 2.828x | [2.599, 3.044] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | numeric-boundary | json-canon-rs | 2.311x | [2.071, 2.510] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | rfc-key-sorting | schubfach | 1.030x | [0.924, 1.142] | 0.5871 | 0.6197 | false |
| e2e | canonicalize | rfc-key-sorting | schubfach-rs | 1.098x | [1.021, 1.180] | 0.0214 | 0.0266 | true |
| e2e | canonicalize | rfc-key-sorting | schubfach-rs | 2.916x | [2.666, 3.134] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | rfc-key-sorting | json-canon-rs | 2.735x | [2.463, 2.976] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | small | json-canon | 1.001x | [0.877, 1.129] | 0.9942 | 0.9942 | false |
| e2e | canonicalize | small | schubfach-rs | 1.210x | [1.098, 1.326] | 0.0014 | 0.0018 | true |
| e2e | canonicalize | small | schubfach-rs | 2.733x | [2.449, 3.015] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | small | json-canon-rs | 2.258x | [2.003, 2.557] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | surrogate-pair | schubfach | 1.009x | [0.904, 1.133] | 0.8762 | 0.8943 | false |
| e2e | canonicalize | surrogate-pair | schubfach-rs | 1.125x | [1.034, 1.233] | 0.0210 | 0.0262 | true |
| e2e | canonicalize | surrogate-pair | schubfach-rs | 2.819x | [2.549, 3.094] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | surrogate-pair | json-canon-rs | 2.529x | [2.261, 2.816] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | unicode | json-canon | 1.087x | [0.964, 1.213] | 0.1568 | 0.1776 | false |
| e2e | canonicalize | unicode | schubfach-rs | 1.145x | [1.059, 1.245] | 0.0034 | 0.0044 | true |
| e2e | canonicalize | unicode | schubfach-rs | 2.989x | [2.743, 3.255] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | unicode | json-canon-rs | 2.402x | [2.166, 2.692] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | verify-whitespace | schubfach | 1.109x | [0.961, 1.267] | 0.1672 | 0.1885 | false |
| e2e | canonicalize | verify-whitespace | schubfach-rs | 1.237x | [1.132, 1.334] | 0.0004 | 0.0005 | true |
| e2e | canonicalize | verify-whitespace | schubfach-rs | 2.605x | [2.345, 2.869] | 0.0002 | 0.0003 | true |
| e2e | canonicalize | verify-whitespace | json-canon-rs | 2.336x | [2.051, 2.662] | 0.0002 | 0.0003 | true |
| e2e | verify | array-2048 | schubfach | 1.268x | [1.205, 1.348] | 0.0002 | 0.0003 | true |
| e2e | verify | array-2048 | schubfach-rs | 1.714x | [1.549, 1.917] | 0.0002 | 0.0003 | true |
| e2e | verify | array-2048 | schubfach-rs | 2.325x | [2.096, 2.544] | 0.0002 | 0.0003 | true |
| e2e | verify | array-2048 | json-canon-rs | 1.719x | [1.563, 1.826] | 0.0002 | 0.0003 | true |
| e2e | verify | array-256 | schubfach | 1.284x | [1.130, 1.454] | 0.0018 | 0.0024 | true |
| e2e | verify | array-256 | schubfach-rs | 1.482x | [1.359, 1.622] | 0.0002 | 0.0003 | true |
| e2e | verify | array-256 | schubfach-rs | 2.232x | [2.011, 2.503] | 0.0002 | 0.0003 | true |
| e2e | verify | array-256 | json-canon-rs | 1.935x | [1.739, 2.166] | 0.0002 | 0.0003 | true |
| e2e | verify | canonical-minimal | schubfach | 1.002x | [0.882, 1.141] | 0.9748 | 0.9868 | false |
| e2e | verify | canonical-minimal | schubfach-rs | 1.043x | [0.966, 1.116] | 0.2881 | 0.3163 | false |
| e2e | verify | canonical-minimal | schubfach-rs | 2.488x | [2.269, 2.732] | 0.0002 | 0.0003 | true |
| e2e | verify | canonical-minimal | json-canon-rs | 2.390x | [2.141, 2.694] | 0.0002 | 0.0003 | true |
| e2e | verify | control-escapes | schubfach | 1.136x | [1.027, 1.259] | 0.0310 | 0.0379 | true |
| e2e | verify | control-escapes | schubfach-rs | 1.148x | [1.087, 1.235] | 0.0004 | 0.0005 | true |
| e2e | verify | control-escapes | schubfach-rs | 2.601x | [2.386, 2.843] | 0.0002 | 0.0003 | true |
| e2e | verify | control-escapes | json-canon-rs | 2.574x | [2.353, 2.780] | 0.0002 | 0.0003 | true |
| e2e | verify | deep-64 | schubfach | 1.211x | [1.075, 1.379] | 0.0074 | 0.0094 | true |
| e2e | verify | deep-64 | schubfach-rs | 1.085x | [0.983, 1.198] | 0.1254 | 0.1447 | false |
| e2e | verify | deep-64 | schubfach-rs | 2.644x | [2.307, 3.012] | 0.0002 | 0.0003 | true |
| e2e | verify | deep-64 | json-canon-rs | 2.952x | [2.715, 3.256] | 0.0002 | 0.0003 | true |
| e2e | verify | escaped-key-order | schubfach | 1.080x | [0.955, 1.229] | 0.2691 | 0.2995 | false |
| e2e | verify | escaped-key-order | schubfach-rs | 1.168x | [1.090, 1.250] | 0.0006 | 0.0008 | true |
| e2e | verify | escaped-key-order | schubfach-rs | 2.599x | [2.323, 2.877] | 0.0002 | 0.0003 | true |
| e2e | verify | escaped-key-order | json-canon-rs | 2.404x | [2.184, 2.652] | 0.0002 | 0.0003 | true |
| e2e | verify | long-string | schubfach | 1.035x | [0.895, 1.196] | 0.6385 | 0.6711 | false |
| e2e | verify | long-string | schubfach-rs | 1.057x | [0.987, 1.126] | 0.1222 | 0.1423 | false |
| e2e | verify | long-string | schubfach-rs | 2.709x | [2.407, 2.986] | 0.0002 | 0.0003 | true |
| e2e | verify | long-string | json-canon-rs | 2.655x | [2.361, 2.960] | 0.0002 | 0.0003 | true |
| e2e | verify | nested-mixed | schubfach | 1.153x | [1.016, 1.302] | 0.0314 | 0.0382 | true |
| e2e | verify | nested-mixed | schubfach-rs | 1.170x | [1.060, 1.274] | 0.0054 | 0.0069 | true |
| e2e | verify | nested-mixed | schubfach-rs | 2.659x | [2.400, 2.916] | 0.0002 | 0.0003 | true |
| e2e | verify | nested-mixed | json-canon-rs | 2.621x | [2.335, 2.951] | 0.0002 | 0.0003 | true |
| e2e | verify | numeric-boundary | schubfach | 1.116x | [0.988, 1.280] | 0.1252 | 0.1447 | false |
| e2e | verify | numeric-boundary | schubfach-rs | 1.313x | [1.225, 1.383] | 0.0002 | 0.0003 | true |
| e2e | verify | numeric-boundary | schubfach-rs | 2.628x | [2.326, 2.894] | 0.0002 | 0.0003 | true |
| e2e | verify | numeric-boundary | json-canon-rs | 2.235x | [2.028, 2.460] | 0.0002 | 0.0003 | true |
| e2e | verify | rfc-key-sorting | schubfach | 1.137x | [1.016, 1.267] | 0.0364 | 0.0441 | true |
| e2e | verify | rfc-key-sorting | schubfach-rs | 1.124x | [1.023, 1.215] | 0.0206 | 0.0258 | true |
| e2e | verify | rfc-key-sorting | schubfach-rs | 2.794x | [2.504, 3.060] | 0.0002 | 0.0003 | true |
| e2e | verify | rfc-key-sorting | json-canon-rs | 2.826x | [2.548, 3.139] | 0.0002 | 0.0003 | true |
| e2e | verify | small | schubfach | 1.214x | [1.080, 1.357] | 0.0020 | 0.0026 | true |
| e2e | verify | small | schubfach-rs | 1.296x | [1.220, 1.385] | 0.0002 | 0.0003 | true |
| e2e | verify | small | schubfach-rs | 2.565x | [2.289, 2.846] | 0.0002 | 0.0003 | true |
| e2e | verify | small | json-canon-rs | 2.403x | [2.202, 2.577] | 0.0002 | 0.0003 | true |
| e2e | verify | surrogate-pair | schubfach | 1.122x | [0.993, 1.266] | 0.0852 | 0.1012 | false |
| e2e | verify | surrogate-pair | schubfach-rs | 1.084x | [0.977, 1.194] | 0.1422 | 0.1633 | false |
| e2e | verify | surrogate-pair | schubfach-rs | 2.716x | [2.422, 3.001] | 0.0002 | 0.0003 | true |
| e2e | verify | surrogate-pair | json-canon-rs | 2.814x | [2.481, 3.161] | 0.0002 | 0.0003 | true |
| e2e | verify | unicode | schubfach | 1.188x | [1.022, 1.368] | 0.0286 | 0.0353 | true |
| e2e | verify | unicode | schubfach-rs | 1.151x | [1.084, 1.248] | 0.0006 | 0.0008 | true |
| e2e | verify | unicode | schubfach-rs | 2.564x | [2.257, 2.894] | 0.0002 | 0.0003 | true |
| e2e | verify | unicode | json-canon-rs | 2.645x | [2.376, 2.880] | 0.0002 | 0.0003 | true |
| e2e | verify | verify-whitespace | schubfach | 1.084x | [0.946, 1.263] | 0.2821 | 0.3125 | false |
| e2e | verify | verify-whitespace | schubfach-rs | 1.243x | [1.150, 1.329] | 0.0002 | 0.0003 | true |
| e2e | verify | verify-whitespace | schubfach-rs | 2.616x | [2.291, 2.938] | 0.0002 | 0.0003 | true |
| e2e | verify | verify-whitespace | json-canon-rs | 2.281x | [2.062, 2.524] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-2048 | schubfach | 1.604x | [1.596, 1.612] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-2048 | schubfach-rs | 2.431x | [2.401, 2.458] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-2048 | schubfach-rs | 2.144x | [2.122, 2.164] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-2048 | json-canon-rs | 1.415x | [1.405, 1.427] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-256 | schubfach | 1.550x | [1.540, 1.563] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-256 | schubfach-rs | 2.047x | [2.005, 2.085] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-256 | schubfach-rs | 1.945x | [1.915, 1.968] | 0.0002 | 0.0003 | true |
| api | canonicalize | array-256 | json-canon-rs | 1.473x | [1.449, 1.496] | 0.0002 | 0.0003 | true |
| api | canonicalize | canonical-minimal | schubfach | 1.514x | [1.506, 1.522] | 0.0002 | 0.0003 | true |
| api | canonicalize | canonical-minimal | schubfach-rs | 2.176x | [1.981, 2.314] | 0.0002 | 0.0003 | true |
| api | canonicalize | canonical-minimal | schubfach-rs | 1.842x | [1.682, 1.953] | 0.0002 | 0.0003 | true |
| api | canonicalize | canonical-minimal | json-canon-rs | 1.281x | [1.248, 1.309] | 0.0002 | 0.0003 | true |
| api | canonicalize | control-escapes | json-canon | 1.065x | [1.059, 1.074] | 0.0002 | 0.0003 | true |
| api | canonicalize | control-escapes | schubfach-rs | 1.000x | [0.968, 1.035] | 0.9892 | 0.9942 | false |
| api | canonicalize | control-escapes | schubfach-rs | 2.206x | [2.166, 2.254] | 0.0002 | 0.0003 | true |
| api | canonicalize | control-escapes | json-canon-rs | 2.072x | [2.016, 2.131] | 0.0002 | 0.0003 | true |
| api | canonicalize | deep-64 | schubfach | 1.000x | [0.995, 1.005] | 0.9936 | 0.9942 | false |
| api | canonicalize | deep-64 | json-canon-rs | 1.021x | [0.987, 1.135] | 0.9502 | 0.9659 | false |
| api | canonicalize | deep-64 | schubfach-rs | 1.730x | [1.574, 1.788] | 0.0002 | 0.0003 | true |
| api | canonicalize | deep-64 | json-canon-rs | 1.767x | [1.753, 1.787] | 0.0002 | 0.0003 | true |
| api | canonicalize | deep | schubfach | 1.001x | [0.994, 1.007] | 0.8556 | 0.8843 | false |
| api | canonicalize | escaped-key-order | schubfach | 1.597x | [1.590, 1.605] | 0.0002 | 0.0003 | true |
| api | canonicalize | escaped-key-order | schubfach-rs | 2.503x | [2.429, 2.575] | 0.0002 | 0.0003 | true |
| api | canonicalize | escaped-key-order | schubfach-rs | 2.371x | [2.324, 2.424] | 0.0002 | 0.0003 | true |
| api | canonicalize | escaped-key-order | json-canon-rs | 1.513x | [1.486, 1.548] | 0.0002 | 0.0003 | true |
| api | canonicalize | large | schubfach | 1.725x | [1.717, 1.733] | 0.0002 | 0.0003 | true |
| api | canonicalize | long-string | json-canon | 1.003x | [0.992, 1.016] | 0.6959 | 0.7252 | false |
| api | canonicalize | long-string | schubfach-rs | 1.001x | [0.999, 1.003] | 0.2845 | 0.3138 | false |
| api | canonicalize | long-string | schubfach-rs | 4.394x | [4.364, 4.429] | 0.0002 | 0.0003 | true |
| api | canonicalize | long-string | json-canon-rs | 4.378x | [4.329, 4.414] | 0.0002 | 0.0003 | true |
| api | canonicalize | medium | schubfach | 1.709x | [1.700, 1.720] | 0.0002 | 0.0003 | true |
| api | canonicalize | mixed-prod | schubfach | 1.324x | [1.313, 1.335] | 0.0002 | 0.0003 | true |
| api | canonicalize | nested-mixed | schubfach | 1.414x | [1.406, 1.423] | 0.0002 | 0.0003 | true |
| api | canonicalize | nested-mixed | schubfach-rs | 1.998x | [1.918, 2.091] | 0.0002 | 0.0003 | true |
| api | canonicalize | nested-mixed | schubfach-rs | 2.592x | [2.507, 2.690] | 0.0002 | 0.0003 | true |
| api | canonicalize | nested-mixed | json-canon-rs | 1.834x | [1.780, 1.867] | 0.0002 | 0.0003 | true |
| api | canonicalize | number-heavy | schubfach | 1.551x | [1.539, 1.564] | 0.0002 | 0.0003 | true |
| api | canonicalize | numeric-boundary | schubfach | 1.482x | [1.470, 1.496] | 0.0002 | 0.0003 | true |
| api | canonicalize | numeric-boundary | schubfach-rs | 9.191x | [8.815, 9.382] | 0.0002 | 0.0003 | true |
| api | canonicalize | numeric-boundary | schubfach-rs | 12.998x | [12.499, 13.235] | 0.0002 | 0.0003 | true |
| api | canonicalize | numeric-boundary | json-canon-rs | 2.096x | [2.067, 2.127] | 0.0002 | 0.0003 | true |
| api | canonicalize | rfc-key-sorting | json-canon | 1.010x | [1.006, 1.014] | 0.0002 | 0.0003 | true |
| api | canonicalize | rfc-key-sorting | json-canon-rs | 1.035x | [1.001, 1.072] | 0.0768 | 0.0916 | false |
| api | canonicalize | rfc-key-sorting | schubfach-rs | 2.320x | [2.290, 2.353] | 0.0002 | 0.0003 | true |
| api | canonicalize | rfc-key-sorting | json-canon-rs | 2.378x | [2.303, 2.455] | 0.0002 | 0.0003 | true |
| api | canonicalize | small | schubfach | 1.338x | [1.315, 1.351] | 0.0002 | 0.0003 | true |
| api | canonicalize | small | schubfach-rs | 1.715x | [1.655, 1.780] | 0.0002 | 0.0003 | true |
| api | canonicalize | small | schubfach-rs | 2.146x | [2.080, 2.222] | 0.0002 | 0.0003 | true |
| api | canonicalize | small | json-canon-rs | 1.675x | [1.642, 1.713] | 0.0002 | 0.0003 | true |
| api | canonicalize | surrogate-pair | schubfach | 1.015x | [0.997, 1.038] | 0.1984 | 0.2217 | false |
| api | canonicalize | surrogate-pair | schubfach-rs | 1.051x | [0.993, 1.119] | 0.1452 | 0.1660 | false |
| api | canonicalize | surrogate-pair | schubfach-rs | 2.424x | [2.315, 2.496] | 0.0002 | 0.0003 | true |
| api | canonicalize | surrogate-pair | json-canon-rs | 2.341x | [2.204, 2.450] | 0.0002 | 0.0003 | true |
| api | canonicalize | unicode | json-canon | 1.002x | [0.998, 1.006] | 0.3375 | 0.3673 | false |
| api | canonicalize | unicode | schubfach-rs | 1.172x | [1.049, 1.358] | 0.0416 | 0.0501 | false |
| api | canonicalize | unicode | schubfach | 1.079x | [1.065, 1.103] | 0.0002 | 0.0003 | true |
| api | canonicalize | unicode | json-canon | 1.267x | [1.136, 1.457] | 0.0002 | 0.0003 | true |
| api | canonicalize | verify-whitespace | schubfach | 1.724x | [1.715, 1.730] | 0.0002 | 0.0003 | true |
| api | canonicalize | verify-whitespace | schubfach-rs | 2.826x | [2.721, 2.933] | 0.0002 | 0.0003 | true |
| api | canonicalize | verify-whitespace | schubfach-rs | 2.623x | [2.538, 2.691] | 0.0002 | 0.0003 | true |
| api | canonicalize | verify-whitespace | json-canon-rs | 1.600x | [1.553, 1.631] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-2048 | schubfach | 1.624x | [1.617, 1.631] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-2048 | schubfach-rs | 2.406x | [2.376, 2.451] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-2048 | schubfach-rs | 2.053x | [2.030, 2.075] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-2048 | json-canon-rs | 1.386x | [1.359, 1.397] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-256 | schubfach | 1.534x | [1.526, 1.541] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-256 | schubfach-rs | 2.188x | [1.947, 2.320] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-256 | schubfach-rs | 2.236x | [1.991, 2.367] | 0.0002 | 0.0003 | true |
| api | verify | canonical/array-256 | json-canon-rs | 1.567x | [1.556, 1.576] | 0.0002 | 0.0003 | true |
| api | verify | canonical/canonical-minimal | schubfach | 1.523x | [1.514, 1.533] | 0.0002 | 0.0003 | true |
| api | verify | canonical/canonical-minimal | schubfach-rs | 2.363x | [2.236, 2.476] | 0.0002 | 0.0003 | true |
| api | verify | canonical/canonical-minimal | schubfach-rs | 2.055x | [1.947, 2.123] | 0.0002 | 0.0003 | true |
| api | verify | canonical/canonical-minimal | json-canon-rs | 1.325x | [1.278, 1.355] | 0.0002 | 0.0003 | true |
| api | verify | canonical/control-escapes | json-canon | 1.044x | [1.040, 1.049] | 0.0002 | 0.0003 | true |
| api | verify | canonical/control-escapes | json-canon-rs | 1.010x | [0.983, 1.041] | 0.5247 | 0.5635 | false |
| api | verify | canonical/control-escapes | schubfach-rs | 2.167x | [2.120, 2.206] | 0.0002 | 0.0003 | true |
| api | verify | canonical/control-escapes | json-canon-rs | 2.096x | [2.054, 2.143] | 0.0002 | 0.0003 | true |
| api | verify | canonical/deep-64 | schubfach | 1.003x | [0.999, 1.007] | 0.1822 | 0.2045 | false |
| api | verify | canonical/deep-64 | json-canon-rs | 1.005x | [0.997, 1.019] | 0.3689 | 0.3979 | false |
| api | verify | canonical/deep-64 | schubfach-rs | 1.833x | [1.813, 1.847] | 0.0004 | 0.0005 | true |
| api | verify | canonical/deep-64 | json-canon-rs | 1.849x | [1.839, 1.865] | 0.0002 | 0.0003 | true |
| api | verify | canonical/deep | schubfach | 1.011x | [1.007, 1.018] | 0.0004 | 0.0005 | true |
| api | verify | canonical/escaped-key-order | schubfach | 1.639x | [1.630, 1.648] | 0.0002 | 0.0003 | true |
| api | verify | canonical/escaped-key-order | schubfach-rs | 2.534x | [2.429, 2.666] | 0.0002 | 0.0003 | true |
| api | verify | canonical/escaped-key-order | schubfach-rs | 2.412x | [2.315, 2.472] | 0.0002 | 0.0003 | true |
| api | verify | canonical/escaped-key-order | json-canon-rs | 1.560x | [1.480, 1.596] | 0.0002 | 0.0003 | true |
| api | verify | canonical/large | schubfach | 1.735x | [1.727, 1.743] | 0.0002 | 0.0003 | true |
| api | verify | canonical/long-string | json-canon | 1.003x | [0.993, 1.012] | 0.5451 | 0.5828 | false |
| api | verify | canonical/long-string | json-canon-rs | 1.003x | [0.993, 1.016] | 0.5703 | 0.6046 | false |
| api | verify | canonical/long-string | schubfach-rs | 4.445x | [4.407, 4.498] | 0.0004 | 0.0005 | true |
| api | verify | canonical/long-string | json-canon-rs | 4.446x | [4.408, 4.505] | 0.0002 | 0.0003 | true |
| api | verify | canonical/medium | schubfach | 1.690x | [1.679, 1.700] | 0.0002 | 0.0003 | true |
| api | verify | canonical/mixed-prod | schubfach | 1.320x | [1.314, 1.325] | 0.0002 | 0.0003 | true |
| api | verify | canonical/nested-mixed | schubfach | 1.400x | [1.394, 1.406] | 0.0002 | 0.0003 | true |
| api | verify | canonical/nested-mixed | schubfach-rs | 1.935x | [1.873, 1.999] | 0.0002 | 0.0003 | true |
| api | verify | canonical/nested-mixed | schubfach-rs | 2.616x | [2.551, 2.691] | 0.0002 | 0.0003 | true |
| api | verify | canonical/nested-mixed | json-canon-rs | 1.893x | [1.848, 1.927] | 0.0002 | 0.0003 | true |
| api | verify | canonical/number-heavy | schubfach | 1.552x | [1.541, 1.562] | 0.0002 | 0.0003 | true |
| api | verify | canonical/numeric-boundary | schubfach | 1.474x | [1.463, 1.486] | 0.0002 | 0.0003 | true |
| api | verify | canonical/numeric-boundary | schubfach-rs | 9.091x | [8.972, 9.208] | 0.0002 | 0.0003 | true |
| api | verify | canonical/numeric-boundary | schubfach-rs | 13.307x | [13.134, 13.465] | 0.0002 | 0.0003 | true |
| api | verify | canonical/numeric-boundary | json-canon-rs | 2.158x | [2.141, 2.177] | 0.0002 | 0.0003 | true |
| api | verify | canonical/rfc-key-sorting | json-canon | 1.005x | [1.001, 1.009] | 0.0462 | 0.0554 | false |
| api | verify | canonical/rfc-key-sorting | schubfach-rs | 1.007x | [0.963, 1.040] | 0.7556 | 0.7842 | false |
| api | verify | canonical/rfc-key-sorting | schubfach-rs | 2.177x | [2.111, 2.225] | 0.0002 | 0.0003 | true |
| api | verify | canonical/rfc-key-sorting | json-canon-rs | 2.152x | [2.101, 2.218] | 0.0002 | 0.0003 | true |
| api | verify | canonical/small | schubfach | 1.353x | [1.348, 1.360] | 0.0002 | 0.0003 | true |
| api | verify | canonical/small | schubfach-rs | 1.683x | [1.604, 1.760] | 0.0002 | 0.0003 | true |
| api | verify | canonical/small | schubfach-rs | 2.032x | [1.945, 2.103] | 0.0002 | 0.0003 | true |
| api | verify | canonical/small | json-canon-rs | 1.635x | [1.572, 1.667] | 0.0002 | 0.0003 | true |
| api | verify | canonical/surrogate-pair | json-canon | 1.008x | [1.005, 1.012] | 0.0006 | 0.0008 | true |
| api | verify | canonical/surrogate-pair | schubfach-rs | 1.051x | [0.976, 1.166] | 0.3519 | 0.3813 | false |
| api | verify | canonical/surrogate-pair | schubfach-rs | 2.325x | [2.234, 2.414] | 0.0002 | 0.0003 | true |
| api | verify | canonical/surrogate-pair | json-canon-rs | 2.195x | [1.991, 2.344] | 0.0002 | 0.0003 | true |
| api | verify | canonical/unicode | json-canon | 1.001x | [0.997, 1.006] | 0.6795 | 0.7111 | false |
| api | verify | canonical/unicode | json-canon-rs | 1.040x | [1.024, 1.059] | 0.0008 | 0.0011 | true |
| api | verify | canonical/unicode | schubfach | 1.096x | [1.083, 1.115] | 0.0004 | 0.0005 | true |
| api | verify | canonical/unicode | json-canon | 1.055x | [1.045, 1.066] | 0.0002 | 0.0003 | true |
| api | verify | canonical/verify-whitespace | schubfach | 1.702x | [1.695, 1.709] | 0.0004 | 0.0005 | true |
| api | verify | canonical/verify-whitespace | schubfach-rs | 2.691x | [2.510, 2.823] | 0.0002 | 0.0003 | true |
| api | verify | canonical/verify-whitespace | schubfach-rs | 2.568x | [2.400, 2.689] | 0.0002 | 0.0003 | true |
| api | verify | canonical/verify-whitespace | json-canon-rs | 1.624x | [1.595, 1.652] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/control-escapes | json-canon | 1.048x | [1.042, 1.054] | 0.0004 | 0.0005 | true |
| api | verify | noncanonical/escaped-key-order | schubfach | 1.604x | [1.596, 1.613] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/large | schubfach | 1.719x | [1.711, 1.730] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/medium | schubfach | 1.671x | [1.664, 1.680] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/mixed-prod | schubfach | 1.299x | [1.293, 1.306] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/nested-mixed | schubfach | 1.394x | [1.388, 1.401] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/number-heavy | schubfach | 1.538x | [1.526, 1.552] | 0.0004 | 0.0005 | true |
| api | verify | noncanonical/numeric-boundary | schubfach | 1.469x | [1.461, 1.480] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/rfc-key-sorting | json-canon | 1.022x | [1.016, 1.027] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/small | schubfach | 1.339x | [1.332, 1.345] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/surrogate-pair | json-canon | 1.019x | [1.014, 1.023] | 0.0002 | 0.0003 | true |
| api | verify | noncanonical/unicode | json-canon | 1.009x | [1.002, 1.015] | 0.0160 | 0.0202 | true |
| api | verify | noncanonical/verify-whitespace | schubfach | 1.702x | [1.694, 1.710] | 0.0002 | 0.0003 | true |

## benchstat Snippet

```text
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
```

## Production Recommendation

Recommend `schubfach-rs` based on BH-significant wins with conformance/oracle gates passing.
