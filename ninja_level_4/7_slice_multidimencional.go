package main

import (
	"fmt"
)

func main() {
	
	x := [][]string{
		[]string{
			"José",
			"Humberto",
			"Programar",
		},
		[]string{
			"Josééééé",
			"Doisberto",
			"Programar mais",
		},
	}
	
	for _, val1 := range x{
		fmt.Println(val1[0])
		for _, val2 := range val1 {
			fmt.Println("\t", val2)
		}
	}
}
