# ifconfig

## 参考网站

https://ipcrystal.com

## 镜像相关

### 创建本地镜像

```shell
bash image_build.sh [IMAGE] [VERSION]
```

- IMAGE：可选，默认镜像的名称 `iproute/ifconfig`
- VERSION: 可选, 默认镜像版本

### 推送镜像

```shell
bash image_push.sh <REGISTRY> [IMAGE] [VERSION]
```

- REGISTRY: 必选，镜像仓库，一般是私仓
- IMAGE：可选，默认镜像的名称 `iproute/ifconfig`
- VERSION: 可选, 默认镜像版本

### 删除本地镜像

```shell
# 保留latest,删除其他所有
bash image_remove.sh

# 删除所有版本
bash image_remove.sh all

# 删除<none>的镜像版本
bash image_remove.sh none
```

## 需要下载的ip数据库

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