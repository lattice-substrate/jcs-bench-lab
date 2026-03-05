# Statistical Comparison

Generated at: 2026-03-05T06:10:00Z

Source: `/home/lenny/jcs-bench-lab/results/latest-cli-runs.csv`

Protocol: permutation test + bootstrap 95% CI, alpha=0.0500, resamples=3000

| track | mode | workload | winner | speedup | ci95 | p-value | effect d | significant | practical |
|---|---|---|---|---:|---|---:|---:|---|---|
| e2e | canonicalize | array-2048 | json-canon | 1.086x | [0.676, 1.927] | 0.7591 | 0.166 | false | true |
| e2e | canonicalize | array-256 | schubfach | 1.171x | [0.750, 1.858] | 0.5165 | -0.368 | false | true |
| e2e | canonicalize | canonical-minimal | schubfach | 1.326x | [1.124, 1.828] | 0.0003 | -1.283 | true | true |
| e2e | canonicalize | control-escapes | schubfach | 1.439x | [0.865, 2.809] | 0.2403 | -0.714 | false | true |
| e2e | canonicalize | deep-64 | schubfach | 1.293x | [1.045, 1.757] | 0.0497 | -1.096 | true | true |
| e2e | canonicalize | escaped-key-order | schubfach | 1.014x | [0.686, 1.472] | 0.9600 | -0.036 | false | false |
| e2e | canonicalize | long-string | schubfach | 1.077x | [0.735, 1.594] | 0.6811 | -0.201 | false | true |
| e2e | canonicalize | nested-mixed | json-canon | 1.344x | [0.853, 2.280] | 0.2329 | 0.663 | false | true |
| e2e | canonicalize | numeric-boundary | json-canon | 1.077x | [0.657, 1.817] | 0.8870 | 0.151 | false | true |
| e2e | canonicalize | rfc-key-sorting | json-canon | 1.387x | [1.053, 2.075] | 0.0610 | 1.138 | false | true |
| e2e | canonicalize | small | json-canon | 1.401x | [1.008, 2.201] | 0.0760 | 0.961 | false | true |
| e2e | canonicalize | surrogate-pair | schubfach | 1.223x | [0.850, 1.831] | 0.3242 | -0.563 | false | true |
| e2e | canonicalize | unicode | schubfach | 1.212x | [0.785, 2.059] | 0.4209 | -0.440 | false | true |
| e2e | canonicalize | verify-whitespace | schubfach | 1.002x | [0.615, 1.519] | 0.9960 | -0.004 | false | false |
| e2e | verify | array-2048 | schubfach | 1.217x | [0.858, 1.641] | 0.2692 | -0.609 | false | true |
| e2e | verify | array-256 | schubfach | 1.250x | [0.869, 1.829] | 0.2912 | -0.636 | false | true |
| e2e | verify | canonical-minimal | schubfach | 1.157x | [0.711, 1.785] | 0.5402 | -0.320 | false | true |
| e2e | verify | control-escapes | schubfach | 1.238x | [0.996, 1.568] | 0.0906 | -1.000 | false | true |
| e2e | verify | deep-64 | schubfach | 1.250x | [0.919, 1.908] | 0.2289 | -0.666 | false | true |
| e2e | verify | escaped-key-order | schubfach | 1.091x | [0.691, 1.645] | 0.7178 | -0.199 | false | true |
| e2e | verify | long-string | schubfach | 1.033x | [0.648, 1.601] | 0.8944 | -0.072 | false | true |
| e2e | verify | nested-mixed | json-canon | 1.092x | [0.854, 1.548] | 0.6101 | 0.304 | false | true |
| e2e | verify | numeric-boundary | schubfach | 1.411x | [0.903, 2.543] | 0.1866 | -0.735 | false | true |
| e2e | verify | rfc-key-sorting | json-canon | 1.054x | [0.892, 1.338] | 0.7304 | 0.262 | false | true |
| e2e | verify | small | schubfach | 1.003x | [0.657, 1.440] | 0.9937 | -0.008 | false | false |
| e2e | verify | surrogate-pair | schubfach | 1.086x | [0.787, 1.482] | 0.6315 | -0.264 | false | true |
| e2e | verify | unicode | json-canon | 1.116x | [0.880, 1.556] | 0.4585 | 0.397 | false | true |
| e2e | verify | verify-whitespace | schubfach | 1.139x | [0.806, 1.725] | 0.5202 | -0.352 | false | true |
