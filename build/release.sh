#!/bin/sh
# Builds release archives for all supported platforms into build/dist.
set -e

if [ ! -f "build/release.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

VERSION=$(sed -n 's/^const version = "\(.*\)"$/\1/p' main.go)
DIST="$PWD/build/dist"
rm -rf "$DIST" && mkdir -p "$DIST"

build() {
    goos=$1; goarch=$2; name=$3; ext=$4
    dir=$(mktemp -d)
    bin=stratum-ping
    [ "$goos" = "windows" ] && bin=stratum-ping.exe
    CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -trimpath -ldflags "-s -w" -o "$dir/$bin"
    if [ "$ext" = "zip" ]; then
        (cd "$dir" && zip -q "$DIST/stratum-ping-$name.zip" "$bin")
    else
        tar -C "$dir" -czf "$DIST/stratum-ping-$name.tar.gz" "$bin"
    fi
    rm -rf "$dir"
}

build linux   amd64 linux-x64   tar.gz
build darwin  arm64 mac-arm64   zip
build darwin  amd64 mac-x64     zip
build windows amd64 win-x64     zip

(cd "$DIST" && sha256sum stratum-ping-* > SHA256SUMS)
echo "stratum-ping $VERSION:"; ls -l "$DIST"
