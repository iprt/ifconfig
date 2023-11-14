#!/bin/bash

REGISTRY=$1
IMAGE=$2
VERSION=$3

if [ -z "$REGISTRY" ]; then
    echo "registry is empty"
fi

if [ -z "$IMAGE" ]; then
    echo "default image is iproute/ifconfig"
    IMAGE=iproute/ifconfig
fi

if [ -z "$VERSION" ]; then
    echo "default version latest"
    VERSION=latest
fi

CON=$(docker image ls $IMAGE:$VERSION | wc -l)

if [ "$CON" -eq 2 ]; then
  echo tag new version
  docker tag $IMAGE:$VERSION "$REGISTRY"/$IMAGE:$VERSION:$VERSION
  echo "docker push $REGISTRY/$IMAGE:$VERSION:$VERSION"

  docker push "$REGISTRY/$IMAGE:$VERSION"
else
  echo "image $IMAGE:$VERSION does not exist"
fi