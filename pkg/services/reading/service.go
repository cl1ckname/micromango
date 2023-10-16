package reading

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/reading"
)

type service struct {
	pb.UnimplementedReadingServer
	db *gorm.DB
}

func Run(c Config) {
	database := Connect(c.DbAddr)
	serv := service{
		db: database,
	}
	addr := fmt.Sprintf(c.Addr)
	if err := common.RunGRPCServer(addr, func(registrar grpc.ServiceRegistrar) {
		pb.RegisterReadingServer(registrar, &serv)
	}); err != nil {
		panic(err)
	}
}

func (s *service) GetMangaContent(_ context.Context, req *pb.MangaContentRequest) (*pb.MangaContentResponse, error) {
	m, err := getMangaContent(s.db, req.MangaId)
	if err != nil {
		return nil, err
	}
	resp := mangaContentToPb(m)
	return resp, nil
}

func (s *service) AddMangaContent(context.Context, *pb.AddMangaContentRequest) (*pb.MangaContentResponse, error) {
	m, err := addMangaContent(s.db)
	return mangaContentToPb(m), err
}

func mangaContentToPb(m MangaContent) *pb.MangaContentResponse {
	chapters := make([]*pb.MangaContentResponse_ChapterHead, len(m.Chapters))
	for i, c := range m.Chapters {
		chapters[i] = &pb.MangaContentResponse_ChapterHead{
			ChapterId:     c.ChapterId.String(),
			ChapterNumber: c.ChapterNumber,
			Title:         c.Title,
		}
	}
	return &pb.MangaContentResponse{
		MangaId:  m.MangaId.String(),
		Chapters: chapters,
	}
}

func (s *service) GetChapter(_ context.Context, req *pb.ChapterRequest) (*pb.ChapterResponse, error) {
	c, err := getChapter(s.db, req.ChapterId)
	if err != nil {
		return nil, err
	}
	return chapterToPb(c), nil
}

func (s *service) AddChapter(_ context.Context, req *pb.AddChapterRequest) (*pb.ChapterResponse, error) {
	c, err := addChapter(s.db, req)
	if err != nil {
		return nil, err
	}
	return chapterToPb(c), nil
}

func chapterToPb(c Chapter) *pb.ChapterResponse {
	pages := make([]*pb.ChapterResponse_PageHead, len(c.Pages))
	for i, p := range c.Pages {
		pages[i] = &pb.ChapterResponse_PageHead{
			PageId:     p.PageId.String(),
			PageNumber: p.PageNumber,
		}
	}
	return &pb.ChapterResponse{
		ChapterId:     c.ChapterId.String(),
		MangaId:       c.MangaId.String(),
		ChapterNumber: c.ChapterNumber,
		Title:         c.Title,
		Pages:         pages,
	}
}

func (s *service) GetPage(_ context.Context, req *pb.PageRequest) (*pb.PageResponse, error) {
	p, err := getPage(s.db, req.PageId)
	if err != nil {
		return nil, err
	}
	return pageToPB(p), nil
}

func (s *service) AddPage(_ context.Context, req *pb.AddPageRequest) (*pb.PageResponse, error) {
	p, err := addPage(s.db, req)
	if err != nil {
		return nil, err
	}
	return pageToPB(p), nil
}

func pageToPB(p Page) *pb.PageResponse {
	return &pb.PageResponse{
		PageId:     p.PageId.String(),
		ChapterId:  p.ChapterId.String(),
		PageNumber: p.PageNumber,
		Image:      p.Image,
	}
}
