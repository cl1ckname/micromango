package utils

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"micromango/pkg/common"
	"micromango/pkg/grpc/share"
)

func GrpcDialOrFatal(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func FileFromPb(pbFile *share.File) *common.File {
	if pbFile == nil {
		return nil
	}
	return &common.File{
		Filename: pbFile.Filename,
		File:     pbFile.File,
	}
}
