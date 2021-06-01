package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) apresentar_se() {
	fmt.Printf("Olá meu nome é %v %v e tenho %v anos.\n", p.nome, p.sobrenome, p.idade)
}

func main() {
	
	jose := pessoa{
		nome: "José",
		sobrenome: "Humberto",
		idade: 26,
	}
	
	jose.apresentar_se()
}
