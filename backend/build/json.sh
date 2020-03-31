#!/bin/sh

export DATABASE_URL=${1}
export EXPORT_DIR=${2}
rm -rf ${EXPORT_DIR} 2>/dev/null
mkdir -p ${EXPORT_DIR}/{repos,tags} 2>/dev/null
go run ./cmd/exporter/main.go
