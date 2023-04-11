package main

import (
	"SocialNetworkRestApi/api/internal/server/handlers"
	database "SocialNetworkRestApi/api/pkg/db/sqlite"
	service "SocialNetworkRestApi/api/pkg/services"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	port int
}

func main() {
	config := &Config{port: 8090}

	//DATABASE
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = database.RunMigrateScripts(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully migrated DB..")

	database.Seed(db)

	http.HandleFunc("/", service.Authenticate(handlers.Home))
	http.HandleFunc("/login", handlers.Login)

	fmt.Printf("Starting server on port %d\n", config.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil); err != nil {
		log.Fatal(err)
	}

}
