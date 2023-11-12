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
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	if resp == nil {
		return ctx.JSON(http.StatusOK, make([]catalog.MangaResponse, 0))
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) GetMangas(ctx echo.Context) error {
	mangas, err := s.catalog.GetMangas(context.TODO(), &catalog.Empty{})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, mangas.Mangas)
}

func (s *server) AddManga(ctx echo.Context) error {
	var addMangaReq catalog.AddMangaRequest
	if err := ctx.Bind(&addMangaReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.catalog.AddManga(context.TODO(), &addMangaReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}
