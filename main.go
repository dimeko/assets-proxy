package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/master-assets-app/api"

	"github.com/master-assets-app/db_adapter"
)

const (
	port string = "8001"
)

func main() {
	fmt.Println("Starting server at port:", port)
	db := db_adapter.Connect()
	fmt.Println("Connected to db:")
	routeHandler := api.Routes()
	// credentials := handlers.AllowCredentials()
	// methods := handlers.AllowedMethods([]string{"POST"})
	// origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	httpServer := &http.Server{
		Handler:      routeHandler,
		Addr:         ":8001",
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(httpServer.ListenAndServe())

	defer db.Close()
}
