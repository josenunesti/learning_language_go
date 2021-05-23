package main

import (
	"fmt"
)

func main() {

	clima := "t"

	if clima == "chuvoso" {
		fmt.Println("Vamos ficar em casa")
	} else if clima == "nublado" {
		fmt.Println("Vamos levar um guarda chuva")
	} else if clima == "ensolarado" {
		fmt.Println("Vamos sair")
	} else {
		fmt.Println("Informe um clima válido")
	}
}
