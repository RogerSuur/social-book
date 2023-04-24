package handlers

import (
	"encoding/json"
	"net/http"
)

type profileJSON struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Birthday    string `json:"birthday"`
	Nickname    string `json:"nickname"`
	About       string `json:"about"`
	AvatarImage string `json:"avatarImage"`
	CreatedAt   string `json:"createdAt"`
	IsPublic    bool   `json:"isPublic"`
}

type follower struct {
	UserID      int    `json:"userID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Nickname    string `json:"nickname"`
	AvatarImage string `json:"avatarImage"`
	Accepted    bool   `json:"accepted"`
}

type userFollowers struct {
	Followers []follower `json:"followers"`
}

type userFollowing struct {
	Following []follower `json:"following"`
}

func (app *Application) Profile(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//app.Logger.Println("Profile page accessed")

		userID, err := app.UserService.GetUserID(r)
		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		userData, err := app.UserService.GetUserData(int64(userID))
		if err != nil {
			app.Logger.Printf("Cannot get user data: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		resp := make(map[string]interface{})
		resp["user"] = userData
		resp["message"] = "User profile data retrieved"
		resp["status"] = "success"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			app.Logger.Printf("Cannot marshal JSON: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(jsonResp)

	}
}

func (app *Application) Followers(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		userID, err := app.UserService.GetUserID(r)
		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, "Cannot get user ID", http.StatusUnauthorized)
			return
		}

		userFollowers, err := app.UserService.GetUserFollowers(int64(userID))

		if err != nil {
			app.Logger.Printf("Cannot get user followers: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(rw).Encode(userFollowers)

		if err != nil {
			app.Logger.Printf("Cannot encode user followers: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}
}

func (app *Application) Following(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		userID, err := app.UserService.GetUserID(r)
		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, "Cannot get user ID", http.StatusUnauthorized)
			return
		}

		userFollowing, err := app.UserService.GetUserFollowing(int64(userID))

		if err != nil {
			app.Logger.Printf("Cannot get user following: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(rw).Encode(userFollowing)

		if err != nil {
			app.Logger.Printf("Cannot encode user following: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}
}
