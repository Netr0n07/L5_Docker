# Etap 1 - Budowanie aplikacji w Go
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY Simpleweb/main.go .
ARG VERSION=1.0
ENV VERSION=$VERSION

ENV CGO_ENABLED=0
ENV GO111MODULE=off

RUN go build -buildmode=exe -ldflags "-X main.version=${VERSION}" -o simpleweb

# Etap 2 - Finalny obraz scratch
FROM scratch

ADD https://dl-cdn.alpinelinux.org/alpine/v3.21/releases/x86_64/alpine-minirootfs-3.21.3-x86_64.tar.gz /

COPY --from=builder /app/simpleweb /simpleweb

EXPOSE 8080
CMD ["/simpleweb"]
