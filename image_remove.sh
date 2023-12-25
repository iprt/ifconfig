#!/bin/bash

function rmi() {
  if [ "all" == "$2" ]; then
    # 删除所有镜像
    # shellcheck disable=SC2046
    # docker image rm -f $(docker image ls "$1" |grep -v "REPOSITORY" |awk '{print $3}')
    docker image rm -f $(docker image ls|grep -v "REPOSITORY"|grep "$1"|awk '{print $3}')
  elif [ "none" == "$2" ]; then
    # 删除 <none> 的镜像
    # shellcheck disable=SC2046
    # docker image rm -f $(docker image ls "$1" |grep -v "REPOSITORY" |grep "<none>" |awk '{print $3}')
    docker image rm -f $(docker image ls|grep -v "REPOSITORY"|grep "<none>"|grep "$1"|awk '{print $3}')
  else
    # 保留latest镜像
    # shellcheck disable=SC2046
    # docker image rm -f $(docker image ls "$1" |grep -v "REPOSITORY" |grep -v "latest" |awk '{print $3}')
    docker image rm -f $(docker image ls|grep -v "REPOSITORY"|grep -v "latest" |grep "$1"|awk '{print $3}')
  fi
}

IMAGE=iproute/ifconfig


if [ -z "$1" ]; then
  rmi $IMAGE all
else
  rmi $IMAGE "$1"
fi
