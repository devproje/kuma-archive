#!/bin/bash
echo "building source..."
bun run build

echo "getting system info..."
VERSION="$(node -e "console.log(require('./package.json').version);")"
OS_INFO="$(echo "$(uname -s)" | awk '{print tolower($0)}')"
ARCH="$(uname -m)"

rm -rf dist/src/

echo "generated dist directory..."
mkdir -p dist/
mkdir -p dist/generated
mkdir -p dist/src/kuma-archive

echo "copy executable files..."
mv web/ dist/src/kuma-archive/
mv kuma-archive dist/src/kuma-archive/

echo "entering directory..."
cd dist/src/

echo "compressing build artifact..."
tar zcf ../generated/kuma-archive-${VERSION}-${OS_INFO}-${ARCH}.tar.gz kuma-archive/

echo "leave directory..."
cd ../../

echo "packaging complete!"
