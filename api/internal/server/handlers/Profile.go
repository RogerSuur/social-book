package handlers

import (
	"encoding/json"
	"net/http"
)

func (app *Application) Profile(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//app.Logger.Println("Profile page accessed")

		userID, err := app.UserService.GetUserID(r)
		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, "Cannot get user ID", http.StatusUnauthorized)
			return
		}

		userData, err := app.UserService.GetUserData(int64(userID))
		if err != nil {
			app.Logger.Printf("Cannot get user data: %s", err)
			http.Error(rw, "Cannot get user data", http.StatusUnauthorized)
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
