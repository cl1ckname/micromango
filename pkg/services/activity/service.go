package activity

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/activity"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/share"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	conn := utils.GrpcDialOrFatal(c.CatalogAddr)
	catalogService := catalog.NewCatalogClient(conn)

	serv := service{
		db:      database,
		catalog: catalogService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterActivityServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	db *gorm.DB
	pb.UnimplementedActivityServer
	catalog catalog.CatalogClient
}

func (s *service) Like(ctx context.Context, req *pb.LikeRequest) (*share.Empty, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	if _, err := SaveLike(s.db, mangaId, userId); err != nil {
		return nil, err
	}
	likes, err := LikesCount(s.db, mangaId)
	if err != nil {
		return nil, err
	}
	if _, err := s.catalog.SetLikes(ctx, &catalog.SetLikesRequest{
		MangaId: req.MangaId,
		Likes:   likes,
	}); err != nil {
		return nil, err
	}
	return &share.Empty{}, nil
}

func (s *service) Dislike(_ context.Context, req *pb.DislikeRequest) (*share.Empty, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	if err := RemoveLike(s.db, mangaId, userId); err != nil {
		return nil, err
	}
	return &share.Empty{}, nil
}

func (s *service) HasLike(_ context.Context, req *pb.HasLikeRequest) (*pb.HasLikeResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	has, err := HasLike(s.db, userId, mangaId)
	if err != nil {
		return nil, err
	}
	return &pb.HasLikeResponse{Has: has}, nil
}

func (s *service) RateManga(ctx context.Context, req *pb.RateMangaRequest) (*share.Empty, error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	err = SaveRate(s.db, userId, mangaId, req.Rate)
	if err != nil {
		return nil, err
	}
	avg, err := AvgRate(s.db, mangaId)
	if err != nil {
		return nil, err
	}
	if _, err := s.catalog.SetAvgRate(ctx, &catalog.SetAvgRateRateRequest{
		MangaId: req.MangaId,
		Rate:    avg.Rate,
		Rates:   avg.Voters,
	}); err != nil {
		return nil, err
	}

	return &share.Empty{}, err
}

func (s *service) UserRate(_ context.Context, req *pb.UserRateRequest) (*pb.UserRateResponse, error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	rate, err := UserRate(s.db, userId, mangaId)
	if err != nil {
		return nil, err
	}
	return &pb.UserRateResponse{Rate: rate}, nil
}

func (s *service) UserRateList(_ context.Context, req *pb.UserRateListRequest) (*pb.UserRateListResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	rates, err := RateList(s.db, userId, req.MangaId)
	if err != nil {
		return nil, err
	}
	resp := pb.UserRateListResponse{Rates: make(map[string]uint32, len(rates))}
	for _, m := range rates {
		resp.Rates[m.MangaId.String()] = m.Rate
	}
	return &resp, nil
}

func (s *service) ReadChapter(_ context.Context, req *pb.ReadChapterRequest) (*share.Empty, error) {
	chapterId, err := uuid.Parse(req.ChapterId)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	if err := ReadChapter(s.db, userId, chapterId); err != nil {
		return nil, err
	}
	return &share.Empty{}, nil
}
