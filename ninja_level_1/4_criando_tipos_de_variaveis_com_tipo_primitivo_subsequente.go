package main

import (
	"fmt"
)

type comida int
var x comida

func main() {

	fmt.Printf("%v\t %T\n", x, x)
	x = 42
	fmt.Println(x)
}
