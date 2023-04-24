package handlers

import (
	"SocialNetworkRestApi/api/pkg/services"
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	Logger      *log.Logger
	UserService services.IUserService
}

func (app *Application) Home(rw http.ResponseWriter, r *http.Request) {

	// added cors headers in Authenticate middleware
	//utils.SetCors(&rw, r)
	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		app.Logger.Println("Cannot access homepage")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
