package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int
	Name string `json:"name"` //to lowercase
	Age  int
}

var users = []User{
	{ID: 1, Name: "GunP", Age: 23},
	{ID: 2, Name: "GunP2", Age: 24},
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
			return
		}
		w.Write(b)
		return
	}
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500"))
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error 400"))
		}

		users = append(users, u)
		fmt.Fprintf(w, "hello %s created", "POST")

		// data := []byte(`{"name":"GP001", "Method":"POST"}`) //Convert type to byte
		// w.Write(data)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", handle)

	log.Println("Server start")
	log.Fatal(http.ListenAndServe(":2500", nil))
	log.Println("Exit")
}
