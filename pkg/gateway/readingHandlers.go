package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/reading"
	"net/http"
	"strconv"
)

func (s *server) GetMangaContent(ctx echo.Context) error {
	var getMangaContentReq reading.MangaContentRequest
	getMangaContentReq.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&getMangaContentReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.reading.GetMangaContent(context.TODO(), &getMangaContentReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddMangaContent(ctx echo.Context) error {
	var addMangaContentRequest reading.AddMangaContentRequest
	addMangaContentRequest.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&addMangaContentRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.reading.AddMangaContent(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) GetChapter(ctx echo.Context) error {
	var getChapterReq reading.ChapterRequest
	getChapterReq.ChapterId = ctx.Param("chapterId")
	if err := ctx.Bind(&getChapterReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.reading.GetChapter(context.TODO(), &getChapterReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddChapter(ctx echo.Context) error {
	var addChapterContentRequest reading.AddChapterRequest
	addChapterContentRequest.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&addChapterContentRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.reading.AddChapter(context.TODO(), &addChapterContentRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) GetPage(ctx echo.Context) error {
	var addMangaContentRequest reading.PageRequest
	addMangaContentRequest.PageId = ctx.Param("pageId")
	if err := ctx.Bind(&addMangaContentRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	resp, err := s.reading.GetPage(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddPage(ctx echo.Context) error {
	var addMangaContentRequest reading.AddPageRequest
	addMangaContentRequest.ChapterId = ctx.FormValue("chapterId")
	chapterNumberStr := ctx.FormValue("number")
	chapterNumber, err := strconv.ParseUint(chapterNumberStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	addMangaContentRequest.Number = uint32(chapterNumber)
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	fileBytes, err := utils.ReadFormFile(file)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	addMangaContentRequest.Image = fileBytes
	addMangaContentRequest.ChapterId = ctx.Param("chapterId")

	resp, err := s.reading.AddPage(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, struct{ Message string }{err.Error()})
	}
	return ctx.JSON(http.StatusOK, resp)
}
