package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	
	defer wg.Wait()
	
	wg.Add(2)
	go print()
	go print()
}

func print() {
	defer wg.Done()

	fmt.Println("Hello World")
}
