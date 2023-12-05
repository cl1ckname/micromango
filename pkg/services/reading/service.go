package reading

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/static"
)

type service struct {
	pb.UnimplementedReadingServer
	db     *gorm.DB
	static static.StaticClient
}

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)
	conn := utils.GrpcDialOrFatal(c.StaticServiceAddr)
	staticService := static.NewStaticClient(conn)

	serv := service{
		db:     database,
		static: staticService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterReadingServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

func (s *service) GetMangaContent(_ context.Context, req *pb.MangaContentRequest) (*pb.MangaContentResponse, error) {
	m, err := getMangaContent(s.db, req.MangaId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("content for %s not found", req.MangaId))
		}
		return nil, err
	}
	resp := chaptersToPb(m)
	return resp, nil
}

func chaptersToPb(m []Chapter) *pb.MangaContentResponse {
	chapters := make([]*pb.MangaContentResponse_ChapterHead, len(m))
	for i, c := range m {
		chapters[i] = &pb.MangaContentResponse_ChapterHead{
			ChapterId: c.ChapterId.String(),
			Number:    c.Number,
			Title:     c.Title,
			CreatedAt: c.CreatedAt.String(),
		}
	}
	return &pb.MangaContentResponse{
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

func (s *service) UpdateChapter(_ context.Context, req *pb.UpdateChapterRequest) (*pb.ChapterResponse, error) {
	c, err := updateChapter(s.db, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("chapter %s not found", req.ChapterId))
		}
		return nil, err
	}
	return chapterToPb(c), nil
}

func chapterToPb(c Chapter) *pb.ChapterResponse {
	pages := make([]*pb.ChapterResponse_PageHead, len(c.Pages))
	for i, p := range c.Pages {
		pages[i] = &pb.ChapterResponse_PageHead{
			PageId: p.PageId.String(),
			Number: p.Number,
			Image:  p.Image,
		}
	}
	return &pb.ChapterResponse{
		ChapterId: c.ChapterId.String(),
		MangaId:   c.MangaId.String(),
		Number:    c.Number,
		Title:     c.Title,
		Pages:     pages,
		CreatedAt: c.CreatedAt.String(),
	}
}

func (s *service) GetPage(_ context.Context, req *pb.PageRequest) (*pb.PageResponse, error) {
	p, err := getPage(s.db, req.PageId)
	if err != nil {
		return nil, err
	}
	return pageToPB(p), nil
}

func (s *service) AddPage(ctx context.Context, req *pb.AddPageRequest) (*pb.PageResponse, error) {
	var imageUrl string
	if len(req.Image) > 0 {
		res, err := s.static.UploadPage(ctx, &static.UploadPageRequest{
			MangaId:   req.MangaId,
			ChapterId: req.ChapterId,
			Image:     req.Image,
			Type:      0, // FIXME
		})
		if err != nil {
			return nil, err
		}
		imageUrl = res.ImageId
	}

	chapterId, err := uuid.Parse(req.ChapterId)
	if err != nil {
		return nil, err
	}
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	p, err := addPage(s.db, Page{
		PageId:    uuid.New(),
		MangaId:   mangaId,
		ChapterId: chapterId,
		Number:    req.Number,
		Image:     imageUrl,
	})
	if err != nil {
		return nil, err
	}
	return pageToPB(p), nil
}

func pageToPB(p Page) *pb.PageResponse {
	return &pb.PageResponse{
		PageId:    p.PageId.String(),
		ChapterId: p.ChapterId.String(),
		Number:    p.Number,
		Image:     p.Image,
	}
}
