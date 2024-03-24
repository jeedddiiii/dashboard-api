package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDb() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=jedi2002 dbname=facedetection sslmode=disable")

	// db, err = sql.Open("postgres", "host=host.docker.internal port=5432 user=postgres password=jedi2002 dbname=facedetection sslmode=disable")
	if err != nil {
		fmt.Println("Error opening database:", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
