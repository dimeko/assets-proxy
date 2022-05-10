package db_adapter

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DB struct {
	*sql.DB
}

func Connect() *DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_database := os.Getenv("DB_DATABASE")

	connection := db_user + ":" + db_pass + "@tcp(" + db_host + ")/" + db_database
	fmt.Println(connection)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	return &DB{db}
}
