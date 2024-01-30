package app

import (
	"context"
	"google.golang.org/grpc"
	"micromango/pkg/common"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/activity"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading"
	daoactivity "micromango/pkg/services/reading/dao/activity"
	"micromango/pkg/services/reading/dao/sqlite"
	controller "micromango/pkg/services/reading/grpc"
	"micromango/pkg/services/reading/usecases"
)

func Run(ctx context.Context, c reading.Config) <-chan error {
	db := sqlite.Connect(c.DbAddr)
	conn := utils.GrpcDialOrFatal(c.ActivityServiceAddr)
	activityClient := activity.NewActivityClient(conn)
	activityService := &daoactivity.Repository{Client: activityClient}

	cs := usecases.NewCase(db, activityService)
	server := controller.NewServer(cs)

	baseServer := grpc.NewServer()
	pb.RegisterReadingServer(baseServer, server)
	return common.StartGrpcService(ctx, c.Addr, baseServer)
}
