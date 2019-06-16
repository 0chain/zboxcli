ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

ZBOXCMD=zboxcmd

GOMODCORE           := $(GOMODBASE)/zcncore
VERSION_FILE        := $(ROOT_DIR)/core/version/version.go
MAJOR_VERSION       := "1.0"

PLATFORMOS := $(shell uname | tr "[:upper:]" "[:lower:]")

include _util/printer.mk

.PHONY: install-all herumi-all gosdk-all show

default: help show

#GO BUILD SDK
gomod-download:
	go mod download -json

gomod-clean:
	go clean -i -r -x -modcache  ./...

build: gomod-download
	go build -x -v -tags bn256 -o $(ZBOXCMD) main.go

zboxcmd-test:
	go test -tags bn256 ./...

install: build
	cp $(ZBOXCMD) $(ROOT_DIR)/sample

clean:
	@rm -rf $(OUTDIR)

show:
	@echo "GOPATH=$(GOPATH)"
	@echo "GOROOT=$(GOROOT)"
	@echo "BLS git branch=$(bls_branch)"
	@echo "MCL git branch=$(mcl_branch)"

help:
	@echo "Supported commands:"
	@ecgo "\tmake show              - Display environment and make variables"
	@echo "\tmake install       - Install all build and project dependencies"
	@echo "\tmake gosdk-all         - Install GO modules and packages"
	@echo "\tmake herumi-all        - Download, build and install HERUMI packages"
	@echo "\tmake clean             - Deletes all the built output files"
