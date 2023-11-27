package static

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "golang.org/x/image/draw"
	"google.golang.org/grpc"
	"image"
	"io"
	"log"
	"micromango/pkg/common"
	pb "micromango/pkg/grpc/static"
	"os"
	"path"
)

type Config struct {
	ServerAddr  string
	GatewayAddr string
	StaticDir   string
}

type server struct {
	staticDir   string
	gatewayAddr string
	pb.UnimplementedStaticServer
}

func (s *server) UploadCover(_ context.Context, req *pb.UploadCoverRequest) (*pb.UploadImageResponse, error) {
	img, err := ToImage(req.Image, req.Type)
	if err != nil {
		return nil, err
	}
	img = Resize(img, common.COVER_W, common.COVER_H)

	filename := uuid.NewString() + ".jpg"

	if err := s.saveCoverImage(req.MangaId, filename, img); err != nil {
		return nil, err
	}
	extFileName := s.gatewayAddr + "/static/" + path.Join("mangas", req.MangaId, filename)
	return &pb.UploadImageResponse{ImageId: extFileName}, nil
}

func (s *server) saveCoverImage(mangaId, filename string, img image.Image) error {
	mangaDirPath := path.Join(s.staticDir, "mangas", mangaId)
	if err := createFolderIfNotExists(mangaDirPath); err != nil {
		return err
	}
	filePath := path.Join(mangaDirPath, filename)
	log.Println(filePath)
	return SaveImage(filePath, img)
}

func (s *server) UploadPage(_ context.Context, req *pb.UploadPageRequest) (*pb.UploadImageResponse, error) {
	img, err := ToImage(req.Image, req.Type)
	if err != nil {
		return nil, err
	}
	img = Resize(img, common.PAGE_W, common.PAGE_H)

	filename := uuid.NewString() + ".jpg"
	if err := s.savePageImage(req.MangaId, req.ChapterId, filename, img); err != nil {
		return nil, err
	}

	extFileName := s.gatewayAddr + "/static/" + path.Join("mangas", req.MangaId, req.ChapterId, filename)
	return &pb.UploadImageResponse{ImageId: extFileName}, nil
}

func (s *server) savePageImage(mangaId, chapterId, filename string, img image.Image) error {
	mangaDirPath := path.Join(s.staticDir, "mangas", mangaId)
	if err := createFolderIfNotExists(mangaDirPath); err != nil {
		return err
	}
	chapterDirPath := path.Join(mangaDirPath, chapterId)
	if err := createFolderIfNotExists(chapterDirPath); err != nil {
		return err
	}
	filePath := path.Join(chapterDirPath, filename)
	return SaveImage(filePath, img)
}

func createFolderIfNotExists(folderPath string) error {
	if _, err := os.Stat(folderPath); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(folderPath, 0700)
			if err != nil {
				return fmt.Errorf("create folder error: %v", err)
			}
			return nil
		}
		return fmt.Errorf("stat error: %v", err)
	}
	return nil
}

func (s *server) GetImage(_ context.Context, req *pb.GetImageRequest) (*pb.ImageResponse, error) {
	id := req.ImageId
	filepath := path.Join(s.staticDir, id)
	log.Println(filepath)
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	imageBytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return &pb.ImageResponse{Image: imageBytes}, nil
}

func Run(ctx context.Context, c Config) <-chan error {
	s := server{
		staticDir:   c.StaticDir,
		gatewayAddr: c.GatewayAddr,
	}
	baseServer := grpc.NewServer()
	pb.RegisterStaticServer(baseServer, &s)
	return common.StartGrpcService(ctx, c.ServerAddr, baseServer)
}
