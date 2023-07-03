package handlers

import (
	"SocialNetworkRestApi/api/internal/server/websocket"
	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	Logger              *log.Logger
	WS                  *websocket.WebsocketServer
	UserService         services.IUserService
	NotificationService services.INotificationService
	PostService         services.IPostService
	CommentService      services.ICommentService
	ChatService         services.IChatService
}

func InitApp(repositories *models.Repositories, logger *log.Logger) *Application {

	userServices := services.InitUserService(
		logger,
		repositories.UserRepo,
		repositories.SessionRepo,
		repositories.FollowerRepo,
		repositories.NotificationRepo,
	)

	notificationServices := services.InitNotificationService(
		logger,
		repositories.UserRepo,
		repositories.FollowerRepo,
		repositories.NotificationRepo,
	)

	chatServices := services.InitChatService(
		logger,
		repositories.UserRepo,
		repositories.MessageRepo,
	)

	return &Application{
		Logger: logger,
		WS: websocket.InitWebsocket(
			logger,
			userServices,
			notificationServices,
			chatServices,
		),
		UserService:         userServices,
		NotificationService: notificationServices,
		PostService:         services.InitPostService(logger, repositories.PostRepo, repositories.AllowedPostRepo),
		CommentService:      services.InitCommentService(logger, repositories.CommentRepo),
		ChatService:         chatServices,
	}
}

func (app *Application) Home(rw http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		app.Logger.Println("Cannot access homepage")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
