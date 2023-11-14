FROM alpine:latest
MAINTAINER "mimotronik@gmail.com"

LABEL email="mimotronik@gamil.com" \
      author="zhuzhenjie"

WORKDIR /opt/app

ADD ifconfig ifconfig
ADD static/ static/
ADD views/ views/

EXPOSE 8080

ENTRYPOINT ["./ifconfig"]