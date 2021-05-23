package main

import (
	"fmt"
)

func main() {

	for x := 10; x <= 100; x++ {
		fmt.Printf("O resto da divisão por %v é %v\n", x, x%4)
	}
}
