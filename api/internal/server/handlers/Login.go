package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/models"
)

type JSONdata struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

func (app *Application) Login(rw http.ResponseWriter, r *http.Request) {

	utils.SetCors(&rw, r)

	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(rw, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		}
		r.Body = http.MaxBytesReader(rw, r.Body, 1024)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &JSONdata{}
		err := decoder.Decode(JSONdata)

		if err != nil {
			log.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		userData := &models.User{
			Email:    JSONdata.Email,
			Password: JSONdata.Password,
		}

		/*
			userData := &models.User{
				FirstName: "Test",
				LastName:  "User",
				Email:     "jarmo@test.ee",
				Password:  "kood1234",
				Birthday:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				Nickname:  "tester",
				About:     "I am a disco dancer",
				ImagePath: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
				IsPublic:  true,
			}
		*/

		sessionToken, err := app.Service.UserLogin(userData)
		if err != nil {
			log.Printf("Cannot login user: %s", err)
			http.Error(rw, "err", http.StatusUnauthorized)
			return
		}

		app.Service.SetCookie(rw, sessionToken)

		_, err = fmt.Fprintf(rw, "Successful login, cookie set")
		if err != nil {
			log.Printf("Cannot access login page: %s", err)
			http.Error(rw, "Cannot access login page", http.StatusInternalServerError)
			return
		}
	}
}
