FROM alpine:3.16 AS runner
RUN addgroup -g 5001 -S api
RUN adduser -h /apt/api -s /sbin/nologin -G api -S -u 5001 api
WORKDIR /apt/api
USER api:api
CMD /build/api "api:5000" /configs/api.env
