package main

func main(){
	r:= func(a,b int) bool{
		return a >b
	}(2,3)

	println("a>b",r)
}

