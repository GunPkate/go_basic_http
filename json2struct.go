package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	// name string 
	Name string `json:"name"` //ตัวเล็กตัวใหญ่ได้ ชื่อต้องตรงกับ body 
	// Name string `json:"nickname"` 
	Age  int
}

func main() {
	data := []byte(`  {
		"ID": 1,
		"name": "PGGun",
		"Age": 23
	  }`)

	var u = &User{} //User (main) Unmashal(json) คนละ package ส่ง pointer &
	err := json.Unmarshal(data, u) //Unmashal ไม่สามารถเข้าถึง private member ของ user เปลี่ยนเป็นตัวอักษรเล็ก User package main
	fmt.Printf("% #v\n", u)
	fmt.Println(err)
}
