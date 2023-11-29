package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"micromango/pkg/common/utils"
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
		Cover:       []byte("favicon.ico"),
		Description: utils.Ptr("my test manga"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	if resp.Title != "micromango" || resp.Cover != "favicon.ico" || resp.Description != "my test manga" {
		t.Error("invalid fields")
	}
	_, err = uuid.Parse(resp.MangaId)
	require.NoError(t, err)

	getResp, err := cc.GetManga(context.TODO(), &pb.MangaRequest{MangaId: resp.MangaId})
	require.NoError(t, err)
	if getResp.Title != resp.Title {
		t.Error("invalid title")
	}
	if getResp.MangaId != resp.MangaId {
		t.Error("invalid id")
	}
	if getResp.Cover != resp.Cover {
		t.Error("invalid cover: ", getResp.Cover, " ", resp.Cover)
	}

	getsResp, err := cc.GetMangas(context.TODO(), &pb.Empty{})
	if err != nil {
		t.Error("error call getMangas: ", err)
	}
	if len(getsResp.Mangas) != 1 {
		t.Error("empty getMangas response")
	}

	defer cancel()

}
