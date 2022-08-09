# Build Auth API
FROM golang:1.18.5-alpine3.16 AS builder
RUN apk add build-base
WORKDIR /build
WORKDIR /spAInews
COPY . .
WORKDIR /spAInews/auth-api/cmd/auth-api
RUN go build -ldflags="-s -w" -trimpath -buildvcs=false -o /build/server

# Prepare image
FROM alpine:3.16 AS runner
RUN addgroup -g 5001 -S auth-api
RUN adduser -h /apt/auth-api -s /sbin/nologin -G auth-api -S -u 5001 auth-api
WORKDIR /apt/auth-api
COPY --from=builder /build/server .
RUN chmod 777 /apt/auth-api/server
USER auth-api:auth-api
EXPOSE 5000
CMD /apt/auth-api/server

