package main

import (
	"fmt"
)

type exemplo struct {
	map_fild   map[int]string
	slice_fild []bool
}

func main() {
	
	// Forma 01
	//	exemplo1 := exemplo{
	//		map_fild: map[int]string{1:"José", 2: "Humberto"},
	//		slice_fild: []bool{true, false, true},
	//	}
	
	// Forma 02
	map_new := map[int]string{}
	map_new[0] = "José"
	map_new[10] = "Humberto"

	exemplo1 := exemplo{
		map_fild:   map_new,
		slice_fild: []bool{true, false, true},
	}

	fmt.Println(exemplo1)
}
