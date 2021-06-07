package main

import (
	"fmt"
	_ "fmt"
	"runtime"
	_ "runtime"
)

func main(){

	fmt.Printf("Arquitetura do processador: %v\nSO: %v", runtime.GOARCH, runtime.GOOS)

}
