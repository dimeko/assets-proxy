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
	db_host := os.Getenv("MYSQL_HOST")
	db_user := os.Getenv("MYSQL_USER")
	db_pass := os.Getenv("MYSQL_PASSWORD")
	db_database := os.Getenv("MYSQL_DB")
	db_port := os.Getenv("MYSQL_PORT")

	connection := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_database
	fmt.Println("kfjsdfkjsndfsd", connection)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
