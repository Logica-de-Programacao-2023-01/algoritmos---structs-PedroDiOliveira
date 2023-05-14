package main

import "fmt"

type musica struct {
	titulo  string
	artista string
	duracao float64
}

type playlist struct {
	nome    string
	musicas []musica
}

func spotify(p playlist) {
	nomeP := p.nome
	fmt.Println(nomeP)
	for _, musica := range p.musicas {
		fmt.Println("Título:", musica.titulo)
		fmt.Println("Artista:", musica.artista)
		fmt.Println("Duração:", musica.duracao)
	}

	duracaoTotal := 0.0
	for _, m := range p.musicas {
		duracaoTotal += m.duracao
	}

	fmt.Println("O tempo total de duração é igual a:", duracaoTotal)
}

func main() {
	// Exemplo de criação de uma playlist
	p := playlist{
		nome: "Minha Playlist",
		musicas: []musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 3.5},
			{titulo: "Música 2", artista: "Artista 2", duracao: 4.2},
			{titulo: "Música 3", artista: "Artista 3", duracao: 2.8},
		},
	}

	spotify(p)
}
