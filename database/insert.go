package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq"
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connection Error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name,age) VALUES($1,$2) RETURNING id", "GunP123", 23)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't insert data", err)
	}
	fmt.Println("Insert todo success id",id)
}
