package catalog

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	pb "micromango/pkg/grpc/catalog"
	"micromango/pkg/services/catalog/db"
	"net"
)

type service struct {
	pb.UnimplementedCatalogServer
	db *gorm.DB
}

func (s *service) GetManga(_ context.Context, req *pb.MangaRequest) (*pb.MangaResponse, error) {
	m, err := db.GetManga(s.db, req.GetMangaId())
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
	m, err := db.AddManga(s.db, req)
	return &pb.MangaResponse{
		MangaId:       m.MangaId.String(),
		Title:         m.Title,
		Cover:         m.Cover,
		Description:   m.Description,
		ChapterNumber: 0,
	}, err
}

func Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50002))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	database := db.Connect("catalog.sqlite")
	serv := service{
		db: database,
	}
	pb.RegisterCatalogServer(s, &serv)
	log.Println("customer listens 3234")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
