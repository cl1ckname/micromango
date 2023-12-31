package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/reading"
	"net/http"
	"strconv"
)

type readingHandler struct {
	reading reading.ReadingClient
}

func RegisterReading(g *echo.Group, r reading.ReadingClient) {
	h := readingHandler{r}
	readingGroup := g.Group("/content")

	readingGroup.GET("/:mangaId", h.GetMangaContent)
	readingGroup.GET("/:mangaId/chapter/:chapterId", h.GetChapter)
	readingGroup.PUT("/:mangaId/chapter/:chapterId", h.UpdateChapter)
	readingGroup.POST("/:mangaId/chapter", h.AddChapter)
	readingGroup.GET("/:mangaId/chapter/:chapterId/page/:pageId", h.GetPage)
	readingGroup.POST("/:mangaId/chapter/:chapterId/page", h.AddPage)

}

func (s *readingHandler) GetMangaContent(ctx echo.Context) error {
	var getMangaContentReq reading.MangaContentRequest
	getMangaContentReq.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&getMangaContentReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp, err := s.reading.GetMangaContent(context.TODO(), &getMangaContentReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *readingHandler) GetChapter(ctx echo.Context) error {
	var getChapterReq reading.ChapterRequest
	getChapterReq.ChapterId = ctx.Param("chapterId")
	if err := ctx.Bind(&getChapterReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp, err := s.reading.GetChapter(context.TODO(), &getChapterReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *readingHandler) AddChapter(ctx echo.Context) error {
	var addChapterContentRequest reading.AddChapterRequest
	addChapterContentRequest.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&addChapterContentRequest); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp, err := s.reading.AddChapter(context.TODO(), &addChapterContentRequest)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *readingHandler) UpdateChapter(ctx echo.Context) error {
	var updateChapterReq reading.UpdateChapterRequest
	if err := ctx.Bind(&updateChapterReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp, err := s.reading.UpdateChapter(context.TODO(), &updateChapterReq)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *readingHandler) GetPage(ctx echo.Context) error {
	var addMangaContentRequest reading.PageRequest
	addMangaContentRequest.PageId = ctx.Param("pageId")
	if err := ctx.Bind(&addMangaContentRequest); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	resp, err := s.reading.GetPage(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *readingHandler) AddPage(ctx echo.Context) error {
	var addMangaContentRequest reading.AddPageRequest
	addMangaContentRequest.MangaId = ctx.FormValue("mangaId")
	chapterNumberStr := ctx.FormValue("number")
	chapterNumber, err := strconv.ParseUint(chapterNumberStr, 10, 32)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	addMangaContentRequest.Number = uint32(chapterNumber)
	file, err := ctx.FormFile("image")
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	image, err := utils.ReadFormFile(file)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	addMangaContentRequest.Image = image
	addMangaContentRequest.ChapterId = ctx.Param("chapterId")

	resp, err := s.reading.AddPage(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return ctx.JSON(http.StatusOK, resp)
}
