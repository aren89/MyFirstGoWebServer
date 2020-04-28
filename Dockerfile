FROM golang:1.13-alpine as builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/server

FROM alpine:latest

RUN apk update && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]
