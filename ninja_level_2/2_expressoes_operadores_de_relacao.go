package main

import (
	"fmt"
)

func main() {
	a := 1 == 1
	b := 1 != 1
	c := 1 <= 1
	d := 1 < 1
	e := 1 >= 1
	f := 1 > 1
	
	fmt.Printf(" %v\n %v\n %v\n %v\n %v\n %v\n", a, b, c, d, e, f)
}
