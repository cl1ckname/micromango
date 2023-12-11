package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/share"
	"micromango/pkg/grpc/user"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	addr := os.Getenv("SERVER_ADDR")
	n := flag.Int("n", 0, "Number of users to register")
	flag.Parse()

	log.Println("retrieving manga")
	mangasRsp, err := http.Get(addr + "/api/catalog")
	if err != nil {
		log.Fatal(err)
	}
	var mangas []*share.MangaPreviewResponse
	if err := json.NewDecoder(mangasRsp.Body).Decode(&mangas); err != nil {
		log.Fatal(err)
	}
	log.Printf("got %d mangas\n", len(mangas))

	log.Printf("registering %d users\n", *n)
	var users []*user.UserResponse
	for i := 0; i < *n; i++ {
		payload, _ := json.Marshal(user.RegisterRequest{
			Username: gofakeit.Name(),
			Email:    gofakeit.Email(),
			Password: "qwe",
		})
		payloadReader := bytes.NewReader(payload)
		regResp, err := http.Post(addr+"/api/user/register", "application/json", payloadReader)
		if err != nil {
			log.Fatal(err)
		}
		if regResp.StatusCode != 201 {
			log.Fatal("invalid register user resp code: ", regResp.StatusCode)
		}
		var resp user.UserResponse
		if err := json.NewDecoder(regResp.Body).Decode(&resp); err != nil {
			log.Fatal(err)
		}
		users = append(users, &resp)
	}
	log.Printf("registered %d users\n", len(users))

	log.Println("generating random lists data")
	for _, u := range users {
		for _, m := range mangas {
			var req profile.AddToListRequest
			req.MangaId = m.MangaId
			req.List = share.ListName(1 + rand.Intn(5))
			reqBytes, _ := json.Marshal(req)
			reqReader := bytes.NewReader(reqBytes)
			resp, err := http.Post(addr+"/api/profile/"+u.UserId+"/list", "application/json", reqReader)
			if err != nil {
				log.Fatal(err)
			}
			if resp.StatusCode != 200 && resp.StatusCode != 201 {
				log.Fatal("invalid status code: ", resp.StatusCode)
			}
		}
	}
}
