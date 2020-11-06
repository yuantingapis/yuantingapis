#!/usr/bin/env -S bash -e

source cmd/lib.sh || {
  echo "Are you at repo root?"
  exit 1
}

api-linter -I $API_COMMON_PROTOS yuanting/yt/v1/*.proto
