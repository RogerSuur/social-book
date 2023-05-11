package handlers

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"net/http"
)

type createPostJSON struct {
	UserId      int    `json:"userId"`
	Content     string `json:"content"`
	ImagePath   string `json:"imagePath"`
	PrivacyType int    `json:"privacyType"`
}

func (app *Application) Post(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		//Create a post method here
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &createPostJSON{}
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

		post := &models.Post{
			UserId:      userId,
			ImagePath:   JSONdata.ImagePath,
			Content:     JSONdata.Content,
			PrivacyType: enums.PrivacyType(JSONdata.PrivacyType),
		}

		err = app.PostService.CreatePost(post)

		if err != nil {
			app.Logger.Printf("Cannot create post: %s", err)
			http.Error(rw, "err", http.StatusBadRequest)
			return
		}

	default:
		http.Error(rw, "err", http.StatusBadRequest)
		return
	}

}

//Get handlers should return posts available to see for currently authenticated user
//It should be possible to get posts with offset fom database
