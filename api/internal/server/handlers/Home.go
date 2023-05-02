package handlers

import (
	"SocialNetworkRestApi/api/pkg/services"
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	Logger         *log.Logger
	UserService    services.IUserService
	PostService    services.IPostService
	CommentService services.ICommentService
}

// func InitApplication(repositories *models.Repositories) *Application {
// 	return &Application{
// 		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
// 		UserService: services.InitUserService(
// 			repositories.UserRepo,
// 			repositories.SessionRepo,
// 			repositories.FollowerRepo,
// 		),
// 		PostService:    services.InitPostService(repositories.PostRepo),
// 		CommentService: services.InitCommentService(repositories.CommentRepo),
// 	}
// }

func (app *Application) Home(rw http.ResponseWriter, r *http.Request) {

	// added cors headers in Authenticate middleware
	//utils.SetCors(&rw, r)
	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		app.Logger.Println("Cannot access homepage")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
