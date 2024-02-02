package controller

import (
	pb "micromango/pkg/grpc/activity"
	"micromango/pkg/services/activity/entity"
)

func rateFromProto(pb *pb.RateMangaRequest) entity.Rate {
	return entity.Rate{
		MangaId: pb.MangaId,
		Rate:    pb.Rate,
		UserId:  pb.UserId,
	}
}
