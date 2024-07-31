#!/usr/bin/env bash

set -e

setup() {

    set -v
    export DEBIAN_FRONTEND=noninteractive
    apt-get update
    apt-get install --assume-yes --no-install-recommends apt-transport-https ca-certificates curl gnupg lsb-release software-properties-common apt-utils zstd dpkg
    set +v

    set -v
    mkdir -p /etc/apt/keyrings
    curl -sLS https://packages.zus.network/zus.asc |
      gpg --dearmor > /etc/apt/keyrings/zus.gpg
    chmod go+r /etc/apt/keyrings/zus.gpg
    set +v

    set -v
    # Use env var DIST_CODE for the package dist name if provided
    if [[ -z $DIST_CODE ]]; then
        CLI_REPO=$(lsb_release -cs)
        shopt -s nocasematch
        ERROR_MSG="Unable to find a package for your system. Please check if an existing package in https://packages.zus.network/aptrepo/dists/ can be used in your system and install with the dist name: 'curl -sL https://packages.zus.network/deb_install.sh | sudo DIST_CODE=<dist_code_name> bash'"
        if [[ ! $(curl -sL https://packages.zus.network/aptrepo/dists/) =~ $CLI_REPO ]]; then
            DIST=$(lsb_release -is)
            if [[ $DIST =~ "Ubuntu" ]]; then
                CLI_REPO="jammy"
            elif [[ $DIST =~ "Debian" ]]; then
                CLI_REPO="bookworm"
            elif [[ $DIST =~ "LinuxMint" ]]; then
                CLI_REPO=$(grep -Po 'UBUNTU_CODENAME=\K.*' /etc/os-release) || true
                if [[ -z $CLI_REPO ]]; then
                    echo "$ERROR_MSG"
                    exit 1
                fi
            else
                echo "$ERROR_MSG"
                exit 1
            fi
        fi
    else
        CLI_REPO=$DIST_CODE
        if [[ ! $(curl -sL https://packages.zus.network/aptrepo/dists/) =~ $CLI_REPO ]]; then
            echo "Unable to find an zbox-cli package with DIST_CODE=$CLI_REPO in https://packages.zus.network/aptrepo/dists/"
            exit 1
        fi
    fi

    if [ -f /etc/apt/sources.list.d/zbox.list ]; then
      rm /etc/apt/sources.list.d/zbox.list
    fi

    echo "Types: deb
URIs: https://packages.zus.network/aptrepo/
Suites: ${CLI_REPO}
Components: main
Architectures: $(dpkg --print-architecture)
Signed-by: /etc/apt/keyrings/zus.gpg" | tee /etc/apt/sources.list.d/zbox.sources
    apt-get update
    set +v

    apt-get install --assume-yes zbox

    # Create or update the config.yaml file
    CONFIG_DIR="${HOME}/.zcn"
    CONFIG_PATH="${CONFIG_DIR}/config.yaml"
    echo "Creating/updating configuration file at ${CONFIG_PATH}..."
    mkdir -p $CONFIG_DIR
    touch $CONFIG_PATH
    cat <<EOT > ${CONFIG_PATH}
---
block_worker: https://dev.zus.network
signature_scheme: bls0chain
min_submit: 50
min_confirmation: 50
confirmation_chain_length: 3

# # OPTIONAL - Uncomment to use/ Add more if you want
# preferred_blobbers:
#   - http://one.devnet-0chain.net:31051
#   - http://one.devnet-0chain.net:31052
#   - http://one.devnet-0chain.net:31053    
EOT

    echo "Installation and configuration complete."

}

setup  # ensure the whole file is downloaded before executing