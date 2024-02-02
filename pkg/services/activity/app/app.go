package app

import (
	"context"
	"google.golang.org/grpc"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	pb "micromango/pkg/grpc/activity"
	"micromango/pkg/grpc/catalog"
	client "micromango/pkg/services/activity/clients/catalog"
	"micromango/pkg/services/activity/clients/sqlite"
	"micromango/pkg/services/activity/config"
	"micromango/pkg/services/activity/controller"
	"micromango/pkg/services/activity/usecases"
)

func Run(ctx context.Context, cfg config.Config) <-chan error {
	database, err := sqlite.Connect(cfg.DbAddr)
	if err != nil {
		panic(err)
	}

	conn := utils.GrpcDialOrFatal(cfg.CatalogAddr)
	catalogGrpcClient := catalog.NewCatalogClient(conn)
	catalogService := client.Client{Client: catalogGrpcClient}

	grpcController := &controller.Controller{
		LikeCase: usecases.Like{LikeRepository: database, Catalog: &catalogService},
		RateCase: usecases.Rate{RateRepository: database, Catalog: &catalogService},
		ReadCase: usecases.Read{ReadRepository: database},
	}

	baseServer := grpc.NewServer()
	pb.RegisterActivityServer(baseServer, grpcController)
	return common.StartGrpcService(ctx, cfg.Addr, baseServer)
}
