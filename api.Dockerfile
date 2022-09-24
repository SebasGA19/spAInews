# Builder

FROM golang:1.18.5-alpine3.16 AS builder
RUN apk add build-base
# Prepare output directory
WORKDIR /build
# Copy code
WORKDIR /spAInews
COPY go.mod .
COPY go.sum .
COPY api api
# Compile
WORKDIR /spAInews/api
RUN go build -ldflags="-s -w" -trimpath -buildvcs=false -o /build ./cmd/...

# Runner

FROM alpine:3.16 AS runner
# Setup user
RUN addgroup -g 5001 -S api
RUN adduser -h /apt/api -s /sbin/nologin -G api -S -u 5001 api
# Copy compiled binary
WORKDIR /apt/api
COPY --from=builder /build/api /apt/api/server
# Execute
USER api:api
CMD /apt/api "api:5000" /configs/api.env
