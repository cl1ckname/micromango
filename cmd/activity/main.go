package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"micromango/pkg/services/activity/app"
	"micromango/pkg/services/activity/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := config.Config{
		Addr:        os.Getenv("ACTIVITY_ADDR"),
		DbAddr:      os.Getenv("ACTIVITY_DB_ADDR"),
		CatalogAddr: os.Getenv("CATALOG_ADDR"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	ok := app.Run(ctx, cfg)

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
