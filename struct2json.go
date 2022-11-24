package main

import (
	"encoding/json" // Convert Struct to json
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	u := User{
		ID: 1, Name: "GunP", Age: 20,
	}
	b, err := json.Marshal(u)
	fmt.Printf("byte: %v \n", b) //print value byte
	fmt.Printf("byte: %s \n", b) //print string
	fmt.Printf("byte: %T \n", b) //print type
	fmt.Println(err)
}
