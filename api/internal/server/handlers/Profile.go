package handlers

import (
	"SocialNetworkRestApi/api/pkg/services"
	"encoding/json"
	"net/http"
)

func (app *Application) Profile(rw http.ResponseWriter, r *http.Request) {

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

func (app *Application) UpdateProfile(rw http.ResponseWriter, r *http.Request) {

	userID, err := app.UserService.GetUserID(r)
	if err != nil {
		app.Logger.Printf("Cannot get user ID: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(rw, r.Body, 1024)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	updateData := &services.ProfileUpdateJSON{}
	err = decoder.Decode(updateData)

	if err != nil {
		app.Logger.Printf("Cannot decode JSON: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = app.UserService.UpdateUserData(int64(userID), *updateData)
	if err != nil {
		app.Logger.Printf("Cannot update user data: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	resp := make(map[string]interface{})
	resp["message"] = "User profile data updated"
	resp["status"] = "success"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		app.Logger.Printf("Cannot marshal JSON: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResp)
}

func (app *Application) Followers(rw http.ResponseWriter, r *http.Request) {

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
}

func (app *Application) Following(rw http.ResponseWriter, r *http.Request) {

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
}
