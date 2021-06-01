package main

import (
	"fmt"
)

func main() {
	
	func (x, y int){
		fmt.Print(x + y)
	}(10, 10)
	
}
