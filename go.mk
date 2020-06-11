#!/usr/bin/env make
# ---------------------------------------------------------------------------------------------------------------------

GO=go

EXECUTABLE ?= sv

DIR_DIST=./dist
DIR_BIN=./bin
DIR_CMD=./cmd

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: go-clean
go-clean:
	@$(GO) clean
	@rm -rf $(DIR_BIN)
	@rm -rf $(DIR_DIST)

.PHONY: go-build
go-build:
	@$(GO) version
	@make go-compile d=${DIR_BIN} o=${EXECUTABLE}ctl
	@make go-compile d=${DIR_BIN} o=${EXECUTABLE}-gate s=${DIR_CMD}/gate/
	@make go-compile d=${DIR_BIN} o=${EXECUTABLE}-grpc s=${DIR_CMD}/grpc/
	@make go-compile d=${DIR_BIN} o=${EXECUTABLE}-sync s=${DIR_CMD}/sync/

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: go-test
go-test:
	@$(GO) test -v -short ./...

# ---------------------------------------------------------------------------------------------------------------------
