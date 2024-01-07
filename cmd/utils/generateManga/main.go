package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"io"
	"log"
	"math/rand"
	"micromango/pkg/common/utils"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	addr := os.Getenv("SERVER_ADDR")
	n := flag.Int("n", 0, "Number of users to register")
	email := flag.String("email", "click@name.me", "Email to get auth")
	password := flag.String("password", "qwe", "Password to get auth")
	flag.Parse()

	coverBytes, err := os.ReadFile("cmd/utils/generateManga/cover.jpg")
	if err != nil {
		log.Fatal(err)
	}

	token := login(*email, *password, addr)

	for i := 0; i < *n; i++ {
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		fw, err := w.CreateFormFile("cover", "cover.jpg")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := fw.Write(coverBytes); err != nil {
			log.Fatal(err)
		}
		title := gofakeit.BookTitle()
		fw, err = w.CreateFormField("title")
		if _, err := fw.Write([]byte(title)); err != nil {
			log.Fatal(err)
		}
		description := gofakeit.ProductDescription()
		fw, err = w.CreateFormField("description")
		if _, err := fw.Write([]byte(description)); err != nil {
			log.Fatal(err)
		}
		if err := w.Close(); err != nil {
			log.Fatal(err)
		}

		var genres []int
		for i := 0; i < rand.Intn(20); i++ {
			sp := rand.Intn(20)
			g := (sp + i) % 20
			genres = append(genres, g)
		}
		genresStr := strings.Join(
			utils.Map(genres, func(i int) string {
				return strconv.FormatInt(int64(i), 10)
			}), ",")

		fw, err = w.CreateFormField("genres")
		if _, err := fw.Write([]byte(genresStr)); err != nil {
			log.Fatal(err)
		}
		if err := w.Close(); err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", addr+"/api/catalog", &b)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		res, err := (&http.Client{}).Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != 200 && res.StatusCode != 201 {
			var respData struct{ Message string }
			if err := json.NewDecoder(res.Body).Decode(&respData); err != nil {
				log.Fatal(err)
			}
			log.Fatal("rate invalid status code: ", res.StatusCode, respData.Message)
		}
	}
}

func login(email string, password string, addr string) string {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{email, password}
	payloadBytes, _ := json.Marshal(payload)
	payloadReader := bytes.NewReader(payloadBytes)
	authReq, err := http.NewRequest("POST", addr+"/api/user/login", payloadReader)
	if err != nil {
		log.Fatal(err)
	}
	authReq.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(authReq)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("body: ", string(bodyBytes))
	if resp.StatusCode != 201 {
		log.Fatal("status: ", resp.StatusCode)
	}

	var authResp struct {
		AccessToken string `json:"accessToken"`
	}
	err = json.Unmarshal(bodyBytes, &authResp)
	if err != nil {
		log.Fatal(err)
	}
	token := authResp.AccessToken
	return token
}
