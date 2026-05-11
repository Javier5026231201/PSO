package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func ConnectDB() {
	// ==========================================
	// PILIH SALAH SATU CONNECTION STRING DI BAWAH INI
	// ==========================================

	// OPSI 1: Jika kamu login SSMS menggunakan Windows Authentication (Tanpa password)
	connString := "server=localhost\\SQLEXPRESS;database=mahasiswa_app;trusted_connection=yes;encrypt=disable"

	// OPSI 2: Jika kamu login SSMS menggunakan SQL Server Authentication (Pakai user 'sa' dan password)
	// connString := "server=localhost\\SQLEXPRESS;user id=sa;password=PasswordKamu;database=mahasiswa_app;encrypt=disable"

	database, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Gagal membuka koneksi SQL Server:", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal("SQL Server tidak merespon:", err)
	}

	fmt.Println("Database SQL Server Connected")
	DB = database
}
