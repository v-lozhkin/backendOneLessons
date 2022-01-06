FROM golang:1.17.5 as builder

WORKDIR /build
COPY . .
RUN go build -o /build/bin/api ./cmd/api

FROM centos:7

COPY --from=builder /build/bin/api /usr/local/bin/api