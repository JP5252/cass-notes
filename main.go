package main

import (
	"log"
	"note-site/api"
	"note-site/database"
)

type Program struct {
	db *database.Database
}

func main() {
	var program Program

	if err := program.db.GetDbConnection(); err != nil {
		log.Fatalf("Could not connect to database: %s\n", err.Error())
	}

	api.StartServer()
}
