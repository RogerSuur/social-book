package handlers

import (
	"SocialNetworkRestApi/api/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Return groups that the user is a member of
func (app *Application) UserGroups(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		groups, err := app.GroupService.GetUserGroups(userId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&groups)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) MyGroups(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		groups, err := app.GroupService.GetUserCreatedGroups(userId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&groups)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) Group(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		groupIdStr := vars["groupId"]
		app.Logger.Printf("Decoding group ID provided in the URL (group %v) for Group handler", groupIdStr)
		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)

		if groupId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		group, err := app.GroupService.GetGroupById(groupId)

		if err != nil {
			app.Logger.Printf("Failed fetching group: %v", err)
			http.Error(rw, "Fetch error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&group)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) GroupMembers(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)

		groupIdStr := vars["groupId"]
		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)

		if groupId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		isGroupMember, err := app.GroupMemberService.IsGroupMember(groupId, userId)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if !isGroupMember {
			app.Logger.Printf("User %d is not a member of this group", userId)
			http.Error(rw, "Not a member of this group", http.StatusForbidden)
			return
		}

		members, err := app.GroupMemberService.GetGroupMembers(groupId)

		if err != nil {
			app.Logger.Printf("Failed fetching groups: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&members)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) CreateGroup(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		//Create a post method here
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &models.GroupJSON{}
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

		result, err := app.GroupService.CreateGroup(JSONdata, userId)

		if err != nil {
			http.Error(rw, "err", http.StatusBadRequest)
			return
		}

		app.Logger.Printf("Group with id %d created successfully", result)
		rw.WriteHeader(http.StatusCreated)

	default:
		http.Error(rw, "err", http.StatusBadRequest)
		return
	}

}

func (app *Application) GroupPosts(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		offset := vars["offset"]
		offsetInt, err := strconv.ParseInt(offset, 10, 64)

		if offsetInt < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		groupIdStr := vars["groupId"]

		groupId, err := strconv.ParseInt(groupIdStr, 10, 64)

		if groupId < 0 || err != nil {
			app.Logger.Printf("DATA PARSE error: %v", err)
			http.Error(rw, "DATA PARSE error", http.StatusBadRequest)
		}

		userId, err := app.UserService.GetUserID(r)

		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		isGroupMember, err := app.GroupMemberService.IsGroupMember(groupId, userId)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if !isGroupMember {
			app.Logger.Printf("User %d is not a member of this group", userId)
			http.Error(rw, "Not a member of this group", http.StatusForbidden)
			return
		}

		feed, err := app.PostService.GetGroupPosts(groupId, offsetInt)

		if err != nil {
			app.Logger.Printf("Failed fetching posts: %v", err)
			http.Error(rw, "JSON error", http.StatusBadRequest)
		}

		json.NewEncoder(rw).Encode(&feed)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}
}

func (app *Application) AddMembers(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		JSONdata := &models.GroupMemberJSON{}
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

		result, err := app.GroupMemberService.AddMembers(userId, *JSONdata)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		app.Logger.Printf("Members added successfully")
		app.Logger.Printf("Members added: %v", result)

		json.NewEncoder(rw).Encode(&result)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}

func (app *Application) UpdateGroupImage(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		// Limit the size of the request body to 5MB
		r.Body = http.MaxBytesReader(rw, r.Body, 20<<18)

		vars := mux.Vars(r)
		groupId := vars["groupId"]
		groupIdInt, err := strconv.ParseInt(groupId, 10, 64)
		if err != nil {
			app.Logger.Printf("Cannot parse group ID: %s", err)
			http.Error(rw, "Cannot parse group ID", http.StatusBadRequest)
			return
		}

		userID, err := app.UserService.GetUserID(r)
		if err != nil {
			app.Logger.Printf("Cannot get user ID: %s", err)
			http.Error(rw, "Cannot get user ID", http.StatusUnauthorized)
			return
		}

		err = r.ParseMultipartForm(20 << 18)
		if err != nil {
			app.Logger.Printf("Cannot parse multipart form: %s", err)
			http.Error(rw, err.Error(), http.StatusRequestEntityTooLarge)
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			app.Logger.Printf("Cannot get image file: %s", err)
			http.Error(rw, err.Error(), http.StatusUnsupportedMediaType)
			return
		}
		defer file.Close()

		err = app.GroupService.UpdateGroupImage(userID, groupIdInt, file, header)
		if err != nil {
			app.Logger.Printf("Cannot update user image: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		resp := make(map[string]interface{})
		resp["message"] = "User image updated"
		resp["status"] = "success"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			app.Logger.Printf("Cannot marshal JSON: %s", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(jsonResp)

	default:
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}

}
