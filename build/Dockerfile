# Dockerfile for development

FROM golang:1.19.1-alpine3.15
RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app