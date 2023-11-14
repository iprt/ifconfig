FROM debian:12-alpine

MAINTAINER "mimotronik@gmail.com"

LABEL email="mimotronik@gamil.com" \
      author="zhuzhenjie"

WORKDIR /opt/app

ADD ifconfig ifconfig
ADD conf/ conf/
ADD static/ static/
ADD views/ views/

EXPOSE 8080

ENTRYPOINT ["./ifconfig"]