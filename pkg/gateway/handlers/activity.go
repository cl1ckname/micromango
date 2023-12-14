package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/activity"
	"micromango/pkg/grpc/user"
	"net/http"
)

type activityHandler struct {
	activity activity.ActivityClient
}

func RegisterActivity(g *echo.Group, a activity.ActivityClient) {
	h := activityHandler{a}
	activityGroup := g.Group("/activity")

	activityGroup.POST("/manga/:mangaId/like", h.Like)
	activityGroup.DELETE("/manga/:mangaId/like", h.Dislike)
}

func (h *activityHandler) Like(ctx echo.Context) error {
	auth, ok := ctx.Get("auth").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "Unauthorized")
	}
	mangaId := ctx.Param("mangaId")
	likeReq := &activity.LikeRequest{
		MangaId: mangaId,
		UserId:  auth.UserId,
	}

	if _, err := h.activity.Like(context.TODO(), likeReq); err != nil {
		return utils.JsonMessage(ctx, http.StatusBadRequest, err.Error())
	}
	return utils.JsonMessage(ctx, http.StatusCreated, "Liked")
}

func (h *activityHandler) Dislike(ctx echo.Context) error {
	auth, ok := ctx.Get("auth").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "Unauthorized")
	}
	mangaId := ctx.Param("mangaId")
	dislikeReq := &activity.DislikeRequest{
		MangaId: mangaId,
		UserId:  auth.UserId,
	}

	if _, err := h.activity.Dislike(context.TODO(), dislikeReq); err != nil {
		return utils.JsonMessage(ctx, http.StatusBadRequest, err.Error())
	}
	return utils.JsonMessage(ctx, http.StatusCreated, "Liked")
}
