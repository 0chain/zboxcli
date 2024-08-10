#!/bin/bash

set -e

# apt-get install -y clang cmake git patch python-is-python3 libssl-dev lzma-dev libxml2-dev xz-utils bzip2 cpio libbz2-dev zlib1g-dev libc++-dev libc++abi-dev build-essential
# dpkg -L libc++-dev
# dpkg -L libc++abi-dev

export OSXCROSS_ROOT=$HOME/osxcross
export PATH=$OSXCROSS_ROOT/target/bin:$PATH
export SDKROOT=$OSXCROSS_ROOT/target/SDK/MacOSX11.3.sdk
# export CFLAGS="-isysroot $SDKROOT -stdlib=libc++"
# export CXXFLAGS="-isysroot $SDKROOT -stdlib=libc++"
# export LDFLAGS="-isysroot $SDKROOT -stdlib=libc++"
# Set up Go environment
export GO111MODULE=on
export CGO_ENABLED=1

VERSION=1.16.0

# Cross-compile for Darwin/amd64
echo "Building for Darwin/amd64..."
CC=o64-clang CXX=o64-clang++ GOOS=darwin GOARCH=amd64 ~/sdk/go1.21.8/bin/go build -x -v -tags bn256 -ldflags "-X main.VersionStr=${VERSION}" -o zboxcli-darwin-amd64 .

# Cross-compile for Darwin/arm64
echo "Building for Darwin/arm64..."
CC=o64-clang CXX=o64-clang++ GOOS=darwin GOARCH=arm64 ~/sdk/go1.21.8/bin/go build -x -v -tags bn256 -ldflags "-X main.VersionStr=${VERSION}" -o zboxcli-darwin-arm64 .

echo "Build completed."
