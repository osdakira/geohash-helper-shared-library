#!/bin/bash

set -euxo pipefail

for os in darwin linux
do
    for arch in amd64 arm64
    do
        CGO_ENABLED=1 \
        GOOS=${os} \
        GOARCH=${arch} \
        go build \
            -buildmode=c-shared \
            -o lib/build/${os}_${arch}_geohash_lib_go.so \
            geohash_lib.go
    done
done
