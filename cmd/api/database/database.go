package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

	// Verificar se a conexão com o banco de dados está funcionando
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v\n", err)
	}

	fmt.Println("Connected to the database!")

	// Agora você pode usar 'db' para executar consultas SQL
	// Por exemplo:
	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		log.Fatalf("Error querying the database: %v\n", err)
	}
	defer rows.Close()

	// Processar os resultados da consulta, se necessário
	for rows.Next() {
		// Processar cada linha retornada
	}

	// Verificar se houve algum erro durante o processamento das linhas
	if err := rows.Err(); err != nil {
		log.Fatalf("Error processing rows: %v\n", err)
	}

}
