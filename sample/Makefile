SAMPLE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

ZBOXCLI=$(SAMPLE_DIR)/zboxcli

SCHEME?=ed25519

include $(SAMPLE_DIR)/scheme/$(SCHEME).mk

CLUSTERS=$(SAMPLE_DIR)/clusters

CONFIG:=--config $(cluster)

CLIENT:=--wallet $(client)

ZBOXCLI_CMD:=$(ZBOXCLI) $(CONFIG)

UPLOAD_DIR=$(SAMPLE_DIR)/upload
DOWNLOAD_DIR=$(SAMPLE_DIR)/download

CLIENT_FILE=random.txt
LOCAL_FILE:=$(UPLOAD_DIR)/$(CLIENT_FILE)
LOCAL_PATH=--localpath $(LOCAL_FILE)
REMOTE_PATH=--remotepath /0chain/$(CLIENT_FILE)

ZCN_DIR=$(HOME)/.zcn

ALLOCATION_FILE=allocation.txt
ALLOCATION=--allocationFileName $(ALLOCATION_FILE)
ALLOCATION_ID=`cat $(HOME)/.zcn/$(ALLOCATION_FILE)`
ALLOCATION_ID_OPTION=--allocation $(ALLOCATION_ID)

setup:
	brew install jq || brew upgrade jq || true
	@mkdir -p $(ZCN_DIR) $(UPLOAD_DIR) $(DOWNLOAD_DIR) || true
	cp $(CLUSTERS)/*.yml $(ZCN_DIR)

generate-local:
	base64 /dev/urandom | dd of=$(LOCAL_FILE) bs=16k count=10

register:
	$(ZBOXCLI_CMD) $(CLIENT) register

newallocation:
	$(ZBOXCLI_CMD) $(CLIENT) $(ALLOCATION) newallocation

upload-file:
	@echo "Upload local=$(LOCAL_PATH) remote=$(REMOTE_PATH)"
	$(ZBOXCLI_CMD) $(CLIENT) upload $(ALLOCATION_ID_OPTION) $(LOCAL_PATH) $(REMOTE_PATH)

update-file:
	$(ZBOXCLI_CMD) $(CLIENT) update $(ALLOCATION_ID_OPTION) $(LOCAL_PATH) $(REMOTE_PATH)

stats:
	$(ZBOXCLI_CMD) $(CLIENT) stats $(ALLOCATION_ID_OPTION) $(REMOTE_PATH)

clean:
	@rm -rf $(LOCAL_FILE)

help:
	@echo "Environment:"
	@echo "\tScheme            : $(SCHEME)"
	@echo "\tClient            : $(client)"
	@echo "\tCluster Config    : $(HOME)/.zcn/$(cluster)"
	@echo ""
	@echo "Supported Targets:"
	@echo "\tmake register       : Initialize Wallets"
	@echo "\tmake generate-local : Generate local file with random text data"
	@echo "\tmake newallocation  : New allocation for storage"
	@echo "\tmake upload-file    : Upload a new file"
	@echo "\tmake update-file    : Update contents of an existing file"
	@echo "\tmake stats          : Display statistics for the remote file"
