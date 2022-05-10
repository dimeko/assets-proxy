package api

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func Routes() *mux.Router {
	mux := mux.NewRouter()

	// client files
	// // mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	// mux.PathPrefix("/admin").Handler(http.StripPrefix("/admin", http.FileServer(http.Dir("./client/dist/")))).Methods("GET")
	// mux.PathPrefix("/admin/").Handler(http.FileServer(http.Dir("./client/dist/")))

	mux.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./client/dist/js/"))))
	mux.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./client/dist/css/"))))
	mux.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./client/dist/fonts/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/login", LoginHandler).Methods("POST")
	u.HandleFunc("/user-check", UserCheck).Methods("GET")
	u.HandleFunc("/logout", LogoutHandler).Methods("GET")
	u.HandleFunc("/healthcheck", Healthcheck).Methods("GET")

	s := a.PathPrefix("/server").Subrouter()
	s.HandleFunc("/get-file", GetFile).Methods("GET")
	s.HandleFunc("/map-image-directory", MapImgDirectory).Methods("GET")
	s.HandleFunc("/map-img-files-directory", MapImgFilesDirectory).Methods("GET")
	s.HandleFunc("/map-db-directory", MapDBDirectory).Methods("GET")
	s.HandleFunc("/upload-image", ImageUpload).Methods("POST")
	s.HandleFunc("/edit-db-file", EditDbFile).Methods("POST")
	// s.HandleFunc("/edit-file", EditFile).Methods("POST")

	// s.HandleFunc("/upload-image", UploadImage).Methods("POST")
	// s.HandleFunc("/delete-image", DeleteImage).Methods("POST")

	spa := spaHandler{staticPath: "./client/dist", indexPath: "index.html"}

	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, err := filepath.Abs(r.URL.Path)
		if err != nil {
			// if we failed to get the absolute path respond with a 400 bad request
			// and stop
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		path = filepath.Join(spa.staticPath, path)
		http.ServeFile(w, r, filepath.Join(spa.staticPath, spa.indexPath))
	})

	return mux
}
