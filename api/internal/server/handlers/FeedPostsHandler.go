package handlers

import (
	"SocialNetworkRestApi/api/internal/server/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) FeedPosts(rw http.ResponseWriter, r *http.Request) {
	utils.SetCors(&rw, r)

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		offset := vars["offset"]

		if len(offset) <= 0 {
			offset = "0"
		}

		feed, err := app.PostService.GetFeedPosts(offset)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&feed)
	// case "OPTIONS":
	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
