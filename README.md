# ifconfig

## reference

https://ipcrystal.com

https://nginx.co

https://kubectl.net

## image

### build image

```shell
bash image_build.sh [IMAGE] [VERSION]
```

- IMAGE：可选，默认镜像的名称 `iproute/ifconfig`
- VERSION: 可选, 默认镜像版本

### push image

```shell
bash image_push.sh <REGISTRY> [IMAGE] [VERSION]
```

- REGISTRY: 必选，镜像仓库，一般是私仓
- IMAGE：可选，默认镜像的名称 `iproute/ifconfig`
- VERSION: 可选, 默认镜像版本 `latest`

## ip database

GeoLite2-City.mmdb

https://github.com/wp-statistics/GeoLite2-City

```text
https://cdn.jsdelivr.net/npm/geolite2-city@1.0.0/GeoLite2-City.mmdb.gz
```

ipipfree.ipdb

```text
https://www.ipip.net/product/ip.html
```

qqwry.dat

```text
https://github.com/FW27623/qqwry
```

ip2region.db

```text
https://github.com/lionsoul2014/ip2region
```

```text
https://gitee.com/lionsoul/ip2region
```