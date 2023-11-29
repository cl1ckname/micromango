package profile

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
