#!/usr/bin/env make
# ---------------------------------------------------------------------------------------------------------------------

GO = go

PACKAGE = github.com/brgv/sv

# ---------------------------------------------------------------------------------------------------------------------

GOFLAGS ?= -v

LDFLAGS ?=

LDFLAGS += -X '$(PACKAGE)/version.Version=v$(VERSION)'
LDFLAGS += -X '$(PACKAGE)/version.Executable=$(EXECUTABLE)'
LDFLAGS += -X '$(PACKAGE)/version.BuildNumber=$(BUILD_NUMBER)'
LDFLAGS += -X '$(PACKAGE)/version.BuildBranch=$(BUILD_BRANCH)'
LDFLAGS += -X '$(PACKAGE)/version.BuildHash=$(BUILD_HASH)'
LDFLAGS += -X '$(PACKAGE)/version.BuildTime=$(BUILD_TIME)'
LDFLAGS += -X '$(PACKAGE)/version.BuildUser=$(BUILD_USER)'

STATIC ?=
ifeq ($(STATIC),1)
LDFLAGS += -s -w -extldflags "-static"
endif

# ---------------------------------------------------------------------------------------------------------------------

.PHONY: go-compile
go-compile:
	@echo "-----------------------------------------------------------------------------"
ifeq ($(os),windows)
	GOOS=$(os) GOARCH=$(arch) GOARM=7 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(d)/$(o).exe $(s)
	#@file $(d)/$(o).exe
	@ls -sh $(d)/$(o).exe
else
	GOOS=$(os) GOARCH=$(arch) GOARM=7 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(d)/$(o) $(s)
	#@file $(d)/$(o)
	@ls -sh $(d)/$(o)
endif
	@echo "-----------------------------------------------------------------------------"

# ---------------------------------------------------------------------------------------------------------------------
