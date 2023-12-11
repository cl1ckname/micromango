package profile

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
	"micromango/pkg/grpc/catalog"
	pb "micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/static"
)

type Config struct {
	Addr               string
	DbAddr             string
	StaticServiceAddr  string
	CatalogServiceAddr string
}

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	conn := utils.GrpcDialOrFatal(c.StaticServiceAddr)
	staticService := static.NewStaticClient(conn)

	conn = utils.GrpcDialOrFatal(c.CatalogServiceAddr)
	catalogService := catalog.NewCatalogClient(conn)

	serv := service{
		db:      database,
		static:  staticService,
		catalog: catalogService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterProfileServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedProfileServer
	db      *gorm.DB
	static  static.StaticClient
	catalog catalog.CatalogClient
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
	if len(req.Picture) != 0 {
		uploadRes, err := s.static.UploadProfilePicture(ctx, &static.UploadProfilePictureRequest{
			UserId: req.UserId,
			Image:  req.Picture,
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
	lr, err := GetList(s.db, req)
	if err != nil {
		return nil, err
	}

	mangaIdList := utils.Map(lr, func(l ListRecord) string {
		return l.MangaId.String()
	})
	fmt.Println("got list", lr)
	previewList, err := s.catalog.GetList(ctx, &catalog.GetListRequest{MangaList: mangaIdList})
	if err != nil {
		return nil, err
	}

	return &pb.ListResponse{Manga: previewList.PreviewList}, err
}

func (s *service) AddToList(_ context.Context, req *pb.AddToListRequest) (*pb.Empty, error) {
	return &pb.Empty{}, AddToList(s.db, req)
}

func (s *service) RemoveFromList(_ context.Context, req *pb.RemoveFromListRequest) (*pb.Empty, error) {
	if err := RemoveFromList(s.db, req); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &pb.Empty{}, nil
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
