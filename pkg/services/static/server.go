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

func (s *server) UploadThumbnail(_ context.Context, req *pb.UploadThumbnailRequest) (*pb.UploadImageResponse, error) {
	img, err := ToImage(req.Thumbnail)
	if err != nil {
		return nil, err
	}
	img = Resize(img, common.COVER_W, common.COVER_H)

	filename := uuid.NewString() + ".jpg"

	if err := s.saveCoverImage(req.MangaId, filename, img); err != nil {
		return nil, err
	}
	extFileName := s.gatewayAddr + "/static/" + path.Join("manga", req.MangaId, filename)
	return &pb.UploadImageResponse{ImageId: extFileName}, nil
}

func (s *server) saveCoverImage(mangaId, filename string, img image.Image) error {
	mangaDirPath := path.Join(s.staticDir, "manga", mangaId)
	if err := createFolderIfNotExists(mangaDirPath); err != nil {
		return err
	}
	filePath := path.Join(mangaDirPath, filename)
	return SaveImage(filePath, img)
}

func (s *server) UploadPage(_ context.Context, req *pb.UploadPageRequest) (*pb.UploadImageResponse, error) {
	img, err := ToImage(req.Page)
	if err != nil {
		return nil, err
	}
	img = Resize(img, common.PAGE_W, common.PAGE_H)

	filename := uuid.NewString() + ".jpg"
	if err := s.savePageImage(req.MangaId, req.ChapterId, filename, img); err != nil {
		return nil, err
	}

	extFileName := s.gatewayAddr + "/static/" + path.Join("manga", req.MangaId, req.ChapterId, filename)
	return &pb.UploadImageResponse{ImageId: extFileName}, nil
}

func (s *server) savePageImage(mangaId, chapterId, filename string, img image.Image) error {
	mangaDirPath := path.Join(s.staticDir, "manga", mangaId)
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

func (s *server) UploadProfilePicture(_ context.Context, req *pb.UploadProfilePictureRequest) (*pb.UploadImageResponse, error) {
	profilePicturePath := path.Join(s.staticDir, "profile")
	if err := createFolderIfNotExists(profilePicturePath); err != nil {
		return nil, err
	}
	profilePicturePath = path.Join(profilePicturePath, req.UserId)
	if err := createFolderIfNotExists(profilePicturePath); err != nil {
		return nil, err
	}
	profilePicturePath = path.Join(profilePicturePath, "profile.jpg")

	img, err := ToImage(req.Picture)
	if err != nil {
		return nil, err
	}
	img = Resize(img, common.PROFILE_W, common.PROFILE_H)
	if err := SaveImage(profilePicturePath, img); err != nil {
		return nil, err
	}
	extPath := s.gatewayAddr + "/" + path.Join("static", "profile", req.UserId, "profile.jpg")
	return &pb.UploadImageResponse{ImageId: extPath}, nil
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
	if err := createFolderIfNotExists(path.Join(c.StaticDir, "manga")); err != nil {
		panic(err)
	}
	if err := createFolderIfNotExists(path.Join(c.StaticDir, "profile")); err != nil {
		panic(err)
	}
	s := server{
		staticDir:   c.StaticDir,
		gatewayAddr: c.GatewayAddr,
	}
	baseServer := grpc.NewServer()
	pb.RegisterStaticServer(baseServer, &s)
	return common.StartGrpcService(ctx, c.ServerAddr, baseServer)
}
