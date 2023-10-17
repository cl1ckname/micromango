package catalog

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/catalog"
)

func Run(c Config) {
	database := Connect(c.DbAddr)
	serv := service{
		db: database,
	}
	addr := fmt.Sprintf(c.Addr)
	log.Info("Starting catalog service on ", addr)
	if err := common.RunGRPCServer(addr, func(registrar grpc.ServiceRegistrar) {
		pb.RegisterCatalogServer(registrar, &serv)
	}); err != nil {
		panic(err)
	}
}

type service struct {
	pb.UnimplementedCatalogServer
	db *gorm.DB
}

func (s *service) GetManga(_ context.Context, req *pb.MangaRequest) (*pb.MangaResponse, error) {
	m, err := GetManga(s.db, req.GetMangaId())
	if err != nil {
		return nil, err
	}
	return &pb.MangaResponse{
		MangaId:       m.MangaId.String(),
		Title:         m.Title,
		Cover:         m.Title,
		Description:   m.Description,
		ChapterNumber: 0,
	}, nil
}

func (s *service) AddManga(_ context.Context, req *pb.AddMangaRequest) (*pb.MangaResponse, error) {
	m, err := AddManga(s.db, req)
	return &pb.MangaResponse{
		MangaId:       m.MangaId.String(),
		Title:         m.Title,
		Cover:         m.Cover,
		Description:   m.Description,
		ChapterNumber: 0,
	}, err
}
