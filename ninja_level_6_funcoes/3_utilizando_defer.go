package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("primeiro statment do medodo main") // Este statment será demostrado por ultimo, pois possui o defer
	fmt.Println("segundo statment do medodo main")
	fmt.Println("terceiro statment do medodo main")
	fmt.Println("quarto statment do medodo main")
}
