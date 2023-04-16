package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/models"
)

type signinJSON struct {
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

		JSONdata := &signinJSON{}
		err := decoder.Decode(JSONdata)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		userData := &models.User{
			Email:    JSONdata.Email,
			Password: JSONdata.Password,
		}

		sessionToken, err := app.Service.UserLogin(userData)
		if err != nil {
			app.Logger.Printf("Cannot login user: %s", err)
			http.Error(rw, "err", http.StatusUnauthorized)
			return
		}

		app.Service.SetCookie(rw, sessionToken)

		_, err = fmt.Fprintf(rw, "Successful login, cookie set")
		if err != nil {
			app.Logger.Printf("Cannot access login page: %s", err)
			http.Error(rw, "Cannot access login page", http.StatusInternalServerError)
			return
		}
	}
}