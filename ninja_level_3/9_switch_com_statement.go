package main

import (
	"fmt"
)

func main() {

	esporteFavorito := "basquete"

	switch esporteFavorito {
	case "futeboll":
		fmt.Printf("Gosta de %v", esporteFavorito)
	case "basquete":
		fmt.Printf("Gosta de %v", esporteFavorito)
	default:
		fmt.Println("Não gosta de nenhum dos esportes acima")
	}
}
