FROM golang:1.21.0 as builder
LABEL authors="clickname"

WORKDIR app

COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=1 GOOS=linux go build -o /service -a -ldflags '-linkmode external -extldflags "-static"' cmd/profile/main.go


FROM scratch
WORKDIR app
COPY --from=builder /service /service
CMD ["/service"]