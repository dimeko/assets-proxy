package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/master-assets-app/db_adapter"
)

type LoginBodyType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseType struct {
	Result string `json:"result"`
	Body   string `json:"body"`
}

type EditDbFileType struct {
	File_path string          `json:"file_path"`
	Data      json.RawMessage `json:"data"`
}

func (l *LoginBodyType) hashedPassword() string {
	hash := md5.Sum([]byte(l.Password))
	return hex.EncodeToString(hash[:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	db := db_adapter.Connect()

	var p LoginBodyType
	json.NewDecoder(r.Body).Decode(&p)

	var usrnm string
	var password string
	fmt.Println("Username: ", p.Username, "Password: ", p.hashedPassword())
	db.QueryRow("SELECT username, password FROM users WHERE username=? AND password=?",
		p.Username, p.hashedPassword()).Scan(&usrnm, &password)

	fmt.Println("Username: ", usrnm, "Password: ", password)
	if usrnm != "" && password != "" {
		session, _ := Session(r)
		session.Values["authenticated"] = true
		session.Values["username"] = usrnm
		// saves all sessions used during the current request
		session.Save(r, w)
		w.Write([]byte("Logged in successfully:" + usrnm))
	} else {
		w.Write([]byte("Wrong credentials"))
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	session, _ := Session(r)
	authenticated := session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if isAuthenticated != false {
			w.Write([]byte("Welcome!"))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get registers and returns a session for the given name and session store
	// session.id is the name of the cookie that will be stored in the client's browser
	session, _ := Session(r)
	// Set the authenticated value on the session to false
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte("Logout Successful"))
}

func UserCheck(w http.ResponseWriter, r *http.Request) {
	session, _ := Session(r)
	authenticated := session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if isAuthenticated != false {
			username := session.Values["username"].(string)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(ResponseType{Result: "SUCCESS", Body: string("username: " + username)})
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			json.NewEncoder(w).Encode(ResponseType{Result: "FAILURE", Body: string("message: No user found")})
			return
		}
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
		json.NewEncoder(w).Encode(ResponseType{Result: "FAILURE", Body: string("message: No user found")})
		return
	}
}

/* Api related controllers */

func GetFile(w http.ResponseWriter, r *http.Request) {
	// Setting url
	url := ProxyUri(r, "get_file")
	// Initiallizing request parameters
	fileName := r.URL.Query()["file_path"]
	req := GetRequest(url, map[string]string{"file_path": fileName[0]})
	// Setting basic auth
	UsersWebsiteAuth(r, req)
	response, err := HttpClient().Do(req)
	if err != nil {
		panic(err.Error())
	}
	HttpResponder(w, response)
}

func MapImgDirectory(w http.ResponseWriter, r *http.Request) {
	url := ProxyUri(r, "img_directory")
	req := GetRequest(url, nil)
	UsersWebsiteAuth(r, req)
	response, err := HttpClient().Do(req)
	if err != nil {
		panic(err.Error())
	}
	HttpResponder(w, response)
}

func MapImgFilesDirectory(w http.ResponseWriter, r *http.Request) {
	url := ProxyUri(r, "img_files_directory")
	req := GetRequest(url, nil)
	UsersWebsiteAuth(r, req)
	response, err := HttpClient().Do(req)
	if err != nil {
		panic(err.Error())
	}
	HttpResponder(w, response)
}

func MapDBDirectory(w http.ResponseWriter, r *http.Request) {
	url := ProxyUri(r, "db_directory")
	req := GetRequest(url, nil)
	UsersWebsiteAuth(r, req)
	response, err := HttpClient().Do(req)
	if err != nil {
		panic(err.Error())
	}
	HttpResponder(w, response)
}

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 10MB in size", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image_file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	imageExtension := filepath.Ext(handler.Filename)
	hashedImgName := strings.Join([]string{(HashString(handler.Filename)), imageExtension}, "")

	tmpfile, err := os.Create("./" + hashedImgName)
	_, err = io.Copy(tmpfile, file)
	defer os.Remove("./" + hashedImgName)
	defer tmpfile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("category")
	if err != nil {
	}
	cat := r.PostFormValue("category")
	_, err = io.Copy(fw, strings.NewReader(cat))
	if err != nil {
		return
	}

	fw, err = writer.CreateFormFile("file_to_upload", hashedImgName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newfile, ferr := os.Open(hashedImgName)

	if ferr != nil {
		panic(ferr)
	}
	_, err = io.Copy(fw, newfile)
	if err != nil {
		return
	}

	writer.Close()

	req := PostRequest(ProxyUri(r, "upload_img"), bytes.NewReader(body.Bytes()))
	UsersWebsiteAuth(r, req)

	defer file.Close()

	req.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := HttpClient().Do(req)
	HttpResponder(w, response)
}

func EditDbFile(w http.ResponseWriter, r *http.Request) {
	var data EditDbFileType
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := PostRequest(ProxyUri(r, "edit_file"), &buf)
	UsersWebsiteAuth(r, req)
	response, err := HttpClient().Do(req)
	HttpResponder(w, response)
}
