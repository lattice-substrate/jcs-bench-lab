# Statistical Comparison

Generated at: 2026-03-05T16:39:25Z

Source: `/home/lenny/jcs-bench-lab/results/latest-cli-runs.csv`

Protocol: permutation test + bootstrap 95% CI, alpha=0.0500, resamples=5000

| track | mode | workload | winner | speedup | ci95 | p-value | effect d | significant | practical | noise floor (ms) | min obs effect (%) |
|---|---|---|---|---:|---|---:|---:|---|---|---:|---:|
| e2e | canonicalize | array-2048 | schubfach | 1.343x | [1.259, 1.432] | 0.0002 | -3.180 | true | true | 0.5401 | 5.10 |
| e2e | canonicalize | array-256 | schubfach | 1.156x | [1.030, 1.305] | 0.0296 | -0.848 | true | true | 0.3533 | 9.01 |
| e2e | canonicalize | canonical-minimal | schubfach | 1.063x | [0.939, 1.207] | 0.3595 | -0.336 | false | true | 0.2742 | 9.62 |
| e2e | canonicalize | control-escapes | schubfach | 1.048x | [0.940, 1.168] | 0.4201 | -0.299 | false | true | 0.2229 | 7.78 |
| e2e | canonicalize | deep-64 | schubfach | 1.055x | [0.913, 1.221] | 0.4797 | -0.257 | false | true | 0.3809 | 12.73 |
| e2e | canonicalize | escaped-key-order | schubfach | 1.099x | [0.995, 1.218] | 0.0886 | -0.652 | false | true | 0.2284 | 7.84 |
| e2e | canonicalize | long-string | schubfach | 1.114x | [0.986, 1.263] | 0.1062 | -0.614 | false | true | 0.2862 | 9.47 |
| e2e | canonicalize | nested-mixed | schubfach | 1.129x | [0.999, 1.275] | 0.0692 | -0.693 | false | true | 0.2573 | 9.10 |
| e2e | canonicalize | numeric-boundary | schubfach | 1.090x | [0.959, 1.236] | 0.2030 | -0.468 | false | true | 0.2825 | 9.93 |
| e2e | canonicalize | rfc-key-sorting | schubfach | 1.123x | [1.005, 1.265] | 0.0620 | -0.712 | false | true | 0.2772 | 9.87 |
| e2e | canonicalize | small | schubfach | 1.056x | [0.940, 1.196] | 0.3907 | -0.318 | false | true | 0.2643 | 9.65 |
| e2e | canonicalize | surrogate-pair | schubfach | 1.191x | [1.043, 1.355] | 0.0168 | -0.940 | true | true | 0.2705 | 10.56 |
| e2e | canonicalize | unicode | schubfach | 1.131x | [0.993, 1.289] | 0.0810 | -0.665 | false | true | 0.2796 | 10.10 |
| e2e | canonicalize | verify-whitespace | json-canon | 1.004x | [0.878, 1.133] | 0.9488 | 0.024 | false | false | 0.2420 | 8.47 |
| e2e | verify | array-2048 | schubfach | 1.285x | [1.195, 1.378] | 0.0002 | -2.417 | true | true | 0.6784 | 6.14 |
| e2e | verify | array-256 | schubfach | 1.155x | [1.027, 1.308] | 0.0298 | -0.840 | true | true | 0.4021 | 10.84 |
| e2e | verify | canonical-minimal | schubfach | 1.071x | [0.948, 1.202] | 0.2849 | -0.398 | false | true | 0.2308 | 7.89 |
| e2e | verify | control-escapes | json-canon | 1.081x | [0.953, 1.238] | 0.2480 | 0.429 | false | true | 0.2978 | 10.23 |
| e2e | verify | deep-64 | schubfach | 1.156x | [1.022, 1.314] | 0.0310 | -0.813 | true | true | 0.2877 | 10.04 |
| e2e | verify | escaped-key-order | schubfach | 1.101x | [0.988, 1.235] | 0.1020 | -0.607 | false | true | 0.2510 | 9.02 |
| e2e | verify | long-string | schubfach | 1.179x | [1.022, 1.353] | 0.0344 | -0.810 | true | true | 0.3153 | 10.50 |
| e2e | verify | nested-mixed | schubfach | 1.033x | [0.896, 1.188] | 0.6779 | -0.155 | false | true | 0.2251 | 7.85 |
| e2e | verify | numeric-boundary | schubfach | 1.055x | [0.918, 1.214] | 0.4797 | -0.263 | false | true | 0.2996 | 10.27 |
| e2e | verify | rfc-key-sorting | schubfach | 1.018x | [0.894, 1.159] | 0.7858 | -0.094 | false | false | 0.2656 | 9.45 |
| e2e | verify | small | schubfach | 1.003x | [0.878, 1.156] | 0.9676 | -0.016 | false | false | 0.2964 | 10.49 |
| e2e | verify | surrogate-pair | schubfach | 1.100x | [0.955, 1.263] | 0.1938 | -0.482 | false | true | 0.2694 | 10.02 |
| e2e | verify | unicode | schubfach | 1.216x | [1.065, 1.390] | 0.0120 | -1.029 | true | true | 0.2642 | 10.42 |
| e2e | verify | verify-whitespace | schubfach | 1.150x | [1.004, 1.327] | 0.0618 | -0.715 | false | true | 0.2786 | 10.37 |
