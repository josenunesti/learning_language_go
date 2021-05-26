package main

import (
	"fmt"
)


func main() {
	
	x := struct {
		map_fild   map[int]string
		slice_fild []bool
	}{
		map_fild: map[int]string{1:"José", 2: "Humberto"},
		slice_fild: []bool{true, false, true},
	}

	fmt.Println(x)
}
