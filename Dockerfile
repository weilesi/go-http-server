FROM golang:1.17 as builder

ENV GO1111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY httpserver/main.go .
COPY httpserver/professional/* httpserver/professional/
COPY httpserver/simple/* httpserver/simple/
COPY httpserver/utils/* httpserver/utils/
RUN ls -ltra
RUN go build -o httpserver/bin/amd64 .

FROM scratch
COPY --from=builder /build/httpserver/bin/amd64 /httpserver
EXPOSE 8099
ENTRYPOINT ["/httpserver"]