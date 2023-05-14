package main

//Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade". Escreva funções que permitam aumentar ou diminuir o salário do funcionário
//em uma determinada porcentagem e uma função que calcule o tempo de serviço do funcionário (considerando que ele começou a trabalhar aos 18 anos).

type funcionario struct {
	nome    string
	salario float64
	idade   int
}
