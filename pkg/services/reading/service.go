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
	"micromango/pkg/grpc/activity"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/static"
)

type service struct {
	pb.UnimplementedReadingServer
	db       *gorm.DB
	static   static.StaticClient
	activity activity.ActivityClient
}

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)
	conn := utils.GrpcDialOrFatal(c.StaticServiceAddr)
	staticService := static.NewStaticClient(conn)

	conn = utils.GrpcDialOrFatal(c.ActivityServiceAddr)
	activityService := activity.NewActivityClient(conn)

	serv := service{
		db:       database,
		static:   staticService,
		activity: activityService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterReadingServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

func (s *service) GetMangaContent(ctx context.Context, req *pb.MangaContentRequest) (*pb.MangaContentResponse, error) {
	m, err := getMangaContent(s.db, req.MangaId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("content for %s not found", req.MangaId))
		}
		return nil, err
	}

	readSet := make(map[string]struct{})
	if req.UserId != nil {
		resp, err := s.activity.ReadChapters(ctx, &activity.ReadChaptersRequest{
			UserId:  *req.UserId,
			MangaId: req.MangaId,
		})
		if err != nil {
			return nil, err
		}
		for _, uid := range resp.ChapterIds {
			readSet[uid] = struct{}{}
		}
	}

	var resp pb.MangaContentResponse
	resp.Chapters = utils.Map(m, func(c ChapterHead) *pb.MangaContentResponse_ChapterHead {
		_, read := readSet[c.ChapterId]
		return &pb.MangaContentResponse_ChapterHead{
			ChapterId: c.ChapterId,
			Number:    c.Number,
			Title:     c.Title,
			Pages:     c.Pages,
			Read:      read,
			CreatedAt: c.CreatedAt.String(),
		}
	})
	resp.MangaId = req.MangaId
	return &resp, nil
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
	if req.Image != nil {
		res, err := s.static.UploadPage(ctx, &static.UploadPageRequest{
			MangaId:   req.MangaId,
			ChapterId: req.ChapterId,
			Page:      req.Image,
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
