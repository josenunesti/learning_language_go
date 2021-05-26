package main

import (
	"fmt"
)

const (
	x = 0
	y int = 0
)

func main() {
	fmt.Printf("Não tipada: %v\t %T", x, x)
	fmt.Printf("\nTipada: %v\t %T", y, y)
}
