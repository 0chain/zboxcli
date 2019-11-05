ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

SAMPLE_DIR:=$(ROOT_DIR)/sample

ZBOXCLI=zboxcli

.PHONY:

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

zboxcli: gomod-download
	$(eval VERSION=$(shell git describe --tags --dirty --always))
	$(eval VERSION=$(VERSION)-$(shell git rev-list -1 HEAD --abbrev-commit))
	go build -x -v -tags bn256 -ldflags "-X main.VersionStr=$(VERSION)" -o $(ZBOXCLI) main.go


zboxcli-test:
	go test -tags bn256 ./...

install: zboxcli zboxcli-test
	cp $(ZBOXCLI) $(ROOT_DIR)/sample

clean: gomod-clean
	@rm -rf $(ROOT_DIR)/$(ZBOXCLI)

help:
	@echo "Environment: "
	@echo "\tGOPATH=$(GOPATH)"
	@echo "\tGOROOT=$(GOROOT)"
	@echo ""
	@echo "Supported commands:"
	@echo "\tmake help              - display environment and make targets"
	@echo ""
	@echo "Install"
	@echo "\tmake install           - build, test and install zboxcli"
	@echo "\tmake zboxcli           - installs the zboxcli"
	@echo "\tmake zboxcli-test      - run zboxcli test"
	@echo ""
	@echo "Clean:"
	@echo "\tmake clean             - deletes all build output files"
	@echo "\tmake gomod-download    - download the go modules"
	@echo "\tmake gomod-clean       - clean the go modules"
