package handlers

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/user"
	"net/http"
)

type catalogHandler struct {
	catalog catalog.CatalogClient
}

func RegisterCatalog(g *echo.Group, c catalog.CatalogClient) {
	h := catalogHandler{c}
	catalogGroup := g.Group("/catalog")

	catalogGroup.GET("", h.GetMangas)
	catalogGroup.POST("", h.AddManga)
	catalogGroup.GET("/:mangaId", h.GetManga)
	catalogGroup.PUT("/:mangaId", h.UpdateManga)
	catalogGroup.DELETE("/:mangaId", h.DeleteManga)

	feedGroup := g.Group("/feed")

	feedGroup.GET("/updates", h.GetUpdated)
}

func (s *catalogHandler) GetManga(ctx echo.Context) error {
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

func (s *catalogHandler) GetMangas(ctx echo.Context) error {
	var req catalog.GetMangasRequest
	req.GenresInclude = utils.ParseQueryIntArray[uint32](ctx.QueryParam("genre"))
	req.GenresExclude = utils.ParseQueryIntArray[uint32](ctx.QueryParam("exclude_genre"))
	if starts := ctx.QueryParam("starts"); starts != "" {
		req.Starts = utils.Ptr(starts)
	}
	if order := ctx.QueryParam("order"); order != "" {
		req.Order = utils.Ptr(order)
	}
	if asc := ctx.QueryParam("asc"); asc == "true" {
		req.Asc = true
	}
	mangas, err := s.catalog.GetMangas(context.TODO(), &req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, mangas.Mangas)
}

func (s *catalogHandler) GetUpdated(ctx echo.Context) error {
	mangas, err := s.catalog.LastUpdates(context.TODO(), &catalog.LastUpdatesRequest{
		Page:   1,
		Number: 10,
	})
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, mangas.Manga)
}

func (s *catalogHandler) AddManga(ctx echo.Context) error {
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "unauthorized")
	}
	var addMangaReq catalog.AddMangaRequest
	addMangaReq.UserId = auth.UserId
	addMangaReq.Title = ctx.FormValue("title")
	addMangaReq.Genres = utils.ParseQueryIntArray[uint32](ctx.FormValue("genres"))
	addMangaReq.Description = utils.Ptr(ctx.FormValue("description"))
	formFile, err := ctx.FormFile("cover")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
	}
	if formFile != nil {
		thumbnail, err := utils.ReadFormFile(formFile)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
		}
		addMangaReq.Thumbnail = thumbnail
	}
	resp, err := s.catalog.AddManga(context.TODO(), &addMangaReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *catalogHandler) UpdateManga(ctx echo.Context) error {
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "unauthorized")
	}
	var updateMangaReq catalog.UpdateMangaRequest
	updateMangaReq.UserId = auth.UserId
	if err := ctx.Bind(&updateMangaReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	updateMangaReq.MangaId = ctx.Param("mangaId")
	if description := ctx.FormValue("description"); description != "" {
		updateMangaReq.Description = utils.Ptr(description)
	}
	updateMangaReq.Genres = utils.ParseQueryIntArray[uint32](ctx.FormValue("genres"))
	if title := ctx.FormValue("title"); title != "" {
		updateMangaReq.Title = utils.Ptr(title)
	}

	thumbnail, err := utils.FormFile(ctx, "thumbnail")
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	updateMangaReq.Thumbnail = thumbnail

	cover, err := utils.FormFile(ctx, "cover")
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	updateMangaReq.Cover = cover

	res, err := s.catalog.UpdateManga(context.TODO(), &updateMangaReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *catalogHandler) DeleteManga(ctx echo.Context) error {
	mangaId := ctx.Param("mangaId")
	if _, err := s.catalog.DeleteManga(context.TODO(), &catalog.DeleteMangaRequest{MangaId: mangaId}); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.String(http.StatusOK, mangaId+" deleted")
}
