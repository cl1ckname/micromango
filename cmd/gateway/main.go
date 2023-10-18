package main

import (
	"context"
	"log"
	"micromango/pkg/gateway"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := gateway.Config{
		Addr:        os.Getenv("GATEWAY_ADDR"),
		UserAddr:    os.Getenv("USER_ADDR"),
		CatalogAddr: os.Getenv("CATALOG_ADDR"),
		ReadingAddr: os.Getenv("READING_ADDR"),
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
