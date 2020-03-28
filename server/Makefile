include .makerc

.PHONY: build mocks coverage

DB_URL  := "database.db"
TEST_DB_URL  := "/tmp/test.db"
GO := $(shell which go)
BINARY := taghub

test:
	DATABASE_URL=$(TEST_DB_URL) gotest -v -cover $(pkg)

tests:
	./build/test.sh $(TEST_DB_URL)

compile: build

build:
	./build/build.sh $(BINARY)

mocks:
	./build/mocks.sh

lint: linter

linter:
	goimports -w src libs

run:
	@GITHUB_PERSONAL_TOKEN=$(GITHUB_PERSONAL_TOKEN) \
	DATABASE_URL=$(DB_URL) \
	go run main.go

createdb:
	@sqlite3 $(DATABASE_URL) < sql/taghub.sql

setup:
	$(GO) get github.com/golang/mock/mockgen@latest
	$(GO) get golang.org/x/tools/cmd/goimports
	$(GO) get -u github.com/rakyll/gotest

json:
	@./build/json.sh $(DB_URL) $(PWD)/../web/src/statics/jsonfiles
