package main

import (
	"micromango/pkg/services/user"
	"os"
)

func main() {
	c := user.Config{
		Addr:      os.Getenv("USER_ADDR"),
		DbAddr:    os.Getenv("USER_DB_ADDR"),
		Salt:      os.Getenv("USER_SALT"),
		JwtSecret: os.Getenv("USER_JWT_SECRET"),
	}
	user.Run(c)
}
