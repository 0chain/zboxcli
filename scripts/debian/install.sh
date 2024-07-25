#!/usr/bin/env bash

set -e

if [[ $# -ge 1 && $1 == "-y" ]]; then
    global_consent=0
else
    global_consent=1
fi

function assert_consent {
    if [[ $2 -eq 0 ]]; then
        return 0
    fi

    echo -n "$1 [Y/n] "
    read consent
    if [[ ! "${consent}" == "y" && ! "${consent}" == "Y" && ! "${consent}" == "" ]]; then
        echo "'${consent}'"
        exit 1
    fi
}

setup() {

    assert_consent "Add packages necessary to modify your apt-package sources?" ${global_consent}
    set -v
    export DEBIAN_FRONTEND=noninteractive
    apt-get update
    apt-get install --assume-yes --no-install-recommends apt-transport-https ca-certificates curl gnupg lsb-release
    set +v

    assert_consent "Add Züs as a trusted package signer?" ${global_consent}
    set -v
    mkdir -p /etc/apt/keyrings
    curl -sLS https://packages.zus.network/zus.asc |
      gpg --dearmor > /etc/apt/keyrings/zus.gpg
    chmod go+r /etc/apt/keyrings/zus.gpg
    set +v

    assert_consent "Add the zbox Repository to your apt sources?" ${global_consent}
    set -v
    # Use env var DIST_CODE for the package dist name if provided
    if [[ -z $DIST_CODE ]]; then
        CLI_REPO=$(lsb_release -cs)
        shopt -s nocasematch
        ERROR_MSG="Unable to find a package for your system. Please check if an existing package in https://packages.zus.network/aptrepo/dists/ can be used in your system and install with the dist name: 'curl -sL https://packages.zus.network/aptrepo/ | sudo DIST_CODE=<dist_code_name> bash'"
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

    assert_consent "Install the zbox?" ${global_consent}
    apt-get install --assume-yes zbox

    # Check if config.yaml already exists
    CONFIG_DIR="${HOME}/.zcn"
    CONFIG_PATH="${CONFIG_DIR}/config.yaml"
    if [ -f "${CONFIG_PATH}" ]; then
        assert_consent "The configuration file ${CONFIG_PATH} already exists. Do you want to update it?" ${global_consent}
    fi

    # Create or update the config.yaml file
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