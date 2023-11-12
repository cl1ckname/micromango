package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"micromango/pkg/gateway"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	c := gateway.Config{
		Addr:        os.Getenv("GATEWAY_ADDR"),
		UserAddr:    os.Getenv("USER_ADDR"),
		CatalogAddr: os.Getenv("CATALOG_ADDR"),
		ReadingAddr: os.Getenv("READING_ADDR"),
		StaticAddr:  os.Getenv("STATIC_ADDR"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	ok := gateway.Run(ctx, c)

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
