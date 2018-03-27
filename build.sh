#!/bin/bash
# Adapted from: https://github.com/cyberark/summon/blob/master/build.sh

# Platforms to build: https://golang.org/doc/install/source#environment
PLATFORMS=(
  'darwin:amd64'
  'linux:amd64'
  'windows:amd64'
)
OUTPUT_DIR='output'

echo "building big-earl binaries in $OUTPUT_DIR/"
docker-compose build

for platform in "${PLATFORMS[@]}"; do
  GOOS=${platform%%:*}
  GOARCH=${platform#*:}

  echo "-----"
  echo "GOOS=$GOOS, GOARCH=$GOARCH"
  echo "....."

  docker-compose run --rm \
    -e GOOS=$GOOS -e GOARCH=$GOARCH \
    big-earl-builder \
    build -v -o $OUTPUT_DIR/big-earl-$GOOS-$GOARCH
done