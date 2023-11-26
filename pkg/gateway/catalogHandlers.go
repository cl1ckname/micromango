package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"io"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/catalog"
	"mime/multipart"
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
	addMangaReq.Title = ctx.FormValue("title")
	addMangaReq.Description = utils.Ptr(ctx.FormValue("description"))
	formFile, err := ctx.FormFile("cover")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	imageBytes, err := readFormFile(formFile)
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

func readFormFile(formFile *multipart.FileHeader) ([]byte, error) {
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
