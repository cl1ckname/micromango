package gateway

import (
	"github.com/labstack/echo/v4"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/user"
)

func Run() {
	e := echo.New()
	serv := server{}

	e.POST("user/register", serv.Register)
	e.POST("user/login", serv.Login)
	e.GET("user/:userId", serv.Login)

	e.GET("content/:mangaId", serv.GetMangaContent)
	e.POST("content", serv.AddMangaContent)
	e.GET("content/:mangaId/chapter/:chapterId", serv.GetChapter)
	e.POST("content/:mangaId/chapter", serv.AddChapter)
	e.GET("content/:mangaId/chapter/:chapterId/page/:pageId", serv.GetPage)
	e.POST("content/:mangaId/chapter/:chapterId/page", serv.AddChapter)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type server struct {
	e       *echo.Echo
	user    user.UserClient
	reading reading.ReadingClient
	catalog catalog.CatalogClient
}
