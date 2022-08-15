package api

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/dimeko/assets-proxy/db"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type SpaHandler struct {
	staticPath string
	indexPath  string
}

type Api struct {
	Db     *db.DB
	Routes *mux.Router
}

var apilogger, _ = zap.NewProduction()

func New(db *db.DB) *Api {
	api := &Api{
		Db:     db,
		Routes: nil,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	api.Routes = routes(mux.NewRouter(), api, logger)
	return api
}

func routes(router *mux.Router, api *Api, logger *zap.Logger) *mux.Router {
	router.HandleFunc("/assets/images/favicon.png", favicon)
	router.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	router.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))
	router.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./client/dist/fonts/"))))
	// api
	a := router.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.Use(authPreLogger(logger))
	u.HandleFunc("/login", api.LoginHandler).
		Methods(http.MethodPost, http.MethodOptions).
		Name("Create customer")
	u.HandleFunc("/user-check", api.UserCheck).Methods("GET")
	u.HandleFunc("/logout", api.LogoutHandler).Methods("POST")
	u.HandleFunc("/healthcheck", api.Healthcheck).Methods("GET")

	s := a.PathPrefix("/server").Subrouter()
	s.Use(proxyPreLogger(logger))
	s.HandleFunc("/get-file", api.GetFile).Methods("GET")
	s.HandleFunc("/map-image-directory", api.MapImgDirectory).Methods("GET")
	s.HandleFunc("/map-img-files-directory", api.MapImgFilesDirectory).Methods("GET")
	s.HandleFunc("/map-db-directory", api.MapDBDirectory).Methods("GET")
	s.HandleFunc("/upload-image", api.ImageUpload).Methods("POST")
	s.HandleFunc("/edit-db-file", api.EditDbFile).Methods("POST")

	spa := SpaHandler{staticPath: "./client/dist", indexPath: "index.html"}

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		path = filepath.Join(spa.staticPath, path)
		http.ServeFile(w, r, filepath.Join(spa.staticPath, spa.indexPath))
	})

	return router
}

func proxyPreLogger(logger *zap.Logger) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info(fmt.Sprintf("(Proxy) Url: %s", r.URL))
			h.ServeHTTP(w, r)
		})
	}
}

func authPreLogger(logger *zap.Logger) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info(fmt.Sprintf("(Auth) Url: %s.", r.URL))
			h.ServeHTTP(w, r)
		})
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/dist/assets/images/favicon.png")
}
