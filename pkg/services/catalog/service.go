package catalog

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/activity"
	pb "micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/share"
	"micromango/pkg/grpc/static"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	conn := utils.GrpcDialOrFatal(c.ReadingServiceAddr)
	readingService := reading.NewReadingClient(conn)

	conn = utils.GrpcDialOrFatal(c.StaticServiceAddr)
	staticService := static.NewStaticClient(conn)

	conn = utils.GrpcDialOrFatal(c.ProfileServiceAddr)
	profileService := profile.NewProfileClient(conn)

	conn = utils.GrpcDialOrFatal(c.ActivityServiceAddr)
	activityService := activity.NewActivityClient(conn)

	serv := service{
		db:       database,
		reading:  readingService,
		static:   staticService,
		profile:  profileService,
		activity: activityService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterCatalogServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedCatalogServer
	db       *gorm.DB
	reading  reading.ReadingClient
	static   static.StaticClient
	profile  profile.ProfileClient
	activity activity.ActivityClient
}

func (s *service) GetManga(ctx context.Context, req *pb.MangaRequest) (*pb.MangaResponse, error) {
	m, err := GetManga(s.db, req.GetMangaId())
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("manga", req.MangaId, "not found")
			return nil, status.Error(codes.NotFound, fmt.Sprintf("manga %s not found", req.MangaId))
		}
		return nil, err
	}
	content, err := s.reading.GetMangaContent(ctx, &reading.MangaContentRequest{MangaId: m.MangaId.String()})
	if err != nil {
		log.Println("get manga", req.MangaId, "content error:", err.Error())
		return nil, err
	}
	resp := m.ToResponse()

	if req.UserId != nil {
		isInResp, err := s.profile.IsInList(ctx, &profile.IsInListRequest{UserId: *req.UserId, MangaId: req.MangaId})
		if err != nil {
			return nil, err
		}
		resp.List = isInResp.In
	}
	resp.Content = content

	likesNumber, err := s.activity.LikesNumber(ctx, &activity.LikesNumberRequest{MangaId: req.MangaId})
	if err != nil {
		return nil, err
	}
	resp.Likes = likesNumber.Number

	if req.UserId != nil {
		userId := *req.UserId
		hasLike, err := s.activity.HasLike(ctx, &activity.HasLikeRequest{MangaId: req.MangaId, UserId: userId})
		if err != nil {
			return nil, err
		}
		resp.Liked = hasLike.Has

		userRate, err := s.activity.UserRate(ctx, &activity.UserRateRequest{UserId: userId, MangaId: req.MangaId})
		if err != nil {
			return nil, err
		}
		resp.UserRate = userRate.Rate
	}

	listStats, err := s.profile.ListStats(ctx, &profile.ListStatsRequests{MangaId: req.MangaId})
	if err != nil {
		return nil, err
	}
	resp.ListStats = listStats.Stats

	rate, err := s.activity.AvgMangaRate(ctx, &activity.AvgMangaRateRequest{MangaId: req.MangaId})
	if err != nil {
		return nil, err
	}
	resp.Rate = &share.AvgMangaRateResponse{
		Rate:   rate.Rate,
		Voters: rate.Voters,
	}
	return resp, nil
}

func (s *service) AddManga(ctx context.Context, req *pb.AddMangaRequest) (*pb.MangaResponse, error) {
	mangaId := uuid.New()
	var coverAddr string
	if len(req.Cover) != 0 {
		uploadResp, err := s.static.UploadCover(ctx, &static.UploadCoverRequest{
			MangaId: mangaId.String(),
			Image:   req.Cover,
			Type:    0, // FIXME
		})
		if err != nil {
			return nil, err
		}
		coverAddr = uploadResp.ImageId
	}

	m, err := AddManga(s.db, Manga{
		MangaId:     mangaId,
		Title:       req.Title,
		Cover:       coverAddr,
		Description: utils.DerefOrDefault(req.Description, ""),
		Genres:      utils.Map(req.Genres, func(i uint32) Genre { return Genre{GenreId: int(i)} }),
	})
	if err != nil {
		return nil, err
	}
	return s.GetManga(ctx, &pb.MangaRequest{MangaId: m.MangaId.String()})
}

func (s *service) GetMangas(ctx context.Context, request *pb.GetMangasRequest) (*pb.MangasResponse, error) {
	ms, err := GetMangas(s.db, request.GenresInclude, request.GenresExclude, request.Starts)
	mangas := make([]*share.MangaPreviewResponse, len(ms))
	for i, m := range ms {
		rate, err := s.activity.AvgMangaRate(ctx, &activity.AvgMangaRateRequest{MangaId: m.MangaId.String()})
		if err != nil {
			return nil, err
		}
		mangas[i] = &share.MangaPreviewResponse{
			MangaId: m.MangaId.String(),
			Title:   m.Title,
			Cover:   m.Cover,
			Rate:    rate.Rate,
		}
	}
	return &pb.MangasResponse{Mangas: mangas}, err
}

func (s *service) UpdateManga(ctx context.Context, req *pb.UpdateMangaRequest) (*pb.MangaResponse, error) {
	mangaToUpdate, err := GetManga(s.db, req.MangaId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("manga %s not found", req.MangaId))
		}
		return nil, err
	}
	mangaToUpdate.Title = utils.DerefOrDefault(req.Title, mangaToUpdate.Title)
	mangaToUpdate.Description = utils.DerefOrDefault(req.Description, mangaToUpdate.Description)
	genres := utils.Map(req.Genres, func(i uint32) Genre {
		return Genre{GenreId: int(i)}
	})
	mangaToUpdate.Genres = genres
	if len(req.Cover) != 0 {
		uploadResp, err := s.static.UploadCover(ctx, &static.UploadCoverRequest{
			MangaId: req.MangaId,
			Image:   req.Cover,
			Type:    0, // FIXME
		})
		if err != nil {
			return nil, err
		}
		mangaToUpdate.Cover = uploadResp.ImageId
	}
	updatedManga, err := SaveManga(s.db, mangaToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedManga.ToResponse(), nil
}

func (s *service) DeleteManga(_ context.Context, req *pb.DeleteMangaRequest) (*pb.Empty, error) {
	empty := new(pb.Empty)
	return empty, DeleteManga(s.db, req.MangaId)
}

func (s *service) GetList(_ context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {
	if len(req.MangaList) == 0 {
		return &pb.GetListResponse{}, nil
	}
	m, err := GetMany(s.db, req.MangaList)
	if err != nil {
		return nil, err
	}

	list := utils.Map(m, func(m Manga) *share.MangaPreviewResponse {
		return &share.MangaPreviewResponse{
			MangaId: m.MangaId.String(),
			Title:   m.Title,
			Cover:   m.Cover,
		}
	})

	return &pb.GetListResponse{PreviewList: list}, nil
}
