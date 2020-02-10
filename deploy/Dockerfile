FROM golang:1.13.5-alpine  as builder
WORKDIR /go/src
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
COPY . /go/src

RUN echo "https://mirrors.aliyun.com/alpine/v3.10/main" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.10/community" >> /etc/apk/repositories \
    && apk update \
    && apk add git \
    && cd /go/src \
    && sh ./deploy/build.sh

# This dockerfile uses the apline:3.10 image
FROM alpine:3.10

MAINTAINER cugriver@163.com

WORKDIR /app

RUN echo "https://mirrors.aliyun.com/alpine/v3.10/main" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.10/community" >> /etc/apk/repositories \
    && apk update \
    && mkdir -p /app/log \
    && mkdir -p /app/etc

#拷贝配置文件
COPY --from=builder /go/src/etc/ums.xml /app/etc/
COPY --from=builder /go/src/bin/ums /app/

#暴露的端口号
EXPOSE 13172

ENTRYPOINT /app/ums