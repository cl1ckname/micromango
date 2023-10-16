package gateway

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/user"
)

func Run(c Config) {
	e := echo.New()
	serv := server{}

	conn, err := grpc.Dial(c.UserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	serv.user = user.NewUserClient(conn)

	conn, err = grpc.Dial(c.CatalogAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	serv.catalog = catalog.NewCatalogClient(conn)

	conn, err = grpc.Dial(c.ReadingAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	serv.reading = reading.NewReadingClient(conn)

	applyHandlers(e, serv)

	panic(e.Start(c.Addr))
}

func applyHandlers(e *echo.Echo, serv server) {
	e.POST("user/register", serv.Register)
	e.POST("user/login", serv.Login)
	e.GET("user/:userId", serv.Login)

	e.GET("content/:mangaId", serv.GetMangaContent)
	e.POST("content", serv.AddMangaContent)
	e.GET("content/:mangaId/chapter/:chapterId", serv.GetChapter)
	e.POST("content/:mangaId/chapter", serv.AddChapter)
	e.GET("content/:mangaId/chapter/:chapterId/page/:pageId", serv.GetPage)
	e.POST("content/:mangaId/chapter/:chapterId/page", serv.AddChapter)

	e.GET("catalog/:mangaId", serv.GetManga)
	e.POST("catalog", serv.AddManga)
}

type server struct {
	e       *echo.Echo
	user    user.UserClient
	reading reading.ReadingClient
	catalog catalog.CatalogClient
}
