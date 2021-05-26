package main

import (
	"fmt"
)

func main() {
	
	clima := "ensolarado"
	
	switch {
		case clima == "chuvoso":
			fmt.Println("Vamos ficar em casa")
		case clima == "ensolarado":
			fmt.Println("Vamos sair")
		default:
			fmt.Println("Informe um clima válido")
	}
}
