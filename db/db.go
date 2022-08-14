package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	db_host := os.Getenv("DOCK_MYSQL_HOST")
	db_user := os.Getenv("DOCK_MYSQL_USER")
	db_pass := os.Getenv("DOCK_MYSQL_PASS")
	db_database := os.Getenv("DOCK_MYSQL_DB")
	db_port := os.Getenv("DOCK_MYSQL_PORT")

	connection := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_database

	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
