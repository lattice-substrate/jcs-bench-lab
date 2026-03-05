#!/usr/bin/env bash
set -euo pipefail

# Package reproducibility artifact for archival submission.
# Creates a self-contained tarball with source, results, oracle vectors, and paper.

STAMP=$(date -u +%Y%m%d)
ARTIFACT="jcs-bench-lab-artifact-${STAMP}"
ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "Packaging artifact from ${ROOT}..."

TMPDIR=$(mktemp -d)
DEST="${TMPDIR}/${ARTIFACT}"
mkdir -p "${DEST}"

# Source code (excluding binaries and build artifacts)
rsync -a --exclude='bin/' --exclude='*.exe' --exclude='.git/' \
  --exclude='workloads/' --exclude='__debug_bin*' \
  "${ROOT}/" "${DEST}/"

# Ensure results directory is included
mkdir -p "${DEST}/results"
if ls "${ROOT}/results/latest-"* 1>/dev/null 2>&1; then
  cp "${ROOT}/results/latest-"* "${DEST}/results/"
fi

# Oracle vectors (already in source tree under impl-*/jcsfloat/testdata/)

# Paper (if built)
if [ -f "${ROOT}/paper/main.pdf" ]; then
  cp "${ROOT}/paper/main.pdf" "${DEST}/paper/"
fi

# Generate manifest
cd "${DEST}"
find . -type f | sort | while read -r f; do
  sha256sum "$f"
done > MANIFEST.sha256

cd "${TMPDIR}"
tar czf "${ROOT}/${ARTIFACT}.tar.gz" "${ARTIFACT}"
rm -rf "${TMPDIR}"

echo "Artifact: ${ROOT}/${ARTIFACT}.tar.gz"
echo "Verify: tar xzf ${ARTIFACT}.tar.gz && cd ${ARTIFACT} && sha256sum -c MANIFEST.sha256"
