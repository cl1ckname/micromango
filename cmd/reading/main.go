package main

import (
	"context"
	"log"
	"micromango/pkg/services/reading"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := reading.Config{
		Addr:   os.Getenv("READING_ADDR"),
		DbAddr: os.Getenv("READING_DB_ADDR"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	ok := reading.Run(ctx, c)

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
