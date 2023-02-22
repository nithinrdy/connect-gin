package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var DatabaseInstance *sql.DB

func DbConn() (db *sql.DB) {
	envFileContents, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	password := envFileContents["DB_PASSWORD"]
	database_host_url := envFileContents["DB_HOST_URL"]
	db_port := envFileContents["DB_PORT"]
	var err2 error
	DatabaseInstance, err2 = sql.Open("postgres", fmt.Sprintf("user=postgres password=%v host=%v port=%v dbname=postgres", password, database_host_url, db_port))
	if err != nil {
		log.Fatal(err2)
	}
	return DatabaseInstance
}
