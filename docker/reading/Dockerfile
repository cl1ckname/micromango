FROM golang:1.21.0 as builder
LABEL authors="clickname"

WORKDIR app

COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg

CMD go run cmd/reading/main.go