#!/bin/bash

set -e

SRC_DIR=$PWD
HOME_DIR=$HOME

echo "PWD=$SRC_DIR , HOME=$HOME_DIR"

# sudo apt-get update
sudo apt-get install -y clang cmake git patch python libssl-dev lzma-dev libxml2-dev xz-utils bzip2 cpio libbz2-dev zlib1g-dev

cd $HOME_DIR
rm -rf $HOME_DIR/osxcross
git clone https://github.com/tpoechtrager/osxcross.git
cd osxcross/tarballs
wget https://github.com/phracker/MacOSX-SDKs/releases/download/11.3/MacOSX11.3.sdk.tar.xz
cd ..
export CC=clang
export CXX=clang++
UNATTENDED=yes ./build.sh

export OSXCROSS_ROOT=$PWD
export PATH=$OSXCROSS_ROOT/target/bin:$PATH
export GO111MODULE=on
export CGO_ENABLED=1

# Cross-compile for Darwin/amd64
echo "Building for Darwin/amd64..."
CC=o64-clang CXX=o64-clang++ GOOS=darwin GOARCH=amd64 go build -x -v -tags bn256 -ldflags "-X main.VersionStr=${VERSION} -linkmode 'external' -extldflags '-static'" -o zboxcli-darwin-amd64 .

# Cross-compile for Darwin/arm64
echo "Building for Darwin/arm64..."
CC=o64-clang CXX=o64-clang++ GOOS=darwin GOARCH=arm64 go build -x -v -tags bn256 -ldflags "-X main.VersionStr=${VERSION} -linkmode 'external' -extldflags '-static'" -o zboxcli-darwin-arm64 .

echo "Build completed."

echo "osxcross setup success!"