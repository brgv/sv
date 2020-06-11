#!/usr/bin/env make
# ---------------------------------------------------------------------------------------------------------------------

GO = go

PACKAGE = github.com/brgv/sv

# ---------------------------------------------------------------------------------------------------------------------

GOFLAGS ?= -v

LDFLAGS ?=

LDFLAGS += -X '$(PACKAGE)/internal/version.Version=$(VERSION)'
LDFLAGS += -X '$(PACKAGE)/internal/version.Executable=$(EXECUTABLE)'
LDFLAGS += -X '$(PACKAGE)/internal/version.BuildNumber=$(BUILD_NUMBER)'
LDFLAGS += -X '$(PACKAGE)/internal/version.BuildBranch=$(BUILD_BRANCH)'
LDFLAGS += -X '$(PACKAGE)/internal/version.BuildHash=$(BUILD_HASH)'
LDFLAGS += -X '$(PACKAGE)/internal/version.BuildTime=$(BUILD_TIME)'
LDFLAGS += -X '$(PACKAGE)/internal/version.BuildUser=$(BUILD_USER)'

STATIC ?=
ifeq ($(STATIC),1)
LDFLAGS += -s -w -extldflags "-static"
endif

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: go-compile
go-compile:
	@echo "-----------------------------------------------------------------------------"
ifeq ($(os),windows)
	GOOS=$(os) GOARCH=$(arch) $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(d)/$(o).exe $(s)
	#@file $(d)/$(o).exe
	@ls -sh $(d)/$(o).exe
else
	GOOS=$(os) GOARCH=$(arch) $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(d)/$(o) $(s)
	#@file $(d)/$(o)
	@ls -sh $(d)/$(o)
endif
	@echo "-----------------------------------------------------------------------------"

# ---------------------------------------------------------------------------------------------------------------------
