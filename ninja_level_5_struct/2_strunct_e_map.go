package main

import (
	"fmt"
)

type pessoa struct {
	nome string 
	sobrenome string
	sabores []string
}

func main() {
	
	// Forma facil 
	//var map_pessoas = map[string]pessoa{}
	
	// Com make
	map_pessoas := make(map[string]pessoa)
	
	map_pessoas["Humberto"] = pessoa{
		nome: "José",
		sobrenome: "Humberto",
		sabores: []string{"Morango", "Chocolate"},
	}
	
	map_pessoas["Silva"] = pessoa{"João", "Silva", []string{"Maracuja", "Limão"}}
	
	fmt.Println(map_pessoas)
	
	for _, pessoa := range map_pessoas{
		
		fmt.Println(pessoa.nome, pessoa.sobrenome)
		for _, sabor := range pessoa.sabores{
			fmt.Printf("\t %v \n", sabor)
		}
		
	}
	
	
}
