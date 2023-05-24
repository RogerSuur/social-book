package handlers

import (
	"SocialNetworkRestApi/api/internal/server/utils"
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type insertCommentJSON struct {
	PostId    int    `json:"postId"`
	Content   string `json:"content"`
	ImagePath string `json:"imagePath"`
}

func (app *Application) Comments(rw http.ResponseWriter, r *http.Request) {
	utils.SetCors(&rw, r)

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		//Get postId from endpoint and parse
		postIdStr := vars["postId"]
		postId, err := strconv.Atoi(postIdStr)

		// fmt.Println("postIdStr", postIdStr)

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

func (app *Application) Comment(rw http.ResponseWriter, r *http.Request) {
	utils.SetCors(&rw, r)

	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &insertCommentJSON{}
		err := decoder.Decode(&JSONdata)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Failed fetching user: %v", err)
			http.Error(rw, "Get user error", http.StatusBadRequest)
		}

		comment := &models.Comment{
			PostId:    JSONdata.PostId,
			UserId:    userId,
			Content:   JSONdata.Content,
			ImagePath: JSONdata.ImagePath,
		}

		err = app.CommentService.CreateComment(comment)

		if err != nil {
			app.Logger.Printf("Creating comment failed: %v", err)
			http.Error(rw, "Error", http.StatusBadRequest)
		}

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
