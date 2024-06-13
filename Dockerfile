# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /discord_bot

EXPOSE 8080

CMD ["/discord_bot"]