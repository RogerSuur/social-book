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
	r.HandleFunc("/following", app.UserService.Authenticate(app.Following)).Methods("GET")
	r.HandleFunc("/followers", app.UserService.Authenticate(app.Followers)).Methods("GET")
	r.HandleFunc("/feedposts/{offset}", app.UserService.Authenticate(app.FeedPosts)).Methods("GET")
	r.HandleFunc("/comments/{postId}", app.UserService.Authenticate(app.Comments)).Methods("GET")
	r.HandleFunc("/post", app.UserService.Authenticate(app.Post)).Methods("POST")

	return r
}
