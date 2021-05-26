package main

import (
	"fmt"
)

func main() {
	
	array := [5]int{10, 20, 30, 40, 50}
	
	for _, val := range array {
		fmt.Println(val)
	}
	
	fmt.Printf("%T", array)
}
