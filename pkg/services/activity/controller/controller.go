package controller

import (
	"context"
	pb "micromango/pkg/grpc/activity"
	"micromango/pkg/grpc/share"
	"micromango/pkg/services/activity/entity"
	"micromango/pkg/services/activity/usecases"
)

type Controller struct {
	pb.UnimplementedActivityServer
	LikeCase usecases.Like
	RateCase usecases.Rate
	ReadCase usecases.Read
}

func (c *Controller) Like(_ context.Context, req *pb.LikeRequest) (*share.Empty, error) {
	if err := c.LikeCase.Like(req.UserId, req.MangaId); err != nil {
		return nil, err
	}
	return &share.Empty{}, nil
}

func (c *Controller) Dislike(_ context.Context, req *pb.DislikeRequest) (*share.Empty, error) {
	if err := c.LikeCase.Dislike(req.UserId, req.MangaId); err != nil {
		return nil, err
	}
	return &share.Empty{}, nil
}

func (c *Controller) HasLike(_ context.Context, req *pb.HasLikeRequest) (*pb.HasLikeResponse, error) {
	has, err := c.LikeCase.HasLike(req.UserId, req.MangaId)
	if err != nil {
		return nil, err
	}
	return &pb.HasLikeResponse{Has: has}, nil
}

func (c *Controller) RateManga(_ context.Context, req *pb.RateMangaRequest) (*share.Empty, error) {
	err := c.RateCase.RateManga(req.UserId, req.MangaId, req.Rate)
	return &share.Empty{}, err
}

func (c *Controller) UserRate(_ context.Context, req *pb.UserRateRequest) (*pb.UserRateResponse, error) {
	rate, err := c.RateCase.GetUserRate(req.UserId, req.MangaId)
	if err != nil {
		return nil, err
	}
	return &pb.UserRateResponse{Rate: rate}, nil
}

func (c *Controller) UserRateList(_ context.Context, req *pb.UserRateListRequest) (*pb.UserRateListResponse, error) {
	rates, err := c.RateCase.GetRateList(req.UserId, req.MangaId)
	if err != nil {
		return nil, err
	}
	return &pb.UserRateListResponse{Rates: rates}, nil
}

func (c *Controller) ReadChapter(_ context.Context, req *pb.ReadChapterRequest) (*share.Empty, error) {
	if err := c.ReadCase.ReadChapter(entity.ReadRecord{
		MangaId:   req.MangaId,
		ChapterId: req.ChapterId,
		UserId:    req.UserId,
	}); err != nil {
		return nil, err
	}
	return &share.Empty{}, nil
}

func (c *Controller) ReadChapters(_ context.Context, req *pb.ReadChaptersRequest) (*pb.ReadChaptersResponse, error) {
	uids, err := c.ReadCase.GetReadChapters(req.UserId, req.MangaId)
	if err != nil {
		return nil, err
	}
	return &pb.ReadChaptersResponse{ChapterIds: uids}, nil
}
