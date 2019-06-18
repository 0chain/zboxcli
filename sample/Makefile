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

ALLOCATION:=--allocation ae5b9040bac6d102c149ca1262b8f9c7baf0b13ff01e7a773b969d4044d690a7

ZCN_DIR=$(HOME)/.zcn

setup:
	brew install jq || brew upgrade jq || true
	@mkdir -p $(ZCN_DIR) $(UPLOAD_DIR) $(DOWNLOAD_DIR) || true
	cp $(CLUSTERS)/*.yml $(ZCN_DIR)

generate-local:
	base64 /dev/urandom | dd of=$(LOCAL_FILE) bs=16k count=10
	#dd if=/dev/urandom of=$(LOCAL_FILE) bs=16k count=16

register:
	$(ZBOXCLI_CMD) $(CLIENT) register

newallocation:
	$(ZBOXCLI_CMD) $(CLIENT) newallocation

upload-file:
	@echo "Upload local=$(LOCAL_PATH) remote=$(REMOTE_PATH)"
	$(ZBOXCLI_CMD) $(CLIENT) upload $(ALLOCATION) $(LOCAL_PATH) $(REMOTE_PATH)

update-file:
	$(ZBOXCLI_CMD) $(CLIENT) update $(ALLOCATION) $(LOCAL_PATH) $(REMOTE_PATH)

stats:
	$(ZBOXCLI_CMD) $(CLIENT) stats $(ALLOCATION) $(REMOTE_PATH)

clean:


