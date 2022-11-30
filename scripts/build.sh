#!/bin/bash
#

BUILD_DIR=../build
export BIGBANG_VERSION=$(cat VERSION)

echo "Creating the deploy package"
if [ ! -d "$BUILD_DIR" ]; then
	error "Creating Build Dir"
	mkdir -p "$BUILD_DIR"
fi

zarf package create --skip-sbom --set BIGBANG_VERSION=${BIGBANG_VERSION} --confirm -l debug
mv zarf-package-big-bang-amd64.tar.zst $BUILD_DIR
