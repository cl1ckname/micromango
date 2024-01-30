package grpc

import (
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading/entity"
)

func ChapterToPb(c entity.Chapter) *pb.ChapterResponse {
	return &pb.ChapterResponse{
		ChapterId: c.ChapterId,
		MangaId:   c.MangaId,
		Number:    c.Number,
		Title:     c.Title,
		Pages: utils.Map(c.Pages, func(p entity.Page) *pb.ChapterResponse_PageHead {
			return &pb.ChapterResponse_PageHead{
				PageId: p.PageId,
				Number: p.Number,
				Image:  p.Image,
			}
		}),
		CreatedAt: "",
	}
}

func ChapterToHead(c entity.ChapterHead) *pb.MangaContentResponse_ChapterHead {
	return &pb.MangaContentResponse_ChapterHead{
		ChapterId: c.ChapterId,
		Number:    c.Number,
		Title:     c.Title,
		Pages:     c.Pages,
		Read:      c.Read,
		CreatedAt: c.CreatedAt.String(),
	}
}

func PageToPb(p entity.Page) *pb.PageResponse {
	return &pb.PageResponse{
		PageId: p.PageId,
		Number: p.Number,
		Image:  p.Image,
	}
}
