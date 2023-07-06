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
	//Session
	r.HandleFunc("/auth", app.UserService.Authenticate(nil)).Methods("GET")
	r.HandleFunc("/login", app.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/signup", app.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", app.UserService.Authenticate(app.Logout)).Methods("GET")
	//Profile
	r.HandleFunc("/profile", app.UserService.Authenticate(app.Profile)).Methods("GET")
	r.HandleFunc("/profile/{id:[0-9]+?}", app.UserService.Authenticate(app.Profile)).Methods("GET")
	r.HandleFunc("/profile/update", app.UserService.Authenticate(app.UpdateProfile)).Methods("POST", "OPTIONS")
	r.HandleFunc("/profile/update/avatar", app.UserService.Authenticate(app.UpdateUserImage)).Methods("POST", "OPTIONS")
	r.HandleFunc("/following", app.UserService.Authenticate(app.Following)).Methods("GET")
	r.HandleFunc("/followers", app.UserService.Authenticate(app.Followers)).Methods("GET")
	//Posts
	r.HandleFunc("/feedposts/{offset}", app.UserService.Authenticate(app.FeedPosts)).Methods("GET")
	r.HandleFunc("/comments/{postId}/{offset}", app.UserService.Authenticate(app.Comments)).Methods("GET")
	r.HandleFunc("/insertcomment", app.UserService.Authenticate(app.Comment)).Methods("POST")
	r.HandleFunc("/post", app.UserService.Authenticate(app.Post)).Methods("POST")
	r.HandleFunc("/profileposts/{offset}", app.UserService.Authenticate(app.ProfilePosts)).Methods("GET")
	//Groups
	r.HandleFunc("/usergroups", app.UserService.Authenticate(app.UserGroups)).Methods("GET")
	r.HandleFunc("/mygroups", app.UserService.Authenticate(app.MyGroups)).Methods("GET")
	r.HandleFunc("/groups/{groupId}", app.UserService.Authenticate(app.Group)).Methods("GET")
	r.HandleFunc("/groupmembers/{groupId}", app.UserService.Authenticate(app.GroupMembers)).Methods("GET")

	// ACCEPTED_EVENTS_URL = "http://localhost:8000/userevents";
	// GROUP_USERS_URL = "http://localhost:8000/groupUsers/{groupId}";
	// SEARCH_URL = "http://localhost:8000/search/";
	return r
}
