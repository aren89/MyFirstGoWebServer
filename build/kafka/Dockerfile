FROM golang:1.13-alpine as builder

RUN apk add \
    && apk add librdkafka-dev pkgconf \
    && apk add gcc \
    && apk add libc-dev

WORKDIR /app

COPY . .

RUN go build -tags musl -o main ./cmd/kafka

FROM alpine:latest

RUN apk update && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]
