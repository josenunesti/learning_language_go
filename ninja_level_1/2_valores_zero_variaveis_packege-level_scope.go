package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {

	fmt.Println(x, y, z)
	fmt.Printf("x: %T\n", x)
	fmt.Printf("y: %T\n", y)
	fmt.Printf("z: %T\n", z)
}
