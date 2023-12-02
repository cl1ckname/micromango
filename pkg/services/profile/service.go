package profile

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/static"
)

type Config struct {
	Addr              string
	DbAddr            string
	StaticServiceAddr string
}

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	conn, err := grpc.Dial(c.StaticServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	staticService := static.NewStaticClient(conn)

	serv := service{
		db:     database,
		static: staticService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterProfileServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedProfileServer
	db     *gorm.DB
	static static.StaticClient
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

func (s *service) GetList(_ context.Context, req *pb.GetListRequest) (*pb.ListResponse, error) {
	lr, err := GetList(s.db, req)
	if err != nil {
		return nil, err
	}
	lrsp := make([]string, len(lr))
	for i := 0; i < len(lr); i++ {
		lrsp[i] = lr[i].MangaId.String()
	}
	return &pb.ListResponse{MangaId: lrsp}, err
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