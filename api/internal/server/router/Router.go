package router

import (
	"SocialNetworkRestApi/api/internal/server/handlers"
	"SocialNetworkRestApi/api/internal/server/utils"

	"github.com/gorilla/mux"
)

func New(app *handlers.Application) *mux.Router {
	r := mux.NewRouter()

	r.Use(utils.CorsMiddleware)

	r.HandleFunc("/", app.UserService.Authenticate(app.Home)).Methods("GET")
	r.HandleFunc("/ws", app.UserService.Authenticate(app.WS.WShandler))
	r.HandleFunc("/auth", app.UserService.Authenticate(nil)).Methods("GET")

	r.HandleFunc("/login", app.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/signup", app.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", app.UserService.Authenticate(app.Logout)).Methods("GET")
	r.HandleFunc("/profile", app.UserService.Authenticate(app.Profile)).Methods("GET")
	r.HandleFunc("/profile/{id:[0-9]+?}", app.UserService.Authenticate(app.Profile)).Methods("GET")
	r.HandleFunc("/profile/update", app.UserService.Authenticate(app.UpdateProfile)).Methods("POST", "OPTIONS")
	r.HandleFunc("/profile/update/avatar", app.UserService.Authenticate(app.UpdateImage)).Methods("POST", "OPTIONS")
	r.HandleFunc("/following", app.UserService.Authenticate(app.Following)).Methods("GET")
	r.HandleFunc("/followers", app.UserService.Authenticate(app.Followers)).Methods("GET")
	r.HandleFunc("/feedposts/{offset}", app.UserService.Authenticate(app.FeedPosts)).Methods("GET")
	r.HandleFunc("/comments/{postId}/{offset}", app.UserService.Authenticate(app.Comments)).Methods("GET")
	r.HandleFunc("/post", app.UserService.Authenticate(app.Post)).Methods("POST")

	return r
}
