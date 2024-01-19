#!/bin/bash
# shellcheck disable=SC2046

bash <(curl -SL https://gitlab.com/iprt/build-project/-/raw/main/docker/rmi.sh) \
-i "iproute/ifconfig" \
-s "all"
