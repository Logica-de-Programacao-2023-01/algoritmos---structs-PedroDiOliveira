package main

import "fmt"

type Endereco struct {
	rua    int
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func local(x Pessoa) {
	fmt.Println(x.endereco)
}

func main() {
	endereco := Endereco{
		rua:    11,
		numero: 2,
		estado: ("DF"),
	}
	pessoa := Pessoa{
		nome:     ("Pedro"),
		idade:    18,
		endereco: endereco,
	}
	local(pessoa)
}
