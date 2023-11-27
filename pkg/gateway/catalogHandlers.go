package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/catalog"
	"net/http"
)

func (s *server) GetManga(ctx echo.Context) error {
	var getMangaReq catalog.MangaRequest
	getMangaReq.MangaId = ctx.Param("mangaId")
	resp, err := s.catalog.GetManga(context.TODO(), &getMangaReq)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
		if st.Code() == codes.NotFound {
			return ctx.JSON(http.StatusNotFound, struct{ Message string }{err.Error()})
		}
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
	addMangaReq.Title = ctx.FormValue("title")
	addMangaReq.Description = utils.Ptr(ctx.FormValue("description"))
	formFile, err := ctx.FormFile("cover")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	imageBytes, err := utils.ReadFormFile(formFile)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	addMangaReq.Cover = imageBytes
	resp, err := s.catalog.AddManga(context.TODO(), &addMangaReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}
