package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p *pessoa) mudaMeMetodo() {
	(*p).nome = "Metodo_" + p.nome
	p.sobrenome = "Metodo_" + p.sobrenome
	p.idade = p.idade * 2
}

func mudaMeFunc(p *pessoa) {
	(*p).nome = "Func_" + p.nome
	p.sobrenome = "func_" + p.sobrenome
	p.idade = p.idade * 3
}

func main() {

	jose := pessoa{nome: "José", sobrenome: "Reis", idade: 14}
	carol := pessoa{nome: "Carol", sobrenome: "Reis", idade: 11}
	
	fmt.Println("Antes de mudar")
	fmt.Println(jose.nome, jose.sobrenome, jose.idade)
	fmt.Println(carol.nome, carol.sobrenome, carol.idade)
	
	// Exemplo com metodo
	jose.mudaMeMetodo()
	// Exemplo com função
	mudaMeFunc(&carol)
	
	fmt.Println("\nDepois de mudar")
	fmt.Println(jose.nome, jose.sobrenome, jose.idade)
	fmt.Println(carol.nome, carol.sobrenome, carol.idade)

}
