package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/master-assets-app/db_adapter"
)

type JSON_MIGRATIONS struct {
	Migrations []string `json:"migrations"`
}

func main() {
	db := db_adapter.Connect()
	file, err := ioutil.ReadFile("migrations/init_db.json")

	if err != nil {
		fmt.Println(err)
	}

	data := JSON_MIGRATIONS{}
	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Migrations); i++ {
		fmt.Println("Executing:", data.Migrations[i])
		_, err := db.Query(data.Migrations[i])

		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
}
