package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"micromango/pkg/grpc/static"
	"net/http"
)

func (s *server) GetStatic(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.String(http.StatusBadRequest, "invalid file id")
	}
	resp, err := s.static.GetImage(context.TODO(), &static.GetImageRequest{ImageId: id})
	if err != nil {
		log.Error(err)
		ctx.Error(err)
		return nil
	}
	if _, err := ctx.Response().Write(resp.Image); err != nil {
		log.Error(err)
		ctx.Error(err)
		return nil
	}
	ctx.Response().Header().Set("Content-Type", "image/jpg")
	return nil
}
