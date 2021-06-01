package main

import (
	"fmt"
)

func soma(x int, y int) int {
	return x + y
}

func operacao_aritmetica(operacao func(int, int) int, x []int) {
	resultado := 0
	for i:=0; i < len(x); i++ {
		resultado = operacao(resultado, x[i])
	}
	
	fmt.Println("Resultado da operação:", resultado)
}

func main() {
	
	si := []int{2,2,4,1,1,-1}
	operacao_aritmetica(soma, si)
}
