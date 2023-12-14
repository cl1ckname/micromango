package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"micromango/pkg/common/utils"
	"micromango/pkg/gateway/handlers"
	mw "micromango/pkg/gateway/middleware"
	"micromango/pkg/grpc/activity"
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
	serv.connectServices(c)

	e.Use(mw.Cors())
	e.Use(mw.Auth(serv.user))

	applyHandlers(e, serv)

	return listenUntilError(ctx, c.Addr, e)
}

func listenUntilError(ctx context.Context, addr string, e *echo.Echo) <-chan error {
	ok := make(chan error)
	go pipeEchoErrorToChan(addr, e, ok)
	go waitContextShutdown(ctx, e, ok)
	return ok
}

func pipeEchoErrorToChan(addr string, e *echo.Echo, ok chan error) {
	if err := e.Start(addr); err != nil {
		log.Println("Server stopped: ", err.Error())
		ok <- err
		close(ok)
	}
}

func waitContextShutdown(ctx context.Context, e *echo.Echo, ok chan error) {
	<-ctx.Done()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := e.Shutdown(timeoutCtx); err != nil {
		log.Println("server failed: ", err.Error())
	}
	close(ok)
}

func (s *server) connectServices(c Config) {
	conn := utils.GrpcDialOrFatal(c.UserAddr)
	s.user = user.NewUserClient(conn)

	conn = utils.GrpcDialOrFatal(c.CatalogAddr)
	s.catalog = catalog.NewCatalogClient(conn)

	conn = utils.GrpcDialOrFatal(c.ReadingAddr)
	s.reading = reading.NewReadingClient(conn)

	conn = utils.GrpcDialOrFatal(c.StaticAddr)
	s.static = static.NewStaticClient(conn)

	conn = utils.GrpcDialOrFatal(c.ProfileAddr)
	s.profile = profile.NewProfileClient(conn)

	conn = utils.GrpcDialOrFatal(c.ActivityAddr)
	s.activity = activity.NewActivityClient(conn)
}

func applyHandlers(e *echo.Echo, serv server) {
	apiGroup := e.Group("/api")

	handlers.RegisterUser(apiGroup, serv.user)
	handlers.RegisterCatalog(apiGroup, serv.catalog)
	handlers.RegisterProfile(apiGroup, serv.profile)
	handlers.RegisterReading(apiGroup, serv.reading)
	handlers.RegisterActivity(apiGroup, serv.activity)
	e.GET("static/:id", handlers.GetStaticHandler(serv.static))
}

type server struct {
	e        *echo.Echo
	user     user.UserClient
	reading  reading.ReadingClient
	catalog  catalog.CatalogClient
	static   static.StaticClient
	profile  profile.ProfileClient
	activity activity.ActivityClient
}
