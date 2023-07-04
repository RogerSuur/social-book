package handlers

import (
	"encoding/json"
	"net/http"
)

// Return groups that the user is a member of
func (app *Application) UserGroups(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		groups, err := app.GroupService.GetUserGroups(userId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&groups)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) MyGroups(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		groups, err := app.GroupService.GetUserCreatedGroups(userId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&groups)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
