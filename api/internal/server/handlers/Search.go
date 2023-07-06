package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) Search(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		searchString := vars["searchcriteria"]
		app.Logger.Println(searchString)

		groupSearchResult, err := app.GroupService.SearchGroupsAndUsers(searchString)

		app.Logger.Println(groupSearchResult)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&groupSearchResult)
	// case "OPTIONS":
	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
