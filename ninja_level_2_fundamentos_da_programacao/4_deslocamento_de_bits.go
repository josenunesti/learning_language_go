package main

import (
	"fmt"
)

func main() {
	n1 := 10
	fmt.Printf("%d\t%b\t%#x\t", n1, n1, n1)
	n1_d := n1 << 1
	fmt.Printf("\n%d\t%b\t%#x\t", n1_d, n1_d, n1_d)
}
