package main

import (
	"fmt"
)

func main() {

	ano := 1994
	for {
		fmt.Println(ano)
		ano++

		if ano > 2021 {
			break
		}
	}
}
