FROM golang:1.13

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o main ./src

EXPOSE 8080

CMD ["/app/main"]
