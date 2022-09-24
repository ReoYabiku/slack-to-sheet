# Dockerfile for deploy

FROM golang:1.19.1-alpine3.15
RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
ENV PORT 80
CMD [ "go",  "build", "./cmd/main.go" ]