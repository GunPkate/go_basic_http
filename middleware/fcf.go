package main

import "fmt"

func sum(a int, b int) int {// signature ตรงกับ Math
	return a + b
}

type Math func(int,int) int//signature

func cal(sn Math) int{
	return sn(5,4)
}

func main() {
	fn := sum
	r1 := fn(1, 2)
	fmt.Println("fn(1,2)",r1)
	
	r2 := cal(fn)
	fmt.Println("cal(fn)",r2)

	r3:= cal(sum)
	fmt.Println("cal(sum)",r3)
}