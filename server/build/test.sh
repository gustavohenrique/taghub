#!/bin/sh

export CGO_ENABLED=1
export GO111MODULE=on

echo -n "Checking gofmt: "
ERRS=$(find ./src ./libs -type f -name \*.go | xargs gofmt -l 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL - the following files need to be gofmt'ed:"
    for e in ${ERRS}; do
        echo "    $e"
    done
    echo
    exit 1
fi
echo "PASS"
echo

if [ "${CIRCLECI}" != "true" ]; then
    echo -n "Checking go vet: "
    TARGETS="./src/... ./libs/..."
    ERRS=$(go vet ${TARGETS} 2>&1 || true)
    if [ -n "${ERRS}" ]; then
        echo "FAIL"
        echo "${ERRS}"
        echo
        exit 1
    fi
    echo "PASS"
    echo
fi

echo -n "Running tests: "
export DATABASE_URL=${1}
if [ "${CIRCLECI}" == "true" ]; then
    go test -v -p 1 -coverprofile=coverage.txt -failfast ./src/...
else
    gotest -v -failfast -p 1 -coverprofile=coverage.txt ./src/...
fi
go tool cover -html ./coverage.txt -o coverage.html
echo "PASS"
echo
