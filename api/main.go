package main

import (
	"SocialNetworkRestApi/api/internal/server/handlers"
	database "SocialNetworkRestApi/api/pkg/db/sqlite"
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

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

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

	fmt.Println("successfully migrated DB..")

	app := &handlers.Application{
		Logger:  logger,
		Service: &services.Service{DB: db},
	}

	database.Seed(db)

	http.HandleFunc("/", app.Service.Authenticate(app.Home))
	http.HandleFunc("/signin", app.Login)

	logger.Printf("Starting server on port %d\n", config.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil); err != nil {
		logger.Fatal(err)
	}

}
