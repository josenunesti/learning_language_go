package main

import (
	"fmt"
)

type comida int
var x comida
var y int

func main() {

	fmt.Printf("Valor Zero do tipo comiga: %v\t Tipo: %T\n", x, x)
	x = 42
	fmt.Printf("Valor da variavel do tipo comida: %v\n", x)
	
	y = int(x)
	
	fmt.Printf("Valor da variavel do tipo comida após atribuição: %v\t Tipo: %T", y, y)
	
}
