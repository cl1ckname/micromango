package grpc

import (
	"context"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading/entity"
	"micromango/pkg/services/reading/usecases"
)

type Server struct {
	readingCase usecases.Case
	pb.UnimplementedReadingServer
}

func NewServer(cs usecases.Case) *Server {
	return &Server{readingCase: cs}
}

func (s *Server) GetMangaContent(_ context.Context, req *pb.MangaContentRequest) (*pb.MangaContentResponse, error) {
	content, err := s.readingCase.GetMangaContent(req.MangaId, req.UserId)
	if err != nil {
		return nil, err
	}
	var response pb.MangaContentResponse
	response.Chapters = utils.Map(content, func(c entity.ChapterHead) *pb.MangaContentResponse_ChapterHead {
		return &pb.MangaContentResponse_ChapterHead{
			ChapterId: c.ChapterId,
			Number:    c.Number,
			Title:     c.Title,
			Pages:     c.Pages,
			Read:      c.Read,
			CreatedAt: c.CreatedAt.String(),
		}
	})
	return &response, nil
}
