#!/usr/bin/env bash
set -euo pipefail
count=$(find . -name '*.go' -type f | wc -l | tr -d ' ')
echo "Go files: $count"
if [ "$count" -lt 1111 ]; then echo "ERROR: repository has fewer than 1,111 Go files"; exit 1; fi
