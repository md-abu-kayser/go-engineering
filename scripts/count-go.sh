#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT_DIR}"

MIN_GO_FILES=1111
TOTAL_GO_FILES="$(find . \
  -type f \
  -name '*.go' \
  -not -path './.git/*' \
  -not -path './vendor/*' \
  -not -path './.tmp/*' \
  | wc -l \
  | tr -d ' ')"

printf 'Go Engineering repository statistics\n'
printf '%s\n' '-------------------------------------'
printf 'Go source files : %s\n' "${TOTAL_GO_FILES}"
printf 'Required minimum : %s\n' "${MIN_GO_FILES}"

if (( TOTAL_GO_FILES < MIN_GO_FILES )); then
  printf '\nERROR: repository contains fewer than %s Go files.\n' "${MIN_GO_FILES}" >&2
  exit 1
fi

printf 'Status           : PASS\n'
