package handlers

import (
	"fmt"
	"log"
	"net/http"

	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/services"
)

type Application struct {
	Logger  *log.Logger
	Service *services.Service
}

func (app *Application) Home(rw http.ResponseWriter, r *http.Request) {

	utils.SetCors(&rw, r)

	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		log.Println("Cannot access homepage")
	}
}
