package catalog

import (
	"context"
	pb "micromango/pkg/grpc/catalog"
)

type Client struct {
	Client pb.CatalogClient
}

func (c *Client) SetLikesNumber(mangaId string, likes uint64) error {
	_, err := c.Client.SetLikes(context.Background(), &pb.SetLikesRequest{MangaId: mangaId, Likes: likes})
	return err
}

func (c *Client) SetAvgRate(mangaId string, rate float32, voters uint64) error {
	_, err := c.Client.SetAvgRate(context.Background(), &pb.SetAvgRateRateRequest{
		MangaId: mangaId,
		Rate:    rate,
		Rates:   voters,
	})
	return err
}
