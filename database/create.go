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

	createTb := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name TEXT,age INT)`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
	log.Println("okay")
}
