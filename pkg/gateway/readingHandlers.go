package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/grpc/reading"
	"net/http"
)

func (s *server) GetMangaContent(ctx echo.Context) error {
	var getMangaContentReq reading.MangaContentRequest
	getMangaContentReq.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&getMangaContentReq); err != nil {
		return err
	}
	resp, err := s.reading.GetMangaContent(context.TODO(), &getMangaContentReq)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddMangaContent(ctx echo.Context) error {
	var addMangaContentRequest reading.AddMangaContentRequest
	addMangaContentRequest.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&addMangaContentRequest); err != nil {
		return err
	}
	resp, err := s.reading.AddMangaContent(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) GetChapter(ctx echo.Context) error {
	var getChapterReq reading.ChapterRequest
	getChapterReq.ChapterId = ctx.Param("chapterId")
	if err := ctx.Bind(&getChapterReq); err != nil {
		return err
	}
	resp, err := s.reading.GetChapter(context.TODO(), &getChapterReq)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddChapter(ctx echo.Context) error {
	var addChapterContentRequest reading.AddChapterRequest
	addChapterContentRequest.MangaId = ctx.Param("mangaId")
	if err := ctx.Bind(&addChapterContentRequest); err != nil {
		return err
	}
	resp, err := s.reading.AddChapter(context.TODO(), &addChapterContentRequest)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) GetPage(ctx echo.Context) error {
	var addMangaContentRequest reading.PageRequest
	addMangaContentRequest.PageId = ctx.Param("pageId")
	if err := ctx.Bind(&addMangaContentRequest); err != nil {
		return err
	}
	resp, err := s.reading.GetPage(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *server) AddPage(ctx echo.Context) error {
	var addMangaContentRequest reading.AddPageRequest
	addMangaContentRequest.ChapterId = ctx.Param("chapterId")
	if err := ctx.Bind(&addMangaContentRequest); err != nil {
		return err
	}
	resp, err := s.reading.AddPage(context.TODO(), &addMangaContentRequest)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}
