# Adversarial Peer Review: Issue Tracker

Generated: 2026-03-05
Status: IN PROGRESS

Each issue below maps to the adversarial peer review findings. Check the box when resolved.

---

## Fatal Flaws

### F1. API Speedup Claims Have No Inferential Statistical Backing
- [x] **Fixed**
- **Severity:** FATAL
- **Location:** Abstract (`paper/main.tex:36`), `paper/sections/discussion.tex:9`, `results/latest-stats.json` (missing API track)
- **Problem:** The headline claim "1.3-1.7x faster on number-heavy workloads in API benchmarks" has zero inferential statistics. `latest-stats.json` contains only `e2e` (CLI) track. API benchmarks are analyzed via raw `testing.B` output with no permutation tests, no BCa CIs, no BH correction. The benchstat summary explicitly says "benchstat unavailable; fallback summary above is non-inferential and not CI-gating."
- **Fix applied:** Extended stats pipeline (`cmd/lab/stats.go`) with `loadAPIBenchRuns()` to parse Go `testing.B` output into `runRecord` entries. These are now processed through the same permutation test + BCa CI + BH correction pipeline as CLI data. Added `--api-bench` flag to the `stats` subcommand. API comparisons appear as `track: "api"` in `latest-stats.json`.

### F2. ARM64 Oracle Test Results Hardcoded as PASS Despite Failures
- [x] **Fixed**
- **Severity:** FATAL
- **Location:** `cmd/lab/main.go:1593-1594`
- **Problem:** `GoldenVectors` and `StressVectors` are hardcoded to `"PASS"` regardless of actual test results. `results/latest-arm64-determinism.json` shows `total_tests: 24, total_passed: 22` for both implementations (2 failures each), yet reports "PASS".
- **Fix applied:** Changed `cmd/lab/main.go` to derive `goldenResult`/`stressResult` from actual test output by matching `TestGoldenOracle` and `TestStressOracle` in `--- PASS:`/`--- FAIL:` lines. Default is `"SKIP"` if test name not found. Evidence file must be regenerated.

---

## Major Issues

### M1. CLI Speedup Range "1.1-1.4x" is Overstated
- [x] **Fixed**
- **Severity:** MAJOR
- **Location:** Abstract (`paper/main.tex:38`), `paper/sections/discussion.tex:11`
- **Problem:** Only 5 of 28 CLI comparisons are BH-significant (speedups 1.21x-1.39x). The "1.1x" lower bound is not supported by any significant result. 23 of 28 comparisons show no significant difference. Many CIs include 1.0.
- **Fix applied:** Revised abstract, introduction, discussion, and conclusion to state "BH-significant comparisons show 1.2-1.4x speedup, while the majority of workloads show no significant difference." Replaced "number-heavy" with "number-dense" to be more precise.

### M2. API "Number-Heavy" Definition is Cherry-Picked
- [x] **Fixed**
- **Severity:** MAJOR
- **Location:** `paper/sections/discussion.tex:6-15`
- **Problem:** `deep` and `deep-64` workloads contain numbers but show ~1.0x speedup. `verify-whitespace` and `escaped-key-order` are not number-heavy but show 1.6-1.7x. The "number-heavy" framing doesn't fully explain the performance pattern.
- **Fix applied:** Added explicit discussion of `deep`/`deep-64` showing ~1.0x despite containing numbers, explaining that speedup correlates with number *density* (fraction of cost from digit generation), not merely number presence. Changed terminology from "number-heavy" to "number-dense."

### M3. Min Observable Effect Formula Missing sqrt(2) Factor
- [x] **Fixed**
- **Severity:** MAJOR
- **Location:** `cmd/lab/stats.go:229-233`, `paper/sections/methodology.tex:49-52`
- **Problem:** Two-sample formula requires `2.8 * sqrt(2) * CV * 100 / sqrt(n)` but code uses `2.8 * CV * 100 / sqrt(n)`. Underestimates minimum detectable effect by ~29%. Also uses only faster implementation's CV/n.
- **Fix applied:** Corrected formula in `cmd/lab/stats.go` to include `math.Sqrt2` factor. Updated comment to explain the two-sample derivation. Updated `methodology.tex` to show corrected formula with sqrt(2) and caveat about single-impl approximation.

### M4. No Statistical Analysis for API Benchmarks in stats.json
- [x] **Fixed**
- **Severity:** MAJOR (overlaps F1)
- **Location:** `cmd/lab/stats.go`, `cmd/lab/main.go`
- **Problem:** The stats subcommand processes only CLI data. API comparisons lack permutation tests, CIs, effect sizes, and multiple comparison correction.
- **Fix applied:** Same as F1 — `loadAPIBenchRuns()` parser added, integrated into `runStats()`. BH correction now applied across both CLI and API comparisons jointly.

