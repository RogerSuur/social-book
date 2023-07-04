package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

		groups, err := app.GroupService.GetUserGroups(int(userId))

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

		groups, err := app.GroupService.GetUserCreatedGroups(int(userId))

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

func (app *Application) Group(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		groupIdStr := vars["groupId"]
		groupId, err := strconv.Atoi(groupIdStr)

		if groupId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		group, err := app.GroupService.GetGroupById(groupId)

		if err != nil {
			app.Logger.Printf("Failed fetching group: %v", err)
			http.Error(rw, "Fetch error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&group)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
