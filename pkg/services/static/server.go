package static

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
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

func (s *server) UploadCover(_ context.Context, req *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {
	id := uuid.NewString()
	var ext string
	var imageBytes []byte
	if req.Type == pb.ImageType_JPG {
		ext = ".jpg"
		imageBytes = req.Image
	} else if req.Type == pb.ImageType_PNG {
		ext = ".png"
		jpgBytes, err := ToJpeg(req.Image)
		if err != nil {
			return nil, err
		}
		imageBytes = jpgBytes
	} else {
		return nil, fmt.Errorf("forbiden image format: %v", req.Type)
	}

	filename := id + ext
	filePath := path.Join(s.staticDir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	if _, err := f.Write(imageBytes); err != nil {
		return nil, err
	}
	extFileName := s.gatewayAddr + "/static/" + filename
	return &pb.UploadImageResponse{ImageId: extFileName}, nil
}

func (s *server) GetImage(_ context.Context, req *pb.GetImageRequest) (*pb.ImageResponse, error) {
	id := req.ImageId
	filepath := path.Join(s.staticDir, id)
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
