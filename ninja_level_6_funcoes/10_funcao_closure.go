package main

import (
	"fmt"
)

func main() {

	x := soma()
	n1 := x()
	n2 := x()
	n3 := x()
	fmt.Println(n1, n2, n3)
	
	y := soma()
	n4 := y()
	n5 := y()
	n6 := y()
	fmt.Println(n4, n5, n6)
}

func soma() func() int {
	x := 1
	return func() int{
		x++
		return x
	}
}
