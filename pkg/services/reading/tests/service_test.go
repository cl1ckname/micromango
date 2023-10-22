package tests

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading"
	"testing"
)

const testServerAddr = ":50011"

func TestService(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reading.Run(ctx, reading.Config{
		Addr:   testServerAddr,
		DbAddr: ":memory:",
	})

	conn, err := grpc.Dial(testServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	cc := pb.NewReadingClient(conn)

	mangaUuid := uuid.NewString()
	res, err := cc.AddMangaContent(context.TODO(), &pb.AddMangaContentRequest{MangaId: mangaUuid})
	require.NoError(t, err)
	expected := &pb.MangaContentResponse{
		MangaId:  mangaUuid,
		Chapters: []*pb.MangaContentResponse_ChapterHead{},
	}
	expBytes, _ := json.Marshal(expected)
	resBytes, _ := json.Marshal(res)
	require.Equal(t, string(expBytes), string(resBytes))

}
