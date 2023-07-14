package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) UserEvents(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("Userevents GET")
		userID, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Failed fetching user: %v", err)
			http.Error(rw, "Get user error", http.StatusBadRequest)
		}

		events, err := app.GroupEventService.GetUserEvents(userID)

		fmt.Println("events", events)
		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&events)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}
}
