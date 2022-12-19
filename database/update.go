package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://mkzchuoq:loPAe5lWPs4gsdvrMf2aKchys2xsGF0x@tiny.db.elephantsql.com/mkzchuoq"
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET name =$1 WHERE id = $2;")

	if err != nil {
		log.Fatal("Statement failed")
	}

	if _, err := stmt.Exec("AAA", 1); err != nil {
		log.Fatal("Update failed")
	}
}
