FROM golang:1.21.0 as builder
LABEL authors="clickname"

WORKDIR app

COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg
COPY static ./static

RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=1 GOOS=linux go build -o /service -a -ldflags '-linkmode external -extldflags "-static"' cmd/user/main.go


FROM scratch
WORKDIR app
COPY --from=builder /service /service
CMD ["/service"]