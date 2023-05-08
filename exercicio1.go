package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calculo(c Circulo) float64 {
	var area float64
	area = math.Pi * c.raio * c.raio
	return area
}

func main() {
	var raio Circulo
	fmt.Println("Digite o raio:")
	fmt.Scan(&raio.raio)

	fmt.Printf("A area do circulo é %f", calculo(raio))

}
