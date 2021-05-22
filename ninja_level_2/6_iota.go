package main

import (
	"fmt"
)

const (
	_ = iota
	year_1 = iota + 2020
	year_2
	year_3
	year_4 
)

func main() {
	fmt.Println(year_1)
	fmt.Println(year_2)
	fmt.Println(year_3)
	fmt.Println(year_4)
}
