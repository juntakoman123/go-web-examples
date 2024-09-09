package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "example:example@(127.0.0.1:3306)/example?parseTime=true")

	if err != nil {
		log.Fatalf("failed to open connection to mysql db %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to verifies a connection to mysql db %v", err)
	}

	log.Println("success")

}
