#!/bin/bash

SHELL_FOLDER=$(
  # shellcheck disable=SC2164
  cd "$(dirname "$0")"
  pwd
)

# shellcheck disable=SC2164
cd "$SHELL_FOLDER"

# re tag
CON=$(docker image ls iproute/ifconfig:latest | wc -l)

if [ "$CON" -eq 2 ]; then
  echo tag new version
  TIMESTAMP_VERSION=$(date '+%Y-%m-%d_%H-%M-%S')
  docker tag iproute/ifconfig:latest iproute/ifconfig:"$TIMESTAMP_VERSION"
fi

docker build -f Dockerfile -t iproute/ifconfig .