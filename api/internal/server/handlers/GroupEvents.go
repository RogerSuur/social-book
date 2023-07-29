package handlers

import (
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) GroupEvents(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		groupIdStr := vars["groupId"]
		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)

		if groupId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		groupEvents, err := app.GroupEventService.GetGroupEvents(groupId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&groupEvents)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) CreateGroupEvent(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		//Create a post method here
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &models.CreateGroupEventFormData{}
		err := decoder.Decode(&JSONdata)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
			return
		}

		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Failed fetching user: %v", err)
			http.Error(rw, "Get user error", http.StatusBadRequest)
			return
		}

		notifications, err := app.GroupEventService.CreateGroupEvent(JSONdata, userId)

		if err != nil {
			http.Error(rw, "err", http.StatusBadRequest)
			return
		}

		err = app.WS.BroadcastGroupNotifications(notifications)

		if err != nil {
			http.Error(rw, "err", http.StatusBadRequest)
			return
		}

	default:
		http.Error(rw, "err", http.StatusBadRequest)
		return
	}

}

func (app *Application) Event(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		eventIdStr := vars["eventId"]
		eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

		if eventId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		event, err := app.GroupEventService.GetEventById(eventId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&event)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
