package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/models"
)

type signupJSON struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	// AvatarImage string `json:"avatarImage"`
	Nickname string `json:"nickname"`
	About    string `json:"about"`
}

func (app *Application) Register(rw http.ResponseWriter, r *http.Request) {

	utils.SetCors(&rw, r)

	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(rw, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		}
		r.Body = http.MaxBytesReader(rw, r.Body, 1024)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &signupJSON{}
		err := decoder.Decode(JSONdata)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		userData := &models.User{
			Email:     JSONdata.Email,
			Password:  JSONdata.Password,
			FirstName: JSONdata.FirstName,
			LastName:  JSONdata.LastName,
			Nickname:  JSONdata.Nickname,
			About:     JSONdata.About,
		}

		sessionToken, err := app.Service.UserRegister(userData)
		if err != nil {
			app.Logger.Printf("Cannot register user: %s", err)
			http.Error(rw, "err", http.StatusUnauthorized)
			return
		}

		app.Service.SetCookie(rw, sessionToken)

		_, err = fmt.Fprintf(rw, "Successful registration")
		if err != nil {
			app.Logger.Printf("Cannot access register page: %s", err)
			http.Error(rw, "Cannot access register page", http.StatusInternalServerError)
			return
		}
	}
}
