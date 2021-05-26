package main

import (
	"fmt"
)

func main() {
	
	x := make([]string, 10)
	x = []string{"São Paulo", "Minas Gerais", "Bahia"}
	
	fmt.Printf("Len: %v\t Cap: %v\n", len(x), cap(x))
	
	for i :=0; i < len(x); i++{
		fmt.Println(x[i])
	}
}
