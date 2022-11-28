FROM alpine:3.7 AS builder
ENV GOPATH=/go
RUN apk add --no-cache go git