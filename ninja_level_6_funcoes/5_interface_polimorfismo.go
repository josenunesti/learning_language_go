package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado_a float64
	lado_b float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado_a * q.lado_b
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {

	quadrado_a := quadrado{
		lado_a: 2,
		lado_b: 2,
	}
	
	circulo_a := circulo{
		raio: 5,
	}

	area_quadrado_a := info(quadrado_a)
	area_circulo_a := info(circulo_a)

	fmt.Println("Area do quadrado_a", area_quadrado_a)
	fmt.Println("Area do circulo_a", area_circulo_a)
}
