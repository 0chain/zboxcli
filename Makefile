PATH  := $(PATH):$(PWD)
#SHELL := env PATH=$(PATH) /bin/bash

ZBOXCMD_BASE=$(PWD)

SCHEME?=ed25519
#SCHEME?=bls0chain

include scheme/$(SCHEME).mk

CLUSTERS=$(ZBOXCMD_BASE)/clusters

CONFIG:=--config $(cluster)

CLIENT:=--wallet $(client)

ZBOXCMD:=$(ZBOXCMD_BASE)/zboxcmd $(CONFIG)

UPLOAD_DIR=$(ZBOXCMD_BASE)/upload
DOWNLOAD_DIR=$(ZBOXCMD_BASE)/download

CLIENT_FILE=random1.txt
LOCAL_FILE:=$(UPLOAD_DIR)/$(CLIENT_FILE)
LOCAL_PATH=--localpath $(LOCAL_FILE)
REMOTE_PATH=--remotepath /0chain/$(CLIENT_FILE)

ALLOCATION:=--allocation f0a7eea03124eb0ba0698efe70954075d1ce1268dd119981ee304382e74eb152
FAUCET:=faucet --methodName pour --input "Generous PayDay"

#TO_CLIENT_ID:=$(shell jq -r '.client_id' $(HOME)/.zcn/$(to))
#FROM_CLIENT_ID:=$(shell jq -r '.client_id' $(HOME)/.zcn/$(from))

init: clean
	@echo "Creating New Wallets"
	$(ZCMD) getbalance $(FROM_WALLET)
	$(ZCMD) getbalance $(TO_WALLET)

show:

ZCN_DIR=$(HOME)/.zcn

setup:
	@brew install jq
	@mkdir -p $(ZCN_DIR) $(UPLOAD_DIR) $(DOWNLOAD_DIR) || true
	cp clusters/*.yml $(ZCN_DIR)

write_local:
	dd if=/dev/urandom of=$(LOCAL_FILE) bs=16k count=16

register:
	$(ZBOXCMD) $(CLIENT) register


newallocation:
	$(ZBOXCMD) $(CLIENT) newallocation

upload_file:
	$(ZBOXCMD) $(CLIENT) upload $(ALLOCATION) $(LOCAL_PATH) $(REMOTE_PATH)

update_file:
	$(ZBOXCMD) $(CLIENT) update $(ALLOCATION) $(LOCAL_PATH) $(REMOTE_PATH)

stats:
	$(ZBOXCMD) $(CLIENT) stats $(ALLOCATION) $(REMOTE_PATH)

send:
	$(ZCMD) $(FROM_WALLET) send --desc "Give loan - Make Happy" --toclientID $(TO_CLIENT_ID) --token 1.5

pay-chain: getbalance0 | send getbalance1 repay getbalance2

repay:
	$(ZCMD) $(TO_WALLET) send --desc "Return loan - Feel sad" --toclientID $(FROM_CLIENT_ID) --token 1.5

getbalance0 getbalance1 getbalance2 getbalance:
	$(ZCMD) getbalance $(FROM_WALLET)
	$(ZCMD) getbalance $(TO_WALLET)

getrich:
	for i in `seq 10`; do make faucet; done

faucet:
	@echo "Add money receiver=$(from)"
	$(ZCMD) $(FAUCET) $(FROM_WALLET)
	@echo "Add money receiver=$(to)"
	$(ZCMD) $(FAUCET) $(TO_WALLET)

clean:
	cd $(HOME)/.zcn;  rm $(from) $(to) || true


