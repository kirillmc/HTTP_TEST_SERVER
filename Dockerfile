FROM golang:1.21.9-alpine3.18 AS builder

COPY . /github.com/kirillmc/HTTP_TEST_SERVER/source/
WORKDIR /github.com/kirillmc/HTTP_TEST_SERVER/source/

RUN go mod download
RUN go build -o ./bin/HTTP_TEST_SERVER cmd/http_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kirillmc/HTTP_TEST_SERVER/source/bin/HTTP_TEST_SERVER .
COPY .env .
CMD ["./HTTP_TEST_SERVER"]



