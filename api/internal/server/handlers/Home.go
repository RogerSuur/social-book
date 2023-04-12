package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
)

type Application struct {
	Logger  *log.Logger
	Service *services.Service
}

func (app *Application) Home(rw http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		log.Println("Cannot access homepage")
	}
}

func (app *Application) Login(rw http.ResponseWriter, r *http.Request) {

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

	sessionToken, err := app.Service.UserLogin(userData)
	if err != nil {
		log.Printf("Cannot login user: %s", err)
		http.Error(rw, "Cannot login user", http.StatusInternalServerError)
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
