package main

import (
	"fmt"
)

type veiculo struct{
	portas	int
	cor	string
}

type caminhonete struct{
	veiculo
	tracao_quatro	bool
}

type sedan struct{
	veiculo
	modelo_luxo bool
}

func main() {
	
	caminhonete1 := caminhonete{veiculo{4, "Branco"}, true}
	
	sedan1 := sedan{
		veiculo: veiculo{
			portas: 2, 
			cor: "Cinza",
		}, 
		modelo_luxo: false,
	}
	
	fmt.Println(caminhonete1)
	fmt.Printf("\t %v \n", caminhonete1.cor)
	fmt.Println(sedan1)
	fmt.Printf("\t %v \n", sedan1.modelo_luxo)
	
	
}
