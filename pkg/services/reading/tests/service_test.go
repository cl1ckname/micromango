package tests

import (
	"context"
	"github.com/golang/protobuf/proto"
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
	require.True(t, proto.Equal(expected, res))

	res, err = cc.GetMangaContent(context.TODO(), &pb.MangaContentRequest{MangaId: mangaUuid})
	require.NoError(t, err)
	require.True(t, proto.Equal(expected, res))

	chapterTitle := "chapter 1"
	cres, err := cc.AddChapter(context.TODO(), &pb.AddChapterRequest{
		MangaId: mangaUuid,
		Title:   chapterTitle,
	})
	chapterId := cres.ChapterId
	_, err = uuid.Parse(chapterId)
	require.NoError(t, err)
	cexpected := &pb.ChapterResponse{
		ChapterId:     chapterId,
		MangaId:       mangaUuid,
		ChapterNumber: 0,
		Title:         chapterTitle,
		Pages:         nil,
	}
	require.True(t, proto.Equal(cres, cexpected))

	expected.Chapters = append(expected.Chapters, &pb.MangaContentResponse_ChapterHead{
		ChapterId:     chapterId,
		ChapterNumber: 0,
		Title:         chapterTitle,
	})
	res, err = cc.GetMangaContent(context.TODO(), &pb.MangaContentRequest{MangaId: mangaUuid})
	require.NoError(t, err)
	require.True(t, proto.Equal(expected, res))

	cres, err = cc.GetChapter(context.TODO(), &pb.ChapterRequest{ChapterId: chapterId})
	require.NoError(t, err)
	require.True(t, proto.Equal(cexpected, cres))

	image := "http://localhost:1234/static/page1.jpg"
	pres, err := cc.AddPage(context.TODO(), &pb.AddPageRequest{
		ChapterId:  chapterId,
		PageNumber: 1,
		Image:      image,
	})
	require.NoError(t, err)
	pid := pres.PageId
	_, err = uuid.Parse(pid)
	require.NoError(t, err)
	pexpected := pb.PageResponse{
		PageId:     pid,
		ChapterId:  chapterId,
		PageNumber: 1,
		Image:      image,
	}
	require.True(t, proto.Equal(&pexpected, pres))

	cexpected.Pages = append(cexpected.Pages, &pb.ChapterResponse_PageHead{
		PageId:     pid,
		PageNumber: 1,
	})
	cres, err = cc.GetChapter(context.TODO(), &pb.ChapterRequest{ChapterId: chapterId})
	require.NoError(t, err)

	require.True(t, proto.Equal(cexpected, cres))
}
