package handlers

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"net/http"
)

type postJSON struct {
	UserId      int    `json:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PrivacyType int    `json:"privacyType"`
}

func (app *Application) Post(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		//Create a post method here
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &postJSON{}
		err := decoder.Decode(JSONdata)

		if err != nil {
			app.Logger.Printf("JSON error: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		post := &models.Post{
			UserId:      JSONdata.UserId,
			Title:       JSONdata.Title,
			Content:     JSONdata.Content,
			PrivacyType: enums.PrivacyType(JSONdata.PrivacyType),
		}

		err = app.Service.CreatePost(post)

		if err != nil {
			app.Logger.Printf("Cannot create post: %s", err)
			http.Error(rw, "err", http.StatusBadRequest)
			return
		}
	}
}

//Get handlers should return posts available to see for currently authenticated user
//It should be possible to get posts with offset fom database
