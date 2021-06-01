package main

import (
	"fmt"
)

func main() {

	x := func () func() {
		return func () {
			fmt.Println("Hello, playground")	
		}
	}()
	
	x()
}
