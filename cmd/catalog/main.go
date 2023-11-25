package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"micromango/pkg/services/catalog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := catalog.Config{
		Addr:               os.Getenv("CATALOG_ADDR"),
		DbAddr:             os.Getenv("CATALOG_DB_ADDR"),
		ReadingServiceAddr: os.Getenv("READING_ADDR"),
		StaticServiceAddr:  os.Getenv("STATIC_ADDR"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	ok := catalog.Run(ctx, config)

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
