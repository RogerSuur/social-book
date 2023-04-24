package main

import (
	"SocialNetworkRestApi/api/internal/server/handlers"
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

	app := &handlers.Application{
		Logger:      logger,
		UserService: services.InitUserService(repos.UserRepo, repos.SessionRepo),
	}

	database.Seed(*repos)

	http.HandleFunc("/", app.UserService.Authenticate(app.Home))
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/signup", app.Register)
	//http.HandleFunc("/logout", app.Service.Authenticate(app.Logout))
	http.HandleFunc("/profile", app.UserService.Authenticate(app.Profile))
	http.HandleFunc("/following", app.UserService.Authenticate(app.Following))
	http.HandleFunc("/followers", app.UserService.Authenticate(app.Followers))

	logger.Printf("Starting server on port %d\n", config.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil); err != nil {
		logger.Fatal(err)
	}

}
