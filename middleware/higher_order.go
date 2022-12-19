package main

import "fmt"

func home(s string) error {
	fmt.Println(s)
	return nil
}

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(c string) error {
		r := c + " green" //Decorator /middleware
		return next(r)
	}
}

func main() {
	home("GP")
	wrapped := Use(home)
	w := wrapped("Hp") //override
	println("result: ", w)
} // Use => fmt.print => return next(c)//home("world")
