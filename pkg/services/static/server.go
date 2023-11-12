package static

import (
	"github.com/labstack/echo/v4"
	"log"
)

func Run(addr string) {
	e := echo.New()
	e.File("/", "static/index.html")
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
