package main

import (
	"micromango/pkg/services/reading"
	"os"
)

func main() {
	c := reading.Config{
		Addr:   os.Getenv("READING_ADDR"),
		DbAddr: os.Getenv("READING_DB_ADDR"),
	}
	reading.Run(c)
}
