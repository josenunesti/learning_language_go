package main

import (
	"fmt"
)

func main() {
	
	multiplicacao := func (x ...int) (int, int) {
		total := 1
		for _, v := range x {
			total *= v
		}
		
		return total, len(x)
	}
	
	resultado, qnt := multiplicacao([]int{1,2,3,4,5}...)
	
	fmt.Print(resultado, qnt)
}
