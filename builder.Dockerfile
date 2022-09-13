FROM golang:1.18.5-alpine3.16 AS builder
RUN apk add build-base
WORKDIR /spAInews
COPY go.mod .
COPY go.sum .
WORKDIR /spAInews/api
CMD go build -ldflags="-s -w" -trimpath -buildvcs=false -o /build ./cmd/...

