package main

import (
	"fmt"
	_ "fmt"
	"sync"
	_ "sync"
	"sync/atomic"
	_ "sync/atomic"
)

var cont int64 // Zero value = 0
var wgr sync.WaitGroup

func main() {
	incrementos := 5000
	wgr.Add(incrementos)

	for i:=0; i < incrementos; i++{
		go incremento()
	}

	wgr.Wait()
 	fmt.Println("Contador: ", cont)
}

func incremento() {
	atomic.AddInt64(&cont, 1)
	wgr.Done()
}

//- Utilizando goroutines, crie um programa incrementador:
//- Tenha uma variável com o valor da contagem
//- Crie um monte de goroutines, onde cada uma deve:
//	- Ler o valor do contador
//	- Salvar este valor
//	- Fazer yield da thread com runtime.Gosched()
//	- Incrementar o valor salvo
//	- Copiar o novo valor para a variável inicial
//- Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//- Demonstre que há uma condição de corrida utilizando a flag -race