# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-short-url

EXPOSE 8080

CMD ["/docker-short-url"]
