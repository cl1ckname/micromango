package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"micromango/pkg/grpc/activity"
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
	var users []struct {
		UserId string
		Email  string
	}
	for i := 0; i < *n; i++ {
		email := gofakeit.Email()
		payload, _ := json.Marshal(user.RegisterRequest{
			Username: gofakeit.Name(),
			Email:    email,
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
		users = append(users, struct {
			UserId string
			Email  string
		}{UserId: resp.UserId, Email: email})
	}
	log.Printf("registered %d users\n", len(users))

	log.Println("generating random lists data")
	for _, u := range users {
		loginReq := user.LoginRequest{
			Email:    u.Email,
			Password: "qwe",
		}
		loginBytes, _ := json.Marshal(&loginReq)
		loginReader := bytes.NewReader(loginBytes)
		loginResp, err := http.Post(addr+"/api/user/login", "application/json", loginReader)
		if err != nil {
			log.Fatal(err)
		}
		var loginData user.LoginResponse
		if err := json.NewDecoder(loginResp.Body).Decode(&loginData); err != nil {
			log.Fatal(err)
		}

		for _, m := range mangas {

			var req profile.AddToListRequest
			req.MangaId = m.MangaId
			req.List = share.ListName(1 + rand.Intn(5))
			reqBytes, _ := json.Marshal(&req)
			reqReader := bytes.NewReader(reqBytes)
			httpReq, _ := http.NewRequest("POST", addr+"/api/profile/"+u.UserId+"/list", reqReader)
			httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginData.AccessToken))
			httpReq.Header.Set("Content-Type", "application/json")
			resp, err := (&http.Client{}).Do(httpReq)
			if err != nil {
				log.Fatal(err)
			}
			if resp.StatusCode != 200 && resp.StatusCode != 201 {
				var respData struct{ Message string }
				json.NewDecoder(resp.Body).Decode(&respData)
				log.Fatal("list invalid status code: ", resp.StatusCode, respData.Message)
			}

			if rand.Int()%2 == 0 {
				req := activity.LikeRequest{
					MangaId: m.MangaId,
					UserId:  u.UserId,
				}
				reqBytes, _ := json.Marshal(req)
				reqReader := bytes.NewReader(reqBytes)
				httpReq, _ := http.NewRequest("POST", addr+"/api/activity/manga/"+m.MangaId+"/like", reqReader)
				httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginData.AccessToken))
				httpReq.Header.Set("Content-Type", "application/json")
				resp, err := (&http.Client{}).Do(httpReq)
				if err != nil {
					log.Fatal(err)
				}
				if resp.StatusCode != 200 && resp.StatusCode != 201 {
					var respData struct{ Message string }
					json.NewDecoder(resp.Body).Decode(&respData)
					log.Fatal("like invalid status code: ", resp.StatusCode, respData.Message)
				}
			}

			if rate := rand.Intn(11); rate != 0 {
				req := activity.RateMangaRequest{
					MangaId: m.MangaId,
					UserId:  u.UserId,
					Rate:    uint32(rate),
				}
				reqBytes, _ := json.Marshal(req)
				reqReader := bytes.NewReader(reqBytes)
				httpReq, _ := http.NewRequest("POST", addr+"/api/activity/manga/"+m.MangaId+"/rate", reqReader)
				httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginData.AccessToken))
				httpReq.Header.Set("Content-Type", "application/json")
				resp, err := (&http.Client{}).Do(httpReq)
				if err != nil {
					log.Fatal(err)
				}
				if resp.StatusCode != 200 && resp.StatusCode != 201 {
					var respData struct{ Message string }
					json.NewDecoder(resp.Body).Decode(&respData)
					log.Fatal("rate invalid status code: ", resp.StatusCode, respData.Message)
				}
			}
		}
	}
}
