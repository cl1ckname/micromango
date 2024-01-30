package grpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	commonerr "micromango/pkg/common/errors"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading/entity"
	"micromango/pkg/services/reading/usecases"
)

type Server struct {
	readingCase usecases.Chapter
	pageCase    usecases.Page
	pb.UnimplementedReadingServer
}

func NewServer(cs usecases.Chapter) *Server {
	return &Server{readingCase: cs}
}

func (s *Server) GetMangaContent(_ context.Context, req *pb.MangaContentRequest) (*pb.MangaContentResponse, error) {
	content, err := s.readingCase.GetMangaContent(req.MangaId, req.UserId)
	if err != nil {
		return nil, err
	}
	var response pb.MangaContentResponse
	response.Chapters = utils.Map(content, ChapterToHead)
	return &response, nil
}

func (s *Server) GetChapter(_ context.Context, req *pb.ChapterRequest) (*pb.ChapterResponse, error) {
	chapter, err := s.readingCase.GetChapter(req.ChapterId)
	if err != nil {
		return nil, err
	}
	return ChapterToPb(chapter), nil
}

func (s *Server) AddChapter(_ context.Context, req *pb.AddChapterRequest) (*pb.ChapterResponse, error) {
	c := entity.Chapter{
		Title:   req.Title,
		MangaId: req.MangaId,
		Number:  req.Number,
	}
	saved, err := s.readingCase.AddChapter(c)
	if err != nil {
		return nil, err
	}
	return ChapterToPb(saved), nil
}

func (s *Server) UpdateChapter(_ context.Context, req *pb.UpdateChapterRequest) (*pb.ChapterResponse, error) {
	updateDto := entity.UpdateChapterDto{
		Title:  req.Title,
		Number: req.Number,
	}
	updatedChapter, err := s.readingCase.UpdateChapter(req.ChapterId, updateDto)
	if err != nil {
		if errors.Is(err, &commonerr.ErrNotFound{}) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("chapter %s not found", req.ChapterId))
		}
		return nil, err
	}
	return ChapterToPb(updatedChapter), nil
}

func (s *Server) AddPage(_ context.Context, req *pb.AddPageRequest) (*pb.PageResponse, error) {
	dto := entity.AddPageDto{
		MangaId:   req.MangaId,
		ChapterId: req.ChapterId,
		Number:    req.Number,
		Image:     utils.FileFromPb(req.Image),
	}
	page, err := s.pageCase.AddPage(dto)
	if err != nil {
		return nil, err
	}
	return PageToPb(page), nil
}

func (s *Server) GetPage(_ context.Context, req *pb.PageRequest) (*pb.PageResponse, error) {
	page, err := s.pageCase.GetPage(req.PageId)
	if err != nil {
		if errors.Is(err, &commonerr.ErrNotFound{}) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("page %s not found", req.PageId))
		}
		return nil, err
	}
	return PageToPb(page), nil
}
