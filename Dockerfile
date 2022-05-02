# syntax = docker/dockerfile:1.3-labs

FROM golang:1-alpine as builder
ARG VERSION=0.0.0
WORKDIR /go/src/trex_exporter
COPY . .
RUN apk --no-cache add git openssh build-base
RUN go build -ldflags "-X version=${VERSION}" -o app .

FROM alpine as production
LABEL maintainer="rluisr" \
  org.opencontainers.image.url="https://github.com/rluisr/trex_exporter" \
  org.opencontainers.image.source="https://github.com/rluisr/trex_exporter" \
  org.opencontainers.image.vendor="rluisr" \
  org.opencontainers.image.title="trex_exporter" \
  org.opencontainers.image.description="T-Rex prometheus exporter" \
  org.opencontainers.image.licenses="WTFPL"
RUN <<EOF
    apk add --no-cache ca-certificates libc6-compat tzdata \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
    rm -rf /var/cache/apk/*
EOF
ENV TZ="Asia/Tokyo"
COPY --from=builder /go/src/trex_exporter/app /app
ENTRYPOINT ["/app"]

