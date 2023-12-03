package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/user"
	"net/http"
)

func (s *server) GetManga(ctx echo.Context) error {
	var getMangaReq catalog.MangaRequest
	if claims, ok := ctx.Get("claims").(*user.UserResponse); ok {
		getMangaReq.UserId = &claims.UserId
	}

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
		if err != http.ErrMissingFile {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
	}
	if formFile != nil {
		imageBytes, err := utils.ReadFormFile(formFile)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
		addMangaReq.Cover = imageBytes
	}
	resp, err := s.catalog.AddManga(context.TODO(), &addMangaReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) UpdateManga(ctx echo.Context) error {
	var updateMangaReq catalog.UpdateMangaRequest
	if err := ctx.Bind(&updateMangaReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	updateMangaReq.MangaId = ctx.Param("mangaId")
	if description := ctx.FormValue("description"); description != "" {
		updateMangaReq.Description = utils.Ptr(description)
	}
	if title := ctx.FormValue("title"); title != "" {
		updateMangaReq.Title = utils.Ptr(title)
	}
	res, err := s.catalog.UpdateManga(context.TODO(), &updateMangaReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *server) DeleteManga(ctx echo.Context) error {
	mangaId := ctx.Param("mangaId")
	if _, err := s.catalog.DeleteManga(context.TODO(), &catalog.DeleteMangaRequest{MangaId: mangaId}); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.String(http.StatusOK, mangaId+" deleted")
}
