package handlers

import (
	"fmt"
	"log"
	"net/http"

	"SocialNetworkRestApi/api/pkg/services"
)

type Application struct {
	Logger  *log.Logger
	Service *services.Service
}

func (app *Application) Home(rw http.ResponseWriter, r *http.Request) {

	// added cors headers in Authenticate middleware
	//utils.SetCors(&rw, r)

	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		app.Logger.Println("Cannot access homepage")
	}
}
