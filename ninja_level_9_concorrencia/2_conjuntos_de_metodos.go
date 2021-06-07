package main

import (
	"fmt"
	_ "fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p *pessoa) falar(){
	fmt.Println("Meu nome é ", p.nome, p.sobrenome, "e minha idade é ", p.idade)
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos){
	h.falar()
}

func main(){

	p1 := pessoa{nome: "José", sobrenome: "Humberto", idade: 26}

	p1.falar() // É uma forma de chamar o metodo
	(&p1).falar() // É a forma correta, porém menos prática

	dizerAlgumaCoisa(&p1) // Chamando a função da forma correta

	//dizerAlgumaCoisa(p1) // Desta forma não funciona, pois não esta sendo passado a referencia da pessoa

}