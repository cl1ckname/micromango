package catalog

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/catalog"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)
	serv := service{
		db: database,
	}
	baseServer := grpc.NewServer()
	pb.RegisterCatalogServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedCatalogServer
	db *gorm.DB
}

func (s *service) GetManga(_ context.Context, req *pb.MangaRequest) (*pb.MangaResponse, error) {
	m, err := GetManga(s.db, req.GetMangaId())
	return m.ToResponse(), err
}

func (s *service) AddManga(_ context.Context, req *pb.AddMangaRequest) (*pb.MangaResponse, error) {
	m, err := AddManga(s.db, req)
	return m.ToResponse(), err
}

func (s *service) GetMangas(context.Context, *pb.Empty) (*pb.MangasResponse, error) {
	ms, err := GetMangas(s.db)
	mangas := make([]*pb.MangaResponse, len(ms))
	for i, m := range ms {
		mangas[i] = m.ToResponse()
	}
	return &pb.MangasResponse{Mangas: mangas}, err
}
