package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func database() {
	// Configuração da conexão com o banco de dados MySQL
	cfg := mysql.Config{
		User:   "username",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "database_name",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
	}

	fmt.Println("Connected to the database!")

	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		log.Fatalf("Error querying the database: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error processing rows: %v\n", err)
	}

}
