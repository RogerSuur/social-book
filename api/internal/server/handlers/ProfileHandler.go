package handlers

import (
	"SocialNetworkRestApi/api/pkg/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) Profile(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var profileId int64

	requestingUserId, err := app.UserService.GetUserID(r)
	if err != nil {
		app.Logger.Printf("Cannot get user ID: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if id == "" {
		// self
		profileId = requestingUserId
	} else {
		app.Logger.Printf("Decoding user ID provided in the URL (user %v) for Profile handler", id)
		profileId, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			app.Logger.Printf("Cannot parse user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	profileData, err := app.UserService.GetUserData(requestingUserId, profileId)
	if err != nil {
		app.Logger.Printf("Cannot get user data: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	resp := make(map[string]interface{})
	resp["user"] = profileData
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

	updateData := &services.ProfileJSON{}
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

	app.Logger.Printf("User profile data updated")

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

func (app *Application) UpdateUserImage(rw http.ResponseWriter, r *http.Request) {

	// Limit the size of the request body to 5MB
	app.Logger.Printf("Request body size: %d", r.ContentLength)
	r.Body = http.MaxBytesReader(rw, r.Body, 20<<18+512)

	userID, err := app.UserService.GetUserID(r)
	if err != nil {
		app.Logger.Printf("Cannot get user ID: %s", err)
		http.Error(rw, "Cannot get user ID", http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(20 << 18)
	if err != nil {
		app.Logger.Printf("Cannot parse multipart form: %s", err)
		http.Error(rw, err.Error(), http.StatusRequestEntityTooLarge)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		app.Logger.Printf("Cannot get image file: %s", err)
		http.Error(rw, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	defer file.Close()

	err = app.UserService.UpdateUserImage(userID, file, header)
	if err != nil {
		app.Logger.Printf("Cannot update user image: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	resp := make(map[string]interface{})
	resp["message"] = "User image updated"
	resp["status"] = "success"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		app.Logger.Printf("Cannot marshal JSON: %s", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write(jsonResp)
}
