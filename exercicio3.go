package main

import "fmt"

// Crie uma struct chamada Triângulo com os campos "base" e "altura". Escreva uma função que receba um Triângulo
// como parâmetro e calcule a área do triângulo (área = base * altura / 2).
type triangulo struct {
	base   float64
	altura float64
}

func areaTriangulo(x triangulo) float64 {
	area := x.base * x.altura / 2
	return area
}

func main() {
	var area triangulo
	fmt.Println("Qual a base do triangulo?")
	fmt.Scan(&area.base)
	fmt.Println("Qual a altura do triangulo?")
	fmt.Scan(&area.altura)
	fmt.Println(areaTriangulo(area))
}
