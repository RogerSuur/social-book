package handlers

import (
	"SocialNetworkRestApi/api/internal/server/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) FeedPosts(w http.ResponseWriter, r *http.Request) {

	utils.SetCors(&w, r)

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		offset := vars["offset"]

		if len(offset) <= 0 {
			offset = "0"
		}

		feed, _ := app.Service.GetFeedPosts(offset)

		json.NewEncoder(w).Encode(&feed)
	// case "OPTIONS":
	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	default:
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

}
