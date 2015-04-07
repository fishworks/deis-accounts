SHELL = /bin/bash

GO = go
INSTALL = install
RM = rm -f

GOBUILD = $(GO) build -o
GOFMT = gofmt -l
GOLINT = golint
GOTEST = $(GO) test -cover --race
GOVET = $(GO) vet

prefix := /usr/local

# The directory to install deis-accounts in
bin_dir = $(prefix)/bin

# the filepath to this repository, relative to $GOPATH/src
repo_path = github.com/fishworks/deis-accounts

# used to reference the output directory for build artifacts
build_dir = bin

SRC_PACKAGES = deis version
REPO_SRC_PACKAGES = $(addprefix $(repo_path)/,$(SRC_PACKAGES))

all: build docs

build:
	$(GOBUILD) $(build_dir)/deis-accounts $(repo_path)

clean:
	$(RM) $(build_dir)/*
	$(MAKE) -C docs clean

docs:
	$(MAKE) -C docs

install:
	$(INSTALL) -c $(build_dir)/deis-accounts $(bin_dir)/deis-accounts
	$(MAKE) -C docs install

test:
# display output, then check
	$(GOFMT) $(SRC_PACKAGES)
	@$(GOFMT) $(SRC_PACKAGES) | read; if [ $$? == 0 ]; then echo "gofmt check failed."; exit 1; fi

# display output, then check
	$(GOLINT) ./...
	$(GOLINT) ./... | read; if [ $$? == 0 ]; then echo "golint check failed."; exit 1; fi

	$(GOTEST) $(REPO_SRC_PACKAGES)
	$(GOVET) $(REPO_SRC_PACKAGES)

.PHONY: all build clean docs install test
