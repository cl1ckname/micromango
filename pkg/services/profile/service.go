package profile

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/activity"
	"micromango/pkg/grpc/catalog"
	pb "micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/share"
	"micromango/pkg/grpc/static"
)

type Config struct {
	Addr                string
	DbAddr              string
	StaticServiceAddr   string
	CatalogServiceAddr  string
	ActivityServiceAddr string
}

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	conn := utils.GrpcDialOrFatal(c.StaticServiceAddr)
	staticService := static.NewStaticClient(conn)

	conn = utils.GrpcDialOrFatal(c.CatalogServiceAddr)
	catalogService := catalog.NewCatalogClient(conn)

	conn = utils.GrpcDialOrFatal(c.ActivityServiceAddr)
	activityService := activity.NewActivityClient(conn)

	serv := service{
		db:       database,
		static:   staticService,
		catalog:  catalogService,
		activity: activityService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterProfileServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedProfileServer
	db       *gorm.DB
	static   static.StaticClient
	catalog  catalog.CatalogClient
	activity activity.ActivityClient
}

func (s *service) Create(_ context.Context, req *pb.CreateRequest) (*pb.Response, error) {
	userUuid, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	profile := Profile{
		UserId:   userUuid,
		Username: req.Username,
	}
	newProfile, err := CreateProfile(s.db, profile)
	if err != nil {
		return nil, err
	}
	return newProfile.ToResponse(), nil
}

func (s *service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Response, error) {
	userUuid, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	userToUpdate, err := FindOne(s.db, userUuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "user "+req.UserId+" not found")
		}
		return nil, err
	}

	userToUpdate.Username = utils.DerefOrDefault(req.Username, userToUpdate.Username)
	userToUpdate.Bio = utils.DerefOrDefault(req.Bio, userToUpdate.Bio)
	if req.Picture != nil {
		uploadRes, err := s.static.UploadProfilePicture(ctx, &static.UploadProfilePictureRequest{
			UserId:  req.UserId,
			Picture: req.Picture,
		})
		if err != nil {
			return nil, err
		}
		userToUpdate.Picture = uploadRes.ImageId
	}
	updatedProfile, err := SaveProfile(s.db, userToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedProfile.ToResponse(), nil
}

func (s *service) Get(_ context.Context, req *pb.GetRequest) (*pb.Response, error) {
	userUuid, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	p, err := FindOne(s.db, userUuid)
	return p.ToResponse(), err
}

func (s *service) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.ListResponse, error) {
	userId, err := uuid.Parse(req.ProfileId)
	if err != nil {
		return nil, err
	}

	resp := pb.ListResponse{
		Lists: make(map[uint32]*pb.ListResponse_ListMapField),
	}
	for i := share.ListName(0); i < 5; i++ {
		mangaList, err := GetList(s.db, userId, i)
		if err != nil {
			return nil, err
		}
		mangaIdList := utils.Map(mangaList, func(l ListRecord) string {
			return l.MangaId.String()
		})
		previewList, err := s.catalog.GetList(ctx, &catalog.GetListRequest{MangaList: mangaIdList})
		if err != nil {
			return nil, err
		}
		rates, err := s.activity.UserRateList(ctx, &activity.UserRateListRequest{
			UserId:  req.ProfileId,
			MangaId: mangaIdList,
		})
		if err != nil {
			return nil, err
		}

		field := &pb.ListResponse_ListMapField{
			Value: utils.Map(previewList.PreviewList, func(c *share.MangaPreviewResponse) *pb.ListResponse_ListEntry {
				rate := rates.Rates[c.MangaId]
				return &pb.ListResponse_ListEntry{
					MangaId: c.MangaId,
					Title:   c.Title,
					Rate:    utils.Ptr(rate),
				}
			}),
		}

		resp.Lists[uint32(i)] = field
	}

	return &resp, err
}

func (s *service) AddToList(_ context.Context, req *pb.AddToListRequest) (*share.Empty, error) {
	return &share.Empty{}, AddToList(s.db, req)
}

func (s *service) RemoveFromList(_ context.Context, req *pb.RemoveFromListRequest) (*share.Empty, error) {
	if err := RemoveFromList(s.db, req); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &share.Empty{}, nil
}

func (s *service) IsInList(_ context.Context, req *pb.IsInListRequest) (*pb.IsInListResponse, error) {
	return FindListRecord(s.db, req)
}

func (s *service) ListStats(_ context.Context, req *pb.ListStatsRequests) (*pb.ListStatsResponse, error) {
	mangaUuid, err := uuid.Parse(req.MangaId)
	if err != nil {
		return nil, err
	}
	res, err := ListStats(s.db, mangaUuid)
	if err != nil {
		return nil, err
	}
	return &pb.ListStatsResponse{Stats: res}, nil
}
