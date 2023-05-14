package main

import "fmt"

//Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade". Escreva funções que permitam aumentar ou diminuir o salário do funcionário
//em uma determinada porcentagem e uma função que calcule o tempo de serviço do funcionário (considerando que ele começou a trabalhar aos 18 anos).

type funcionario struct {
	nome    string
	salario float64
	idade   int
}

func tempoServico(x funcionario) int {
	resultado := x.idade - 17
	return resultado
}

func aumentoSalarial(x funcionario, y float64) float64 {
	y = y / 100
	resultado := y * x.salario
	aumento := x.salario + resultado
	return aumento
}

func reducaoSalarial(z funcionario, r float64) float64 {
	r = r / 100
	resultado := z.salario * r
	reducao := z.salario - resultado
	return reducao
}

func main() {
	f := funcionario{
		nome:    "Pedro",
		salario: 1000.0,
		idade:   26,
	}
	var valor float64
	var x int
	fmt.Println("(1)Redução Salarial (2)Aumento Salarial")
	fmt.Scan(&x)

	if x == 1 {
		fmt.Println("Qual o valor da redução em %?")
		fmt.Scan(&valor)
		fmt.Println(reducaoSalarial(f, valor), "reais!")

	} else if x == 2 {
		fmt.Println("Qual o valor do aumento em %?")
		fmt.Scan(&valor)
		fmt.Println(aumentoSalarial(f, valor), "reais!")
	}
	fmt.Println(tempoServico(f), "anos de servico")
}
