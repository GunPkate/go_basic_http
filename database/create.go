package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //unused lib will disappear to keep it "_" 
)

func main() {
	url := "postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database err", err)
	}
	defer db.Close()

	log.Println("okay")
}
