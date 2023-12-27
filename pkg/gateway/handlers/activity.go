package handlers

import (
	"context"
	"fmt"
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
	activityGroup.POST("/manga/:mangaId/rate", h.Rate)
	activityGroup.POST("/chapter/:chapterId/read", h.Read)
}

func (h *activityHandler) Like(ctx echo.Context) error {
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		fmt.Println("auth", auth, ok)
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
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "Unauthorized")
	}
	mangaId := ctx.Param("mangaId")
	dislikeReq := &activity.DislikeRequest{
		MangaId: mangaId,
		UserId:  auth.UserId,
	}

	if _, err := h.activity.Dislike(context.TODO(), dislikeReq); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return utils.JsonMessage(ctx, http.StatusCreated, "Liked")
}

func (h *activityHandler) Rate(ctx echo.Context) error {
	var req activity.RateMangaRequest
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "Unauthorized")
	}
	if err := ctx.Bind(&req); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	req.MangaId = ctx.Param("mangaId")
	req.UserId = auth.UserId

	if _, err := h.activity.RateManga(context.TODO(), &req); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return utils.JsonMessage(ctx, http.StatusCreated, "Rated")
}

func (h *activityHandler) Read(ctx echo.Context) error {
	var req activity.ReadChapterRequest
	req.ChapterId = ctx.Param("chapterId")
	auth, ok := ctx.Get("claims").(*user.UserResponse)
	if !ok {
		return utils.JsonMessage(ctx, http.StatusUnauthorized, "Unauthorized")
	}
	req.UserId = auth.UserId
	if _, err := h.activity.ReadChapter(context.TODO(), &req); err != nil {
		return utils.ErrorToResponse(ctx, err)
	}
	return utils.JsonMessage(ctx, http.StatusOK, "Read")
}
