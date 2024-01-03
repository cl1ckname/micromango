package catalog

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
			return nil, status.Error(codes.NotFound, fmt.Sprintf("manga %s not found", req.MangaId))
		}
		return nil, err
	}
	resp := m.ToResponse()

	content, err := s.reading.GetMangaContent(ctx, &reading.MangaContentRequest{MangaId: m.MangaId.String()})
	if err != nil {
		return nil, err
	}
	resp.Content = content

	if req.UserId != nil {
		if err := s.attachUserMangaFields(ctx, *req.UserId, req.MangaId, resp); err != nil {
			return nil, err
		}
	}

	listStats, err := s.profile.ListStats(ctx, &profile.ListStatsRequests{MangaId: req.MangaId})
	if err != nil {
		return nil, err
	}
	resp.ListStats = listStats.Stats

	return resp, nil
}

func (s *service) attachUserMangaFields(ctx context.Context, userId, mangaId string, resp *pb.MangaResponse) error {
	isInResp, err := s.profile.IsInList(ctx, &profile.IsInListRequest{UserId: userId, MangaId: mangaId})
	if err != nil {
		return err
	}
	resp.List = isInResp.In

	hasLike, err := s.activity.HasLike(ctx, &activity.HasLikeRequest{MangaId: mangaId, UserId: userId})
	if err != nil {
		return err
	}
	resp.Liked = hasLike.Has

	userRate, err := s.activity.UserRate(ctx, &activity.UserRateRequest{UserId: userId, MangaId: mangaId})
	if err != nil {
		return err
	}
	resp.UserRate = userRate.Rate
	return nil
}

func (s *service) AddManga(ctx context.Context, req *pb.AddMangaRequest) (*pb.MangaResponse, error) {
	mangaId := uuid.New()
	var coverAddr string
	if req.Thumbnail != nil {
		uploadResp, err := s.static.UploadThumbnail(ctx, &static.UploadThumbnailRequest{
			MangaId:   mangaId.String(),
			Thumbnail: req.Thumbnail,
		})
		if err != nil {
			return nil, err
		}
		coverAddr = uploadResp.ImageId
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	m, err := AddManga(s.db, Manga{
		MangaId:     mangaId,
		Title:       req.Title,
		Cover:       coverAddr,
		Description: utils.DerefOrDefault(req.Description, ""),
		Genres:      utils.Map(req.Genres, func(i uint32) Genre { return Genre{GenreId: int(i)} }),
		CreatedBy:   userId,
		UpdatedBy:   userId,
	})
	if err != nil {
		return nil, err
	}
	return s.GetManga(ctx, &pb.MangaRequest{MangaId: m.MangaId.String()})
}

func (s *service) GetMangas(_ context.Context, request *pb.GetMangasRequest) (*pb.MangasResponse, error) {
	opts := GetMangaOpts{
		Include: request.GenresInclude,
		Exclude: request.GenresExclude,
		Starts:  request.Starts,
		Asc:     request.Asc,
	}
	if request.Order != nil {
		opts.Order = utils.Ptr(Order(*request.Order))
	}

	ms, err := GetMangas(s.db, opts)
	mangas := utils.Map(ms, func(m Manga) *share.MangaPreviewResponse {
		return &share.MangaPreviewResponse{
			MangaId: m.MangaId.String(),
			Title:   m.Title,
			Cover:   m.Cover,
			Rate:    m.Rate,
		}
	})

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
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	mangaToUpdate.UpdatedBy = userId
	mangaToUpdate.Title = utils.DerefOrDefault(req.Title, mangaToUpdate.Title)
	mangaToUpdate.Description = utils.DerefOrDefault(req.Description, mangaToUpdate.Description)
	genres := utils.Map(req.Genres, func(i uint32) Genre {
		return Genre{GenreId: int(i)}
	})
	mangaToUpdate.Genres = genres
	if req.Thumbnail != nil {
		uploadResp, err := s.static.UploadThumbnail(ctx, &static.UploadThumbnailRequest{
			MangaId:   req.MangaId,
			Thumbnail: req.Thumbnail,
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
			Rate:    m.Rate,
		}
	})

	return &pb.GetListResponse{PreviewList: list}, nil
}

func (s *service) SetAvgRate(_ context.Context, req *pb.SetAvgRateRateRequest) (*pb.Empty, error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	if err := RateManga(s.db, mangaId, req.Rate, req.Rates); err != nil {
		return nil, err
	}
	return &pb.Empty{}, err
}

func (s *service) SetLikes(_ context.Context, req *pb.SetLikesRequest) (*pb.Empty, error) {
	mangaId, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	if err := LikeManga(s.db, mangaId, req.Likes); err != nil {
		return nil, err
	}
	return &pb.Empty{}, err
}

func (s *service) LastUpdates(_ context.Context, req *pb.LastUpdatesRequest) (*pb.LastUpdatesResponse, error) {
	m, err := LastUpdates(s.db, req.Page, req.Number)
	if err != nil {
		return nil, err
	}
	responseList := utils.Map(m, func(i Manga) *share.MangaPreviewResponse {
		return i.ToPreview()
	})
	return &pb.LastUpdatesResponse{Manga: responseList}, nil
}
