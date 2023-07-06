package handlers

import (
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
