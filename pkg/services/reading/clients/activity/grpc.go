package activity

import (
	"context"
	"micromango/pkg/grpc/activity"
)

type Repository struct {
	Client activity.ActivityClient
}

func (a *Repository) GetReadChapters(userId, mangaId string) ([]string, error) {
	chapters, err := a.Client.ReadChapters(context.TODO(), &activity.ReadChaptersRequest{
		UserId:  userId,
		MangaId: mangaId,
	})
	if err != nil {
		return nil, err
	}
	return chapters.ChapterIds, nil
}
