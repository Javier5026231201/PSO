package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// Sesuaikan dengan kredensial lokal Anda, atau environment variable GCP nanti
	connStr := "host=localhost port=5432 user=postgres password=root dbname=mahasiswa_app sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected Successfully")
}
