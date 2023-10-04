package common

import (
	"google.golang.org/grpc"
	"net"
)

func RunGRPCServer(addr string, cb func(grpc.ServiceRegistrar)) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	serv := grpc.NewServer()

	cb(serv)
	return serv.Serve(lis)
}
