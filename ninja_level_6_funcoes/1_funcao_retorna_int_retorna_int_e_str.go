package main

import (
	"fmt"
)

func main() {
	
	a := retorna_int()
	b, c := retorna_int_str()
	
	fmt.Println(a, b, c)
}

func retorna_int() int {
	return 2
}

func retorna_int_str() (int, string) {
	return 5, "cinco"
}