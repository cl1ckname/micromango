package common

import (
	"context"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

func StartGrpcService(ctx context.Context, addr string, baseServer *grpc.Server) <-chan error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	ok := make(chan error)
	go func() {
		log.Info("Starting service on ", addr)
		if err := baseServer.Serve(lis); err != nil {
			log.Fatal("Server stopped: ", err.Error())
			ok <- err
			close(ok)
		}
	}()
	go func() {
		<-ctx.Done()
		baseServer.GracefulStop()
		close(ok)
	}()
	return ok
}
