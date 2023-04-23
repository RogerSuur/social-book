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

		// parse userdata to JSON and send it to client
		JSONdata := &profileJSON{
			FirstName:   userData.FirstName,
			LastName:    userData.LastName,
			Email:       userData.Email,
			Birthday:    userData.Birthday.Format("02/01/2006"),
			Nickname:    userData.Nickname,
			About:       userData.About,
			AvatarImage: userData.ImagePath,
			CreatedAt:   userData.CreatedAt.Format("02/01/2006 15:04:05"),
			IsPublic:    userData.IsPublic,
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		resp := make(map[string]interface{})
		resp["user"] = JSONdata
		resp["message"] = "User profile data retrieved"
		resp["status"] = "success"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			app.Logger.Printf("Cannot marshal JSON: %s", err)
			http.Error(rw, "Cannot marshal JSON", http.StatusInternalServerError)
			return
		}
		rw.Write(jsonResp)

	}
}
