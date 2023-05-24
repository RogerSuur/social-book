package handlers

import (
	"SocialNetworkRestApi/api/internal/server/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) Comments(rw http.ResponseWriter, r *http.Request) {
	utils.SetCors(&rw, r)

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		//Get postId from endpoint and parse
		postIdStr := vars["postId"]
		postId, err := strconv.Atoi(postIdStr)

		fmt.Println("postIdStr", postIdStr)

		if postId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		//Get offset from endpoint and parse
		offsetStr := vars["offset"]
		offset, err := strconv.Atoi(offsetStr)

		if offset < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		comments, err := app.CommentService.GetPostComments(postId, offset)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&comments)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
