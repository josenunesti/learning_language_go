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
}
