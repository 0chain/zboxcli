SAMPLE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

ZBOXCMD_EXE=$(SAMPLE_DIR)/zboxcmd

SCHEME?=ed25519
#SCHEME?=bls0chain

include $(SAMPLE_DIR)/scheme/$(SCHEME).mk

CLUSTERS=$(SAMPLE_DIR)/clusters

CONFIG:=--config $(cluster)

CLIENT:=--wallet $(client)

ZBOXCMD:=$(ZBOXCMD_EXE) $(CONFIG)

UPLOAD_DIR=$(SAMPLE_DIR)/upload
DOWNLOAD_DIR=$(SAMPLE_DIR)/download

CLIENT_FILE=random5.txt
LOCAL_FILE:=$(UPLOAD_DIR)/$(CLIENT_FILE)
LOCAL_PATH=--localpath $(LOCAL_FILE)
REMOTE_PATH=--remotepath /0chain/$(CLIENT_FILE)

ALLOCATION:=--allocation 613feccd592f894fea1d83dd53b62848832cdac32e1bef4851a494869683d869

ZCN_DIR=$(HOME)/.zcn

setup:
	brew install jq || brew upgrade jq || true
	@mkdir -p $(ZCN_DIR) $(UPLOAD_DIR) $(DOWNLOAD_DIR) || true
	cp $(CLUSTERS)/*.yml $(ZCN_DIR)

generate-local:
	dd if=/dev/urandom of=$(LOCAL_FILE) bs=16k count=16

register:
	$(ZBOXCMD) $(CLIENT) register

newallocation:
	$(ZBOXCMD) $(CLIENT) newallocation

upload-file:
	@echo "Upload local=$(LOCAL_PATH) remote=$(REMOTE_PATH)"
	$(ZBOXCMD) $(CLIENT) upload $(ALLOCATION) $(LOCAL_PATH) $(REMOTE_PATH)

update-file:
	$(ZBOXCMD) $(CLIENT) update $(ALLOCATION) $(LOCAL_PATH) $(REMOTE_PATH)

stats:
	$(ZBOXCMD) $(CLIENT) stats $(ALLOCATION) $(REMOTE_PATH)

clean:


