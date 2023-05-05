package router

import (
	"SocialNetworkRestApi/api/internal/server/handlers"

	"github.com/gorilla/mux"
)

func New(app *handlers.Application) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", app.UserService.Authenticate(app.Home)).Methods("GET")
	r.HandleFunc("/login", app.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/signup", app.Register).Methods("POST", "OPTIONS")
	//r.HandleFunc("/logout", app.UserService.Authenticate(app.Logout))
	r.HandleFunc("/profile", app.UserService.Authenticate(app.Profile)).Methods("GET")
	r.HandleFunc("/profile/update", app.UserService.Authenticate(app.UpdateProfile)).Methods("POST")
	r.HandleFunc("/profile/update/avatar", app.UserService.Authenticate(app.UpdateImage)).Methods("POST")
	r.HandleFunc("/following", app.UserService.Authenticate(app.Following)).Methods("GET")
	r.HandleFunc("/followers", app.UserService.Authenticate(app.Followers)).Methods("GET")

	return r
}
