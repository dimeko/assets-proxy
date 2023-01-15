package api

import (
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type JsonResponse struct {
	Result string          `json:"result"`
	Body   json.RawMessage `json:"body"`
}

func UsersWebsite(r *http.Request, db *sql.DB) string {
	username := SessionUser(r)
	var website string
	db.QueryRow("SELECT website FROM users WHERE username=?",
		username).Scan(&website)

	return website
}

func ProxyUri(r *http.Request, route string, db *sql.DB) string {
	var proxyPath string
	switch route {
	case "get_file":
		proxyPath = "/server/get_file.php"
	case "create_file":
		proxyPath = "/server/create_file.php"
	case "img_files_directory":
		proxyPath = "/server/get_image_files_map.php"
	case "img_directory":
		proxyPath = "/server/get_image_dir_map.php"
	case "db_directory":
		proxyPath = "/server/get_db_dir_map.php"
	case "upload_img":
		proxyPath = "/server/upload_image.php"
	case "edit_file":
		proxyPath = "/server/edit_file.php"
	case "image":
		proxyPath = "/images"
	default:
		proxyPath = "/server/get_file.php"
	}

	return strings.Join([]string{"https://", UsersWebsite(r, db), proxyPath}, "")
}

func UsersWebsiteAuth(r *http.Request, proxyReq *http.Request, db *sql.DB) {
	username := SessionUser(r)
	var auth_username string
	var auth_password string
	db.QueryRow("SELECT website_auth_username, website_auth_password FROM users WHERE username=?",
		username).Scan(&auth_username, &auth_password)
	proxyReq.SetBasicAuth(auth_username, auth_password)
}

func HashString(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func HttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
}

func PostRequest(url string, payload io.Reader) *http.Request {
	apilogger.Info(fmt.Sprintf("POST request to: %s", url))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		apilogger.Error(fmt.Sprintf("Error making POST request. Error: %s", err))
	}
	return req
}

func GetRequest(url string, query map[string]string) *http.Request {
	apilogger.Info(fmt.Sprintf("GET request to: %s", url))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		apilogger.Error(fmt.Sprintf("Error making GET request. Error: %s", err))
	}
	q := req.URL.Query()
	for key, el := range query {
		q.Add(key, el)
	}
	req.URL.RawQuery = q.Encode()

	return req
}

func HttpResponder(w http.ResponseWriter, resp *http.Response, responseType string) {
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	if err != nil {
		apilogger.Error(fmt.Sprintf("Could not decode response: %s", err))
	}

	apilogger.Info(fmt.Sprintf("Method: HttpResponder. Response code: %d", resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if responseType == "image" {
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(body)
		} else {
			var jsonObResponse JsonResponse
			if err := json.Unmarshal([]byte(string(body)), &jsonObResponse); err != nil {
				apilogger.Fatal(fmt.Sprintf("Could not decode %s response.", responseType))
			}
			json.NewEncoder(w).Encode(jsonObResponse)
		}
	} else {
		http.Error(w, "Not found", resp.StatusCode)
		json.NewEncoder(w).Encode(ResponseType{Result: "FAILURE", Body: string("message: Error occured")})
	}
}
