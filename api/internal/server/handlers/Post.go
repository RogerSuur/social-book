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
			return
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

func (app *Application) PostImage(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		vars := mux.Vars(r)
		id := vars["postId"]

		postId, err := strconv.ParseInt(id, 10, 64)
		if postId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		// Limit the size of the request body to 5MB
		//app.Logger.Printf("Request body size: %d", r.ContentLength)
		r.Body = http.MaxBytesReader(rw, r.Body, 20<<18+512)

		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Failed fetching user: %v", err)
			http.Error(rw, "Get user error", http.StatusUnauthorized)
		}

		err = r.ParseMultipartForm(20 << 18)

		if err != nil {
			app.Logger.Printf("Failed parsing form: %v", err)
			http.Error(rw, "Parsing form error", http.StatusRequestEntityTooLarge)
		}

		file, header, err := r.FormFile("image")

		if err != nil {
			app.Logger.Printf("Failed getting file: %v", err)
			http.Error(rw, "Get file error", http.StatusUnsupportedMediaType)
		}

		defer file.Close()

		err = app.PostService.UpdatePostImage(userId, postId, file, header)

		if err != nil {
			app.Logger.Printf("Cannot upload image: %s", err)
			http.Error(rw, "err", http.StatusBadRequest)
			return
		}

		resp := make(map[string]interface{})
		resp["message"] = "Image uploaded successfully"
		resp["status"] = http.StatusOK
		json.NewEncoder(rw).Encode(resp)

	default:
		http.Error(rw, "err", http.StatusBadRequest)
		return
	}

}