### M5. Conformance Gate is Bypassable
- [x] **Fixed**
- **Severity:** MAJOR
- **Location:** `cmd/lab/main.go:151`
- **Problem:** `--skip-conformance` flag lets users generate benchmark data without passing conformance. No field in results records whether the gate was active.
- **Fix applied:** Added `ConformanceGateActive bool` field to `qualityReport` struct. Set to `!skipConformance` in `runBenchCLI`. Gate subcommand (`cmd/lab/gate.go`) now reads `latest-quality.json` and rejects benchmarks that were run with `--skip-conformance`.

---

## Minor Issues

### m1. Noise Floor is Mislabeled
- [x] **Fixed**
- **Location:** `cmd/lab/stats.go:217-228`, `paper/sections/methodology.tex:48`
- **Problem:** `CV * mean = (sigma/mu) * mu = sigma`. The "noise floor" is just the standard deviation. Name implies something distinct.
- **Fix applied:** Updated code comment to say "standard deviation of the faster impl (equivalently CV x mean)". Updated methodology.tex to describe it as "standard deviation of the faster implementation."

### m2. `definition` LaTeX Environment Not Declared
- [x] **Fixed**
- **Location:** `paper/main.tex:12-13`, `paper/sections/algorithms.tex:6,14`
- **Problem:** `\newtheorem{definition}{Definition}` is missing; only `remark` is declared.
- **Fix applied:** Added `\theoremstyle{definition}` and `\newtheorem{definition}{Definition}` before the remark declaration in `paper/main.tex`.

### m3. "Only bits.Mul64" Claim is Imprecise
- [x] **Fixed**
- **Location:** `paper/sections/algorithms.tex:49-51`
- **Problem:** Claim says Schubfach "operates entirely within fixed-width 64/128-bit arithmetic via math/bits."
- **Fix applied:** Changed to "uses fixed-width 64/128-bit arithmetic, relying on math/bits for 128-bit multiplication."

### m4. Process Startup "~1.3 ms" Claimed Without Direct Measurement
- [x] **Fixed**
- **Location:** `paper/sections/discussion.tex:12,19`, `paper/sections/results.tex:25`
- **Problem:** The 1.3 ms figure is inferred, not directly measured.
- **Fix applied:** Qualified as "estimated process startup overhead (approximately 1.3 ms, inferred from the difference between minimal-workload CLI and API times)" in all occurrences.

### m5. PROVENANCE.json "Stable Across V8 Versions" Unverified
- [x] **Fixed**
- **Location:** `paper/sections/threats.tex`
- **Problem:** Stability across V8 versions is claimed but only one version was tested.
- **Fix applied:** Added to threats.tex: "our vectors are verified against only one engine and one version (V8 12.9.202.28). Cross-version stability is expected from the normative spec but not independently verified."

### m6. Go Version Mismatch in go.mod vs Actual
- [x] **Fixed**
- **Location:** `go.mod:3`, `ARTIFACT.md`
- **Problem:** go.mod says `go 1.24`, results show `go1.25.6`.
- **Fix applied:** Updated ARTIFACT.md to note "reported results used Go 1.25.6" alongside the minimum version requirement.

### m7. QEMU Version Not Documented in ARTIFACT.md
- [x] **Fixed**
- **Location:** `ARTIFACT.md`
- **Problem:** arm64 tests require `qemu-aarch64-static` but this is not listed as a dependency.
- **Fix applied:** Added `qemu-aarch64-static` to ARTIFACT.md prerequisites with version note (tested with QEMU 8.2.2).

### m8. 87-93% Startup Overhead Claim Not Directly Verifiable
- [x] **Fixed**
- **Location:** `paper/sections/discussion.tex:20`
- **Problem:** Percentage derived from ratio of two estimates.
- **Fix applied:** Qualified with "roughly 87-93%" in conjunction with the "estimated" startup overhead qualifier.

---

## Remaining Action Items

- [ ] Regenerate `results/latest-arm64-determinism.json` with fixed oracle reporting code
- [ ] Regenerate `results/latest-stats.json` with API comparisons included (run `go run ./cmd/lab stats`)
- [ ] Regenerate paper tables from new stats data
- [ ] Update `results/baseline-stats.json` if gate baseline needs refreshing
- [ ] Verify LaTeX compilation with new `definition` environment
