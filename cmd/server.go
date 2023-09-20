package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/dimeko/assets-proxy/api"
	"github.com/dimeko/assets-proxy/db"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var srvCmd = &cobra.Command{
	Use:   "server",
	Short: "Starting server",
	Run:   start,
}

func init() {
	rootCmd.AddCommand(srvCmd)
}

func start(command *cobra.Command, args []string) {
	StartServer()
}

func StartServer() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	port := os.Getenv("APP_PORT")
	if err != nil && port == "" {
		panic("Cannot find .env file")
	}

	db := db.New()
	defer db.Conn.Close()

	api := api.New(db)
	httpServer := &http.Server{
		Handler:      api.Routes,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		log.Printf("Starting server on port: %s", port)
		log.Fatal(httpServer.ListenAndServe())
	}()

	shutdown := make(chan os.Signal, 0)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("Shutting down server gracefully in 1 second.")
	time.Sleep(time.Second)
	defer cancel()

	log.Fatal(httpServer.Shutdown(ctx))
	os.Exit(0)
}
