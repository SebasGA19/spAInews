# Build Auth API
FROM golang:1.18.5-alpine3.16 AS builder
RUN apk add build-base
WORKDIR /build
WORKDIR /spAInews
COPY . .
WORKDIR /spAInews/api/cmd/api
RUN go build -ldflags="-s -w" -trimpath -buildvcs=false -o /build/server

# Prepare image
FROM alpine:3.16 AS runner
RUN addgroup -g 5001 -S api
RUN adduser -h /apt/api -s /sbin/nologin -G api -S -u 5001 api
WORKDIR /apt/api
COPY --from=builder /build/server .
RUN chmod 777 /apt/api/server
USER api:api
EXPOSE 5000
CMD /apt/api/server

