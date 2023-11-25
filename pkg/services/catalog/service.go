package catalog

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"log"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/static"
)

func Run(ctx context.Context, c Config) <-chan error {
	database := Connect(c.DbAddr)

	conn, err := grpc.Dial(c.ReadingServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	readingService := reading.NewReadingClient(conn)

	conn, err = grpc.Dial(c.StaticServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	staticService := static.NewStaticClient(conn)

	serv := service{
		db:      database,
		reading: readingService,
		static:  staticService,
	}
	baseServer := grpc.NewServer()
	pb.RegisterCatalogServer(baseServer, &serv)

	return common.StartGrpcService(ctx, c.Addr, baseServer)
}

type service struct {
	pb.UnimplementedCatalogServer
	db      *gorm.DB
	reading reading.ReadingClient
	static  static.StaticClient
}

func (s *service) GetManga(ctx context.Context, req *pb.MangaRequest) (*pb.MangaResponse, error) {
	m, err := GetManga(s.db, req.GetMangaId())
	if err != nil {
		return nil, err
	}
	content, err := s.reading.GetMangaContent(ctx, &reading.MangaContentRequest{MangaId: m.MangaId.String()})
	if err != nil {
		return nil, err
	}
	resp := m.ToResponse()
	resp.Content = content
	return resp, nil
}

func (s *service) AddManga(ctx context.Context, req *pb.AddMangaRequest) (*pb.MangaResponse, error) {
	var coverAddr string
	if len(req.Cover) != 0 {
		uploadResp, err := s.static.UploadCover(ctx, &static.UploadImageRequest{Image: req.Cover})
		if err != nil {
			return nil, err
		}
		coverAddr = uploadResp.ImageId
	}

	m, err := AddManga(s.db, Manga{
		Title:       req.Title,
		Cover:       coverAddr,
		Description: utils.DerefOrDefault(req.Description, ""),
	})
	if err != nil {
		return nil, err
	}
	newMangaId := &reading.AddMangaContentRequest{MangaId: m.MangaId.String()}
	if _, err = s.reading.AddMangaContent(ctx, newMangaId); err != nil {
		return nil, err
	}
	return s.GetManga(ctx, &pb.MangaRequest{MangaId: newMangaId.MangaId})
}

func (s *service) GetMangas(context.Context, *pb.Empty) (*pb.MangasResponse, error) {
	ms, err := GetMangas(s.db)
	mangas := make([]*pb.MangaPreviewResponse, len(ms))
	for i, m := range ms {
		mangas[i] = &pb.MangaPreviewResponse{
			MangaId: m.MangaId.String(),
			Title:   m.Title,
			Cover:   m.Cover,
		}
	}
	return &pb.MangasResponse{Mangas: mangas}, err
}
