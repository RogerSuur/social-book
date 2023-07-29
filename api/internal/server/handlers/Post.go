package handlers

import (
	"SocialNetworkRestApi/api/pkg/enums"
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type createPostJSON struct {
	UserId      int      `json:"userId"`
	Content     string   `json:"content"`
	ImagePath   string   `json:"imagePath"`
	PrivacyType int      `json:"privacyType"`
	Receivers   []string `json:"selectedReceivers"`
}

func (app *Application) Post(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
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
			Receivers:   JSONdata.Receivers,
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

func (app *Application) GroupPost(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		vars := mux.Vars(r)

		groupIdStr := vars["groupId"]
		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)

		if groupId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &createPostJSON{}
		err = decoder.Decode(&JSONdata)

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
			PrivacyType: enums.PrivacyType(enums.None),
			Content:     JSONdata.Content,
			ImagePath:   JSONdata.ImagePath,
			GroupId:     groupId,
		}

		err = app.PostService.CreateGroupPost(post)

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
