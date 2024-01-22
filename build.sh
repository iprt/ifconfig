#!/bin/bash
# shellcheck disable=SC2164 disable=SC1090
SHELL_FOLDER=$(
  cd "$(dirname "$0")"
  pwd
)
cd "$SHELL_FOLDER"

source <(curl -SL https://code.kubectl.net/devops/build-project/raw/branch/main/func/log.sh)

log "build" ">>> build ifconfig start <<<"

out_file=ifconfig

log "build" "remove old executable file: rm -rf $out_file"
rm -rf $out_file

log "step 1" "build executable file $out_file"

build_image=golang:1.20
build_image_cache=go_1.20_cache

go_build_url="https://code.kubectl.net/devops/build-project/raw/branch/main/golang/build.sh"

bash <(curl -SL $go_build_url) \
  -i "$build_image" \
  -c "$build_image_cache" \
  -x "go build -v -o $out_file"

docker_build_url="https://code.kubectl.net/devops/build-project/raw/branch/main/docker/build.sh"

registry="registry.cn-shanghai.aliyuncs.com"
image="iproute/ifconfig"

bash <(curl -SL $docker_build_url) \
  -i "$registry/$image" \
  -v "latest" \
  -r "false" \
  -p "true"

log "build" ">>> build ifconfig end <<<"
