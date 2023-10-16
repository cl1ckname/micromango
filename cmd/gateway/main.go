package main

import (
	"micromango/pkg/gateway"
	"os"
)

func main() {
	c := gateway.Config{
		Addr:        os.Getenv("GATEWAY_ADDR"),
		UserAddr:    os.Getenv("USER_ADDR"),
		CatalogAddr: os.Getenv("CATALOG_ADDR"),
		ReadingAddr: os.Getenv("READING_ADDR"),
	}
	gateway.Run(c)
}
