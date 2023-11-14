FROM golang:1.20.0 As buildStage
ARG GOPROXY=https://goproxy.cn,direct
WORKDIR /opt/app
ADD . /opt/app
RUN go build

FROM alpine:latest
MAINTAINER "mimotronik@gmail.com"

LABEL email="mimotronik@gamil.com" \
      author="zhuzhenjie"

WORKDIR /opt/app

COPY --from=buildStage /opt/app/ifconfig /opt/app/

ADD conf/ conf/
ADD static/ static/
ADD views/ views/

EXPOSE 8080

ENTRYPOINT ["./ifconfig"]