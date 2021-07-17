#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

set OTELCOL_IMG="otelcol:latest"
docker-compose --project-directory "${ROOT}"/hack/opentelemetry-collector/ up -d
