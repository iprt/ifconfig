#!/bin/bash

SHELL_FOLDER=$(
  # shellcheck disable=SC2164
  cd "$(dirname "$0")"
  pwd
)

# shellcheck disable=SC2164
cd "$SHELL_FOLDER"

echo "build ifconfig start ..."

echo "remove old pkg"

rm -f ifconfig

GO_IMAGE=golang:1.20
GO_IMAGE_CACHE=go_1.20_cache
OUT_FILE=ifconfig

echo "build ifconfig in docker container $GO_IMAGE"
docker run --rm -v "$PWD":/usr/src/myapp \
-w /usr/src/myapp \
-e CGO_ENABLED=0 \
-e GOPROXY=https://goproxy.cn,direct \
-e GOPATH=/opt/go \
-v "$GO_IMAGE_CACHE":/opt/go \
$GO_IMAGE \
go build -v -o $OUT_FILE

if [ -f "$OUT_FILE" ]; then
  echo  "build $OUT_FILE success"
else
  echo "build $OUT_FILE failed, then exit"
  exit
fi

echo "build ifconfig end"

IMAGE=$1
VERSION=$2

if [ -z "$IMAGE" ]; then
    echo "default image name is: iproute/ifconfig"
    IMAGE=iproute/ifconfig
fi

if [ -z "$VERSION" ]; then
    echo "default image version is: latest"
    VERSION=latest
fi

# re tag
CON=$(docker image ls $IMAGE:$VERSION | wc -l)

if [ "$CON" -eq 2 ]; then
  echo tag new version
  TIMESTAMP_VERSION=$(date '+%Y-%m-%d_%H-%M-%S')
  docker tag "$IMAGE:$VERSION" "$IMAGE:$TIMESTAMP_VERSION"
fi

docker build -f Dockerfile -t iproute/ifconfig .

echo "build ifconfig end ..."
