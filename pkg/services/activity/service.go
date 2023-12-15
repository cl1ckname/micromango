package activity

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/activity"
	"micromango/pkg/grpc/share"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	serv := service{
		db: database,
	}
	baseServer := grpc.NewServer()
	pb.RegisterActivityServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	db *gorm.DB
	pb.UnimplementedActivityServer
}

func (s *service) Like(_ context.Context, req *pb.LikeRequest) (*share.Empty, error) {
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

func (s *service) LikesNumber(_ context.Context, req *pb.LikesNumberRequest) (*pb.LikesNumberResponse, error) {
	mangaUuid, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	num, err := LikesNumber(s.db, mangaUuid)
	if err != nil {
		return nil, err
	}
	return &pb.LikesNumberResponse{Number: num}, nil
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

func (s *service) RateManga(_ context.Context, req *pb.RateMangaRequest) (*share.Empty, error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	err = SaveRate(s.db, userId, mangaId, req.Rate)
	return &share.Empty{}, err
}

func (s *service) AvgMangaRate(_ context.Context, req *pb.AvgMangaRateRequest) (*share.AvgMangaRateResponse, error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	avgRate, err := AvgRate(s.db, mangaId)
	return &share.AvgMangaRateResponse{
		Rate:   avgRate.Rate,
		Voters: avgRate.Voters,
	}, err
}
