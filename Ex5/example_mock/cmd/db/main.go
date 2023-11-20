package main

import (
	"database/sql"
	"fmt"
	"log"

	dbPack "example_mock/internal/db"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		log.Panic(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
