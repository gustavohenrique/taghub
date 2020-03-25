#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

BINARY=${1}
# OS=$(go env GOOS)
# ARCH=$(go env GOARCH)
OS="linux"
ARCH="amd64"
VERSION=$(git describe --tags --always)
NOW=$(date +"%Y%m%d_%H%M")

echo "Building ${NOW} ${VERSION}"

# Disable C code, enable Go modules
export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOOS="${OS}"
export GO111MODULE=on
export GOPRIVATE="github.com/estrategiahq"

APP_NAME=$(go list -m)
go build \
    -v \
    -installsuffix "static" \
    -ldflags "-X ${APP_NAME}/pkg.VERSION=${VERSION} -X ${APP_NAME}/pkg.DATETIME=${NOW} -s -w" \
    -o "${BINARY}-${VERSION}"
