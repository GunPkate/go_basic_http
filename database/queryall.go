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

	fmt.Println("OK")
	defer db.Close()

	stmt, err := db.Prepare("SELECT * from users")
	if err != nil {
		log.Fatal("query error", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("query error", err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("can't scan rows into table", err)
		}
		fmt.Println(id, name, age)
	}
}
