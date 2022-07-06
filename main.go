package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/master-assets-app/api"

	"github.com/master-assets-app/db"
)

const (
	port string = ":8001"
)

func main() {
	fmt.Println("Starting server at port:", port)
	db := db.Connect()
	routeHandler := api.Routes()
	httpServer := &http.Server{
		Handler:      routeHandler,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(httpServer.ListenAndServe())

	defer db.Close()
}
