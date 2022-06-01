PLUGIN = hashicorp
WORKDIR = github.com/richbai90/porter-advanced-build-plugins
SHELL = bash

PORTER_HOME ?= $(HOME)/.porter

COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION ?= $(shell git describe --tags 2> /dev/null || echo v0)

GO = GO111MODULE=on go
LDFLAGS = -w -X $(WORKDIR)/pkg.Version=$(VERSION) -X $(WORKDIR)/pkg.Commit=$(COMMIT)
XBUILD = CGO_ENABLED=0 $(GO) build -a -tags netgo -ldflags '$(LDFLAGS)'
BINDIR = dist

CLIENT_PLATFORM ?= $(shell go env GOOS)
CLIENT_ARCH ?= $(shell go env GOARCH)
SUPPORTED_PLATFORMS = linux darwin windows
SUPPORTED_ARCHES = amd64

ifeq ($(CLIENT_PLATFORM),windows)
FILE_EXT=.exe
else
FILE_EXT=
endif

.PHONY: build
build:
	mkdir -p $(BINDIR)
	$(GO) build -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(PLUGIN)$(FILE_EXT) main.go

xbuild-all:
	$(foreach OS, $(SUPPORTED_PLATFORMS), \
		$(foreach ARCH, $(SUPPORTED_ARCHES), \
				$(MAKE) $(MAKE_OPTS) CLIENT_PLATFORM=$(OS) CLIENT_ARCH=$(ARCH) PLUGIN=$(PLUGIN) xbuild; \
		))

xbuild: $(BINDIR)/$(VERSION)/$(PLUGIN)-$(CLIENT_PLATFORM)-$(CLIENT_ARCH)$(FILE_EXT)
$(BINDIR)/$(VERSION)/$(PLUGIN)-$(CLIENT_PLATFORM)-$(CLIENT_ARCH)$(FILE_EXT):
	mkdir -p $(dir $@)
	GOOS=$(CLIENT_PLATFORM) GOARCH=$(CLIENT_ARCH) $(XBUILD) -o $@ main.go

test: test-unit
	$(BINDIR)/$(PLUGIN)$(FILE_EXT) version

test-unit: build
	$(GO) test ./...

install:
	mkdir -p $(PORTER_HOME)/plugins/$(PLUGIN)
	install $(BINDIR)/$(PLUGIN)$(FILE_EXT) $(PORTER_HOME)/plugins/$(PLUGIN)$(FILE_EXT)

clean:
	-rm -fr $(BINDIR)
