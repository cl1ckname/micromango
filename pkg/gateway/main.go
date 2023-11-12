package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"micromango/pkg/grpc/catalog"
	"micromango/pkg/grpc/reading"
	"micromango/pkg/grpc/user"
	"net/http/httputil"
	"net/url"
	"time"
)

func Run(ctx context.Context, c Config) <-chan error {
	e := echo.New()
	serv := server{}

	conn, err := grpc.Dial(c.UserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	serv.user = user.NewUserClient(conn)

	conn, err = grpc.Dial(c.CatalogAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
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
	applyStaticProxy(e, c)

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

func applyHandlers(e *echo.Echo, serv server) {
	e.POST("api/user/register", serv.Register)
	e.POST("api/user/login", serv.Login)
	e.GET("api/user/:userId", serv.Login)

	e.GET("api/content/:mangaId", serv.GetMangaContent)
	e.POST("api/content", serv.AddMangaContent)
	e.GET("api/content/:mangaId/chapter/:chapterId", serv.GetChapter)
	e.POST("api/content/:mangaId/chapter", serv.AddChapter)
	e.GET("api/content/:mangaId/chapter/:chapterId/page/:pageId", serv.GetPage)
	e.POST("api/content/:mangaId/chapter/:chapterId/page", serv.AddChapter)

	e.GET("api/catalog", serv.GetMangas)
	e.GET("api/catalog/:mangaId", serv.GetManga)
	e.POST("api/catalog", serv.AddManga)
}

func applyStaticProxy(e *echo.Echo, c Config) {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   c.StaticAddr,
	})
	e.Any("/", echo.WrapHandler(proxy))
}

type server struct {
	e       *echo.Echo
	user    user.UserClient
	reading reading.ReadingClient
	catalog catalog.CatalogClient
}
