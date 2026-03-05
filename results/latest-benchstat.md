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

## benchstat Output

```text
│ /tmp/jcs-benchstat-3441469166/schubfach.txt │ /tmp/jcs-benchstat-3441469166/json-canon.txt │
                                      │                   sec/op                    │       sec/op         vs base                 │
Canonicalize/array-2048                                                10.39m ± ∞ ¹          13.40m ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/array-256                                                 1.242m ± ∞ ¹          1.484m ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/canonical-minimal                                         1.034µ ± ∞ ¹          1.147µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/control-escapes                                           4.396µ ± ∞ ¹          4.419µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/deep                                                      135.0µ ± ∞ ¹          141.4µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/deep-64                                                   55.53µ ± ∞ ¹          64.08µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/escaped-key-order                                         2.706µ ± ∞ ¹          3.069µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/large                                                     34.42m ± ∞ ¹          32.71m ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/long-string                                               109.4µ ± ∞ ¹          104.9µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/medium                                                    1.235m ± ∞ ¹          1.611m ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/mixed-prod                                                478.7µ ± ∞ ¹          554.7µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/nested-mixed                                              6.505µ ± ∞ ¹          9.405µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/number-heavy                                              17.50µ ± ∞ ¹          24.91µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/numeric-boundary                                          14.22µ ± ∞ ¹          19.65µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/rfc-key-sorting                                           9.266µ ± ∞ ¹          8.312µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/small                                                     4.030µ ± ∞ ¹          5.605µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/surrogate-pair                                            1.976µ ± ∞ ¹          2.052µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/unicode                                                   4.731µ ± ∞ ¹          5.217µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/verify-whitespace                                         3.699µ ± ∞ ¹          4.280µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/array-2048                                            11.20m ± ∞ ¹          13.11m ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/array-256                                             1.227m ± ∞ ¹          1.704m ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/canonical-minimal                                     1.379µ ± ∞ ¹          1.624µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/control-escapes                                       3.780µ ± ∞ ¹          4.964µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/deep                                                  112.6µ ± ∞ ¹          142.4µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/deep-64                                               63.23µ ± ∞ ¹          65.23µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/escaped-key-order                                     2.413µ ± ∞ ¹          3.476µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/large                                                 33.52m ± ∞ ¹          32.70m ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/long-string                                           143.2µ ± ∞ ¹          127.7µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/medium                                                1.500m ± ∞ ¹          1.923m ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/mixed-prod                                            545.6µ ± ∞ ¹          666.3µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/nested-mixed                                          8.609µ ± ∞ ¹          9.745µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/number-heavy                                          20.35µ ± ∞ ¹          25.87µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/numeric-boundary                                      15.16µ ± ∞ ¹          21.14µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/rfc-key-sorting                                      10.990µ ± ∞ ¹          9.040µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/small                                                 4.072µ ± ∞ ¹          6.616µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/surrogate-pair                                        2.368µ ± ∞ ¹          2.875µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/unicode                                               4.569µ ± ∞ ¹          5.358µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/verify-whitespace                                     2.761µ ± ∞ ¹          4.397µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/control-escapes                                    4.941µ ± ∞ ¹          5.013µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/escaped-key-order                                  2.084µ ± ∞ ¹          2.792µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/large                                              28.96m ± ∞ ¹          32.46m ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/medium                                             754.2µ ± ∞ ¹         1650.6µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/mixed-prod                                         474.1µ ± ∞ ¹          671.6µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/nested-mixed                                       6.415µ ± ∞ ¹          9.879µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/number-heavy                                       16.51µ ± ∞ ¹          23.92µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/numeric-boundary                                   14.55µ ± ∞ ¹          19.89µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/rfc-key-sorting                                    8.347µ ± ∞ ¹          7.785µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/small                                              4.756µ ± ∞ ¹          4.251µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/surrogate-pair                                     2.116µ ± ∞ ¹          1.955µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/unicode                                            4.861µ ± ∞ ¹          5.250µ ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/verify-whitespace                                  2.710µ ± ∞ ¹          3.415µ ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                                                                35.78µ                42.78µ        +19.55%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                                      │ /tmp/jcs-benchstat-3441469166/schubfach.txt │ /tmp/jcs-benchstat-3441469166/json-canon.txt │
                                      │                     B/s                     │         B/s          vs base                 │
Canonicalize/array-2048                                               9.079Mi ± ∞ ¹         7.038Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/array-256                                                9.146Mi ± ∞ ¹         7.658Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/canonical-minimal                                        6.456Mi ± ∞ ¹         5.817Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/control-escapes                                          12.80Mi ± ∞ ¹         12.73Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/deep                                                     8.183Mi ± ∞ ¹         7.811Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/deep-64                                                  8.755Mi ± ∞ ¹         7.591Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/escaped-key-order                                        6.342Mi ± ∞ ¹         5.598Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/large                                                    12.85Mi ± ∞ ¹         13.51Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/long-string                                              142.9Mi ± ∞ ¹         149.2Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/medium                                                  10.605Mi ± ∞ ¹         8.135Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/mixed-prod                                              11.263Mi ± ∞ ¹         9.718Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/nested-mixed                                             7.181Mi ± ∞ ¹         4.969Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/number-heavy                                             7.629Mi ± ∞ ¹         5.360Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/numeric-boundary                                         6.571Mi ± ∞ ¹         4.759Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/rfc-key-sorting                                          20.79Mi ± ∞ ¹         23.17Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/small                                                   10.653Mi ± ∞ ¹         7.658Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/surrogate-pair                                           13.99Mi ± ∞ ¹         13.48Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/unicode                                                  19.76Mi ± ∞ ¹         17.91Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/verify-whitespace                                        6.189Mi ± ∞ ¹         5.350Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/array-2048                                           8.421Mi ± ∞ ¹         7.200Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/array-256                                            9.260Mi ± ∞ ¹         6.666Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/canonical-minimal                                    4.845Mi ± ∞ ¹         4.110Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/control-escapes                                      14.89Mi ± ∞ ¹         11.33Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/deep                                                 9.804Mi ± ∞ ¹         7.753Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/deep-64                                              7.696Mi ± ∞ ¹         7.458Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/escaped-key-order                                    5.140Mi ± ∞ ¹         3.567Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/large                                                13.16Mi ± ∞ ¹         13.49Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/long-string                                          109.3Mi ± ∞ ¹         122.5Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/medium                                               8.717Mi ± ∞ ¹         6.800Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/mixed-prod                                           9.832Mi ± ∞ ¹         8.049Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/nested-mixed                                         5.426Mi ± ∞ ¹         4.797Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/number-heavy                                         6.371Mi ± ∞ ¹         5.016Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/numeric-boundary                                     6.161Mi ± ∞ ¹         4.425Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/rfc-key-sorting                                      15.62Mi ± ∞ ¹         18.99Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/small                                                9.842Mi ± ∞ ¹         6.056Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/surrogate-pair                                       8.459Mi ± ∞ ¹         6.971Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/unicode                                              20.04Mi ± ∞ ¹         17.09Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/verify-whitespace                                    5.875Mi ± ∞ ¹         3.691Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/control-escapes                                   11.39Mi ± ∞ ¹         11.22Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/escaped-key-order                                 8.240Mi ± ∞ ¹         6.151Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/large                                             15.27Mi ± ∞ ¹         13.62Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/medium                                           17.376Mi ± ∞ ¹         7.935Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/mixed-prod                                       11.368Mi ± ∞ ¹         8.030Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/nested-mixed                                      7.286Mi ± ∞ ¹         4.730Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/number-heavy                                      8.087Mi ± ∞ ¹         5.579Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/numeric-boundary                                  6.418Mi ± ∞ ¹         4.702Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/rfc-key-sorting                                   23.08Mi ± ∞ ¹         24.75Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/small                                             9.022Mi ± ∞ ¹        10.099Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/surrogate-pair                                    13.07Mi ± ∞ ¹         14.15Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/unicode                                           19.23Mi ± ∞ ¹         17.81Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/verify-whitespace                                 8.450Mi ± ∞ ¹         6.704Mi ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                                                               10.74Mi               8.987Mi        -16.34%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

                                      │ /tmp/jcs-benchstat-3441469166/schubfach.txt │ /tmp/jcs-benchstat-3441469166/json-canon.txt │
                                      │                    B/op                     │         B/op          vs base                │
Canonicalize/array-2048                                               4.711Mi ± ∞ ¹          4.888Mi ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/array-256                                                590.7Ki ± ∞ ¹          615.1Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/canonical-minimal                                          504.0 ± ∞ ¹            504.0 ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/control-escapes                                          2.016Ki ± ∞ ¹          2.016Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/deep                                                     57.58Ki ± ∞ ¹          57.58Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/deep-64                                                  26.27Ki ± ∞ ¹          26.27Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/escaped-key-order                                        1.141Ki ± ∞ ¹          1.181Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/large                                                    19.95Mi ± ∞ ¹          20.58Mi ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/long-string                                              35.11Ki ± ∞ ¹          35.11Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/medium                                                   604.2Ki ± ∞ ¹          624.8Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/mixed-prod                                               251.0Ki ± ∞ ¹          257.0Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/nested-mixed                                             4.031Ki ± ∞ ¹          4.114Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/number-heavy                                             4.469Ki ± ∞ ¹          5.279Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/numeric-boundary                                         3.133Ki ± ∞ ¹          3.822Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/rfc-key-sorting                                          4.156Ki ± ∞ ¹          4.156Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/small                                                    2.102Ki ± ∞ ¹          2.150Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Canonicalize/surrogate-pair                                           1.141Ki ± ∞ ¹          1.141Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/unicode                                                  1.972Ki ± ∞ ¹          1.972Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Canonicalize/verify-whitespace                                        1.539Ki ± ∞ ¹          1.619Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/array-2048                                           4.710Mi ± ∞ ¹          4.889Mi ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/array-256                                            590.7Ki ± ∞ ¹          615.3Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/canonical-minimal                                      504.0 ± ∞ ¹            504.0 ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/control-escapes                                      2.016Ki ± ∞ ¹          2.016Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/deep                                                 57.58Ki ± ∞ ¹          57.58Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/deep-64                                              26.27Ki ± ∞ ¹          26.27Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/escaped-key-order                                    1.133Ki ± ∞ ¹          1.173Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/large                                                19.95Mi ± ∞ ¹          20.58Mi ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/long-string                                          35.11Ki ± ∞ ¹          35.11Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/medium                                               604.2Ki ± ∞ ¹          625.0Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/mixed-prod                                           251.0Ki ± ∞ ¹          257.1Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/nested-mixed                                         4.031Ki ± ∞ ¹          4.115Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/number-heavy                                         4.469Ki ± ∞ ¹          5.283Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/numeric-boundary                                     3.133Ki ± ∞ ¹          3.824Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/rfc-key-sorting                                      4.141Ki ± ∞ ¹          4.141Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/small                                                2.102Ki ± ∞ ¹          2.151Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/canonical/surrogate-pair                                       1.133Ki ± ∞ ¹          1.133Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/unicode                                              1.956Ki ± ∞ ¹          1.956Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/canonical/verify-whitespace                                    1.539Ki ± ∞ ¹          1.619Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/control-escapes                                   2.016Ki ± ∞ ¹          2.016Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/noncanonical/escaped-key-order                                 1.141Ki ± ∞ ¹          1.181Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/large                                             19.95Mi ± ∞ ¹          20.58Mi ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/medium                                            604.2Ki ± ∞ ¹          624.8Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/mixed-prod                                        251.0Ki ± ∞ ¹          257.0Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/nested-mixed                                      4.031Ki ± ∞ ¹          4.114Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/number-heavy                                      4.469Ki ± ∞ ¹          5.278Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/numeric-boundary                                  3.133Ki ± ∞ ¹          3.822Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/rfc-key-sorting                                   4.156Ki ± ∞ ¹          4.156Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/noncanonical/small                                             2.102Ki ± ∞ ¹          2.150Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
Verify/noncanonical/surrogate-pair                                    1.141Ki ± ∞ ¹          1.141Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/noncanonical/unicode                                           1.972Ki ± ∞ ¹          1.972Ki ± ∞ ¹       ~ (p=1.000 n=1) ³
Verify/noncanonical/verify-whitespace                                 1.539Ki ± ∞ ¹          1.618Ki ± ∞ ¹       ~ (p=1.000 n=1) ²
geomean                                                               15.70Ki                16.29Ki        +3.81%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
³ all samples are equal

                                      │ /tmp/jcs-benchstat-3441469166/schubfach.txt │ /tmp/jcs-benchstat-3441469166/json-canon.txt │
                                      │                  allocs/op                  │      allocs/op       vs base                 │
Canonicalize/array-2048                                                49.13k ± ∞ ¹          57.34k ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/array-256                                                 6.131k ± ∞ ¹          7.151k ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/canonical-minimal                                          8.000 ± ∞ ¹           8.000 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/control-escapes                                            21.00 ± ∞ ¹           21.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/deep                                                       843.0 ± ∞ ¹           843.0 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/deep-64                                                    387.0 ± ∞ ¹           387.0 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/escaped-key-order                                          14.00 ± ∞ ¹           16.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/large                                                     204.8k ± ∞ ¹          237.6k ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/long-string                                                14.00 ± ∞ ¹           14.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/medium                                                    6.380k ± ∞ ¹          7.394k ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/mixed-prod                                                2.340k ± ∞ ¹          2.595k ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/nested-mixed                                               45.00 ± ∞ ¹           49.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/number-heavy                                               60.00 ± ∞ ¹           80.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/numeric-boundary                                           38.00 ± ∞ ¹           52.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/rfc-key-sorting                                            47.00 ± ∞ ¹           47.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/small                                                      22.00 ± ∞ ¹           24.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Canonicalize/surrogate-pair                                             14.00 ± ∞ ¹           14.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/unicode                                                    19.00 ± ∞ ¹           19.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Canonicalize/verify-whitespace                                          19.00 ± ∞ ¹           23.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/array-2048                                            49.13k ± ∞ ¹          57.34k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/array-256                                             6.131k ± ∞ ¹          7.152k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/canonical-minimal                                      8.000 ± ∞ ¹           8.000 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/control-escapes                                        21.00 ± ∞ ¹           21.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/deep                                                   843.0 ± ∞ ¹           843.0 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/deep-64                                                387.0 ± ∞ ¹           387.0 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/escaped-key-order                                      14.00 ± ∞ ¹           16.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/large                                                 204.8k ± ∞ ¹          237.6k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/long-string                                            14.00 ± ∞ ¹           14.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/medium                                                6.380k ± ∞ ¹          7.395k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/mixed-prod                                            2.340k ± ∞ ¹          2.595k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/nested-mixed                                           45.00 ± ∞ ¹           49.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/number-heavy                                           60.00 ± ∞ ¹           80.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/numeric-boundary                                       38.00 ± ∞ ¹           52.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/rfc-key-sorting                                        42.00 ± ∞ ¹           42.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/small                                                  22.00 ± ∞ ¹           24.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/canonical/surrogate-pair                                         14.00 ± ∞ ¹           14.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/unicode                                                19.00 ± ∞ ¹           19.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/canonical/verify-whitespace                                      19.00 ± ∞ ¹           23.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/control-escapes                                     21.00 ± ∞ ¹           21.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/noncanonical/escaped-key-order                                   14.00 ± ∞ ¹           16.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/large                                              204.8k ± ∞ ¹          237.6k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/medium                                             6.380k ± ∞ ¹          7.394k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/mixed-prod                                         2.340k ± ∞ ¹          2.595k ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/nested-mixed                                        45.00 ± ∞ ¹           49.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/number-heavy                                        60.00 ± ∞ ¹           80.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/numeric-boundary                                    38.00 ± ∞ ¹           52.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/rfc-key-sorting                                     47.00 ± ∞ ¹           47.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/noncanonical/small                                               22.00 ± ∞ ¹           24.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
Verify/noncanonical/surrogate-pair                                      14.00 ± ∞ ¹           14.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/noncanonical/unicode                                             19.00 ± ∞ ¹           19.00 ± ∞ ¹        ~ (p=1.000 n=1) ³
Verify/noncanonical/verify-whitespace                                   19.00 ± ∞ ¹           23.00 ± ∞ ¹        ~ (p=1.000 n=1) ²
geomean                                                                 157.7                 174.3        +10.54%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
³ all samples are equal
```
