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

if [ "${CIRCLECI}" != "true" ] && [ "${CI}" != "true" ]; then
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
if [ "${CIRCLECI}" == "true" ] || [ "${CI}" == "true" ]; then
    go test -v -p 1 -covermode=count -coverprofile=coverage.out -failfast ./src/...
    goveralls -coverprofile=coverage.out -service=travis-ci
else
    gotest -v -failfast -p 1 -coverprofile=coverage.out ./src/...
    go tool cover -html ./coverage.out -o coverage.html
fi
echo "PASS"
echo
