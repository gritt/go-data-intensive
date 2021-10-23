include Makefile.vars

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info Available Commands:)
	$(info -> install                 installs project dependencies)
	$(info -> test-unit               run unit tests)
	$(info -> run                     run app)

.PHONY: install
install:
	go mod download

.PHONY: test-unit
test-unit: install
	go test ./... -v -short -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: run
run: install
	go run $(SERVERDIR)

# ignore unknown commands
%:
    @:
