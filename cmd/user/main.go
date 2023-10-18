package main

import (
	"context"
	"log"
	"micromango/pkg/services/user"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := user.Config{
		Addr:      os.Getenv("USER_ADDR"),
		DbAddr:    os.Getenv("USER_DB_ADDR"),
		Salt:      os.Getenv("USER_SALT"),
		JwtSecret: os.Getenv("USER_JWT_SECRET"),
	}
	ctx, cancel := context.WithCancel(context.Background())
	ok := user.Run(ctx, c)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-ok:
		log.Fatal("Server stopped: ", err)
	case s := <-sigCh:
		log.Printf("got signal %v, attempting graceful shutdown", s)
		cancel()
	}
	log.Println("Graceful shutdown")

}
