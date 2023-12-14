package utils

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

func ReadFormFile(formFile *multipart.FileHeader) ([]byte, error) {
	file, err := formFile.Open()
	if err != nil {
		return nil, err
	}
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return imageBytes, nil
}

func ErrorToResponse(ctx echo.Context, err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	if st.Code() == codes.NotFound {
		return ctx.JSON(http.StatusNotFound, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
}

func ParseQueryIntArray[T int | uint32](s string) []T {
	if s == "" {
		return []T{}
	}
	parts := strings.Split(s, ",")
	return Map(parts, func(p string) T {
		i, _ := strconv.ParseInt(p, 10, 32)
		return T(i)
	})
}

func JsonMessage(ctx echo.Context, status int, message string) error {
	return ctx.JSON(status, struct {
		Message string `json:"message"`
	}{message})
}
