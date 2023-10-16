package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/grpc/catalog"
	"net/http"
)

func (s *server) GetManga(ctx echo.Context) error {
	var getMangaReq catalog.MangaRequest
	getMangaReq.MangaId = ctx.Param("mangaId")
	resp, err := s.catalog.GetManga(context.TODO(), &getMangaReq)
	if err != nil {
		ctx.Error(err)
		return nil
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddManga(ctx echo.Context) error {
	var addMangaReq catalog.AddMangaRequest
	if err := ctx.Bind(&addMangaReq); err != nil {
		ctx.Error(err)
		return nil
	}
	resp, err := s.catalog.AddManga(context.TODO(), &addMangaReq)
	if err != nil {
		ctx.Error(err)
		return nil
	}
	return ctx.JSON(http.StatusOK, resp)
}
