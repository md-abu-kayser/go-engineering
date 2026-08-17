#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT_DIR}"

printf '%s\n' '============================================='
printf '%s\n' ' Go Engineering — Repository Verification'
printf '%s\n' '============================================='

printf '\n[1/7] Go file-count invariant\n'
./scripts/count-go.sh

printf '\n[2/7] gofmt check\n'
unformatted="$(gofmt -l .)"
if [[ -n "${unformatted}" ]]; then
  printf '%s\n' "The following files are not gofmt-formatted:" >&2
  printf '%s\n' "${unformatted}" >&2
  exit 1
fi
printf '%s\n' 'PASS: all Go files are gofmt-formatted.'

printf '\n[3/7] module verification\n'
go mod verify

printf '\n[4/7] go vet\n'
go vet ./...

printf '\n[5/7] tests\n'
go test ./...

printf '\n[6/7] lesson-index generator\n'
expected="$(mktemp)"
trap 'rm -f "${expected}"' EXIT
go run ./tools/lesson_index.go > "${expected}"
if ! cmp -s docs/LESSON_INDEX.json "${expected}"; then
  printf '%s\n' 'ERROR: docs/LESSON_INDEX.json is not reproducible from tools/lesson_index.go.' >&2
  diff -u docs/LESSON_INDEX.json "${expected}" || true
  exit 1
fi
printf '%s\n' 'PASS: lesson index is reproducible.'

printf '\n[7/7] repository summary\n'
printf 'Go files   : %s\n' "$(find . -type f -name '*.go' -not -path './.git/*' -not -path './vendor/*' | wc -l | tr -d ' ')"
printf 'Lessons    : %s\n' "$(find . -type f -name 'main.go' -path './level-*/*/main.go' | wc -l | tr -d ' ')"
printf 'Test files : %s\n' "$(find . -type f -name '*_test.go' -not -path './.git/*' | wc -l | tr -d ' ')"
printf 'Projects   : %s\n' "$(find projects -mindepth 1 -maxdepth 1 -type d 2>/dev/null | wc -l | tr -d ' ')"

printf '\n%s\n' 'ALL REPOSITORY CHECKS PASSED.'
