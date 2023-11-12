package main

import (
	"micromango/pkg/services/static"
	"os"
)

func main() {
	static.Run(os.Getenv("STATIC_ADDR"))
}
