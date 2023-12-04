package gateway

import (
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"micromango/pkg/gateway/handlers"
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
	apiGroup := e.Group("api")

	handlers.RegisterUser(apiGroup, serv.user)
	handlers.RegisterCatalog(apiGroup, serv.catalog)
	handlers.RegisterProfile(apiGroup, serv.profile)
	handlers.RegisterReading(apiGroup, serv.reading)
	e.GET("static/:id", handlers.GetStaticHandler(serv.static))
}

type server struct {
	e       *echo.Echo
	user    user.UserClient
	reading reading.ReadingClient
	catalog catalog.CatalogClient
	static  static.StaticClient
	profile profile.ProfileClient
}
