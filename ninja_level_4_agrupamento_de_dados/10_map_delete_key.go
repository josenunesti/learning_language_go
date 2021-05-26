package main

import (
	"fmt"
)

func main() {

	x := map[string]string{
		"José":     "Programar",
		"Josééé": "Estudar",
	}
	
	x["Bill"] = "Ler"
	
	for key, value := range x {
		fmt.Printf("Nome: %v\t Hobbie Favorito: %v\n", key, value)
	}
	
	delete(x, "Bill")
	
	fmt.Printf("\n\n")
	
	for key, value := range x {
		fmt.Printf("Nome: %v\t Hobbie Favorito: %v\n", key, value)
	}
}
