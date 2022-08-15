package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Conn *sql.DB
}

func New() *DB {
	db := &DB{
		Conn: Connect(),
	}
	return db
}
func Connect() *sql.DB {
	db_host := os.Getenv("POSTGRES_HOST")
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_database := os.Getenv("POSTGRES_DB")
	db_port := "3306" //os.Getenv("POSTGRES_PORT")

	connection := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_database
	fmt.Println(connection)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
