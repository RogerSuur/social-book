package main

import (
	"SocialNetworkRestApi/api/internal/server/handlers"
	"SocialNetworkRestApi/api/internal/server/router"
	"SocialNetworkRestApi/api/internal/server/websocket"
	database "SocialNetworkRestApi/api/pkg/db/sqlite"
	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	port int
}

func main() {
	config := &Config{port: 8000}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	//DATABASE
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = database.RunMigrateScripts(db)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("successfully migrated DB..")

	repos := models.InitRepositories(db)
	userServices := services.InitUserService(
		repos.UserRepo,
		repos.SessionRepo,
		repos.FollowerRepo,
		repos.NotificationRepo,
	)
	notificationServices := services.InitNotificationService(
		repos.UserRepo,
		repos.FollowerRepo,
		repos.NotificationRepo,
	)
	chatServices := services.InitChatService(
		repos.UserRepo,
		repos.MessageRepo,
	)

	app := &handlers.Application{
		Logger: logger,
		WS: websocket.InitWebsocket(
			userServices,
			notificationServices,
			chatServices,
		),
		UserService:         userServices,
		NotificationService: notificationServices,
		PostService: services.InitPostService(
			repos.PostRepo,
		),
		CommentService: services.InitCommentService(
			repos.CommentRepo,
		),
		ChatService: chatServices,
	}

	args := os.Args

	if len(args) > 1 {
		switch args[1] {
		case "seed":
			database.Seed(repos)
		default:
			break
		}

	}

	r := router.New(app)

	logger.Printf("Starting server on port %d\n", config.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), r); err != nil {
		logger.Fatal(err)
	}

}
