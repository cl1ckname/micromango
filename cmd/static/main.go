package main

import (
	"github.com/joho/godotenv"
	"log"
	"micromango/pkg/services/static"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	static.Run(os.Getenv("STATIC_ADDR"))
}
