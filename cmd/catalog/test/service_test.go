package test

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "micromango/pkg/grpc/catalog"
	"micromango/pkg/services/catalog"
	"testing"
)

const testServerAddr = ":50010"

func TestService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	catalog.Run(ctx, catalog.Config{
		Addr:   testServerAddr,
		DbAddr: ":memory:",
	})

	conn, err := grpc.Dial(testServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
	}
	cc := pb.NewCatalogClient(conn)

	resp, err := cc.AddManga(context.TODO(), &pb.AddMangaRequest{
		Title:       "micromango",
		Cover:       "favicon.ico",
		Description: "my test manga",
	})
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("empty response")
	}
	if resp.Title != "micromango" || resp.Cover != "favicon.ico" || resp.Description != "my test manga" {
		t.Error("invalid fields")
	}
	if _, err := uuid.Parse(resp.MangaId); err != nil {
		t.Error("invalid manga id: ", err.Error())
	}

	getResp, err := cc.GetManga(context.TODO(), &pb.MangaRequest{MangaId: resp.MangaId})
	if err != nil {
		t.Error(err)
	}
	if getResp.Title != resp.Title {
		t.Error("invalid title")
	}
	if getResp.MangaId != resp.MangaId {
		t.Error("invalid id")
	}
	if getResp.Cover != resp.Cover {
		t.Error("invalid cover: ", getResp.Cover, " ", resp.Cover)
	}

	defer cancel()

}
