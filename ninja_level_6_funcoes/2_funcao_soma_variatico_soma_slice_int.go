package main

import (
	"fmt"
)

func main() {
	
	soma_a:= soma_variadico([]int{1,2,3,4,5}...)
	fmt.Println(soma_a)
	
	soma_b := soma_slice_int([]int{1,2,3,4,5})
	fmt.Println(soma_b)
}

func soma_variadico(x ...int) int {
	total := 0
	for _, v := range x {
		
		total += v
	}
	return total
}

func soma_slice_int(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
