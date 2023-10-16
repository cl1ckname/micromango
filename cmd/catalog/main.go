package main

import (
	"micromango/pkg/services/catalog"
	"os"
)

func main() {
	config := catalog.Config{
		Addr:   os.Getenv("CATALOG_ADDR"),
		DbAddr: os.Getenv("CATALOG_DB_ADDR"),
	}
	catalog.Run(config)
}
