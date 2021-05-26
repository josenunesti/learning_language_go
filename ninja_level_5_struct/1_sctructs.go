package main

import (
	"fmt"
)

type pessoa struct {
	nome 				string
	sobrenome			string
	sabores_favoritos_sorvete		[]string 
}

func main() {
	
	pessoa1 := pessoa {
		nome: "José",
		sobrenome: "Humberto",
		sabores_favoritos_sorvete: []string{"Morango", "Chocolate"},
	}
	
	pessoa2 := pessoa{"João", "Da Silva", []string{"Chocoball", "Açai"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	
	for _, sabor := range pessoa1.sabores_favoritos_sorvete {
		fmt.Printf("\t %v \n", sabor)
	}
	
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	
	for _, sabor := range pessoa2.sabores_favoritos_sorvete{
		fmt.Printf("\t %v \n", sabor)
	}
}
