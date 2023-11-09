FROM debian:12

MAINTAINER "mimotronik@gmail.com"

LABEL email="mimotronik@gamil.com" \
      author="zhuzhenjie"

WORKDIR /opt/app

ADD ifconfig ifconfig
ADD conf/ conf/

EXPOSE 8080

ENTRYPOINT ["./ifconfig"]