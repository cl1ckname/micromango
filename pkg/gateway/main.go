package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	mw "micromango/pkg/gateway/middleware"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/static"
	"micromango/pkg/grpc/user"
	"time"
)

func Run(ctx context.Context, c Config) <-chan error {
	e := echo.New()
	serv := server{}
	e.Use(mw.Cors())
	e.Use(mw.Auth(serv.user))

	serv.connectServices(c)

	applyHandlers(e, serv)

	ok := make(chan error)

	go func() {
		if err := e.Start(c.Addr); err != nil {
			log.Println("Server stopped: ", err.Error())
			ok <- err
			close(ok)
		}
	}()
	go func() {
		<-ctx.Done()
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := e.Shutdown(timeoutCtx); err != nil {
			log.Println("server failed: ", err.Error())
		}
		close(ok)
	}()

	return ok
}

func (s *server) connectServices(c Config) {
	conn, err := grpc.Dial(c.UserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	s.user = user.NewUserClient(conn)

	conn, err = grpc.Dial(c.CatalogAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	s.catalog = catalog.NewCatalogClient(conn)

	conn, err = grpc.Dial(c.ReadingAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	s.reading = reading.NewReadingClient(conn)

	conn, err = grpc.Dial(c.StaticAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	s.static = static.NewStaticClient(conn)

	conn, err = grpc.Dial(c.ProfileAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	s.profile = profile.NewProfileClient(conn)
}

func applyHandlers(e *echo.Echo, serv server) {
	e.POST("api/user/register", serv.Register)
	e.POST("api/user/login", serv.Login)

	e.GET("api/content/:mangaId", serv.GetMangaContent)
	e.POST("api/content", serv.AddMangaContent)
	e.GET("api/content/:mangaId/chapter/:chapterId", serv.GetChapter)
	e.PUT("api/content/:mangaId/chapter/:chapterId", serv.UpdateChapter)
	e.POST("api/content/:mangaId/chapter", serv.AddChapter)
	e.GET("api/content/:mangaId/chapter/:chapterId/page/:pageId", serv.GetPage)
	e.POST("api/content/:mangaId/chapter/:chapterId/page", serv.AddPage)

	e.GET("api/catalog", serv.GetMangas)
	e.POST("api/catalog", serv.AddManga)
	e.GET("api/catalog/:mangaId", serv.GetManga)
	e.PUT("api/catalog/:mangaId", serv.UpdateManga)
	e.DELETE("api/catalog/:mangaId", serv.DeleteManga)

	e.GET("api/profile/:userId", serv.GetProfile)
	e.PUT("api/profile/:userId", serv.UpdateProfile)
	e.GET("api/profile/:userId/list", serv.GetList)
	e.POST("api/profile/:userId/list", serv.AddToList)
	e.DELETE("api/profile/:userId/list", serv.RemoveFromList)

	e.GET("static/:id", serv.GetStatic)
}

type server struct {
	e       *echo.Echo
	user    user.UserClient
	reading reading.ReadingClient
	catalog catalog.CatalogClient
	static  static.StaticClient
	profile profile.ProfileClient
}
