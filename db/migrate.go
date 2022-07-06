package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type JSON_MIGRATIONS struct {
	Migrations []string `json:"migrations"`
}

func main() {
	db := Connect()
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql", driver)
	m.Up()

	defer db.Close()
}
