ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

SAMPLE_DIR:=$(ROOT_DIR)/sample

ZBOX=zbox
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

$(ZBOX): gomod-download
	$(eval VERSION=$(shell git describe --tags --dirty --always))
	go build -x -v -tags bn256 -ldflags "-X main.VersionStr=$(VERSION)" -o $(ZBOX) main.go

zboxcli-test:
	go test -tags bn256 ./...

install: $(ZBOX) zboxcli-test

clean: gomod-clean
	@rm -rf $(ROOT_DIR)/$(ZBOX)

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
	@echo "\tmake zbox              - installs the zboxcli"
	@echo "\tmake zboxcli-test      - run zboxcli test"
	@echo ""
	@echo "Clean:"
	@echo "\tmake clean             - deletes all build output files"
	@echo "\tmake gomod-download    - download the go modules"
	@echo "\tmake gomod-clean       - clean the go modules"

# todo: remove this before merging
#migrate-s3:
#	go run main.go migrate-s3 --bucket="iamrz1-migration" --region="us-east-2" --allocation $(ALLOC)
#
#s3-all:
#	go run main.go migrate-s3 --allocation $(ALLOC)
