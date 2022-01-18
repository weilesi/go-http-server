FROM golang:1.17 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY httpserver .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 8099
ENTRYPOINT ["/httpserver"]