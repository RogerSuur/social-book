package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(rw, "Homepage hit")
	if err != nil {
		log.Println("Cannot access homepage")
	}
}
