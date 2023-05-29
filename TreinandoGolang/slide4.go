package main

import (
	"fmt"
	"sort"
)

type Camisa struct {
	Preco   float64
	Tamanho string
}

func main() {
	listaCamisas := []Camisa{
		{Preco: 19.99, Tamanho: "M"},
		{Preco: 24.99, Tamanho: "XXL"},
		{Preco: 15.99, Tamanho: "S"},
		{Preco: 29.99, Tamanho: "XXXXXXXS"},
	}

	media := calcularMediaPrecosCamisas(listaCamisas)
	fmt.Printf("Média entre o preço da maior camisa e o preço da menor camisa: %.2f\n", media)
}

func calcularMediaPrecosCamisas(camisas []Camisa) float64 {
	var menorPreco float64
	var maiorPreco float64

	sort.Slice(camisas, func(i, j int) bool {
		return compararTamanhos(camisas[i].Tamanho, camisas[j].Tamanho)
	})

	for _, camisa := range camisas {
		tamanho := camisa.Tamanho
		preco := camisa.Preco

		if tamanhoValido(tamanho) {
			if menorPreco == 0 {
				menorPreco = preco
			}
			maiorPreco = preco
		}
	}

	return (maiorPreco + menorPreco) / 2
}

func tamanhoValido(tamanho string) bool {
	if tamanho == "M" {
		return true
	}

	if len(tamanho) < 2 {
		return false
	}

	if tamanho[0] != 'X' {
		return false
	}

	if tamanho[len(tamanho)-1] != 'S' && tamanho[len(tamanho)-1] != 'L' {
		return false
	}

	return true
}

func compararTamanhos(tamanho1, tamanho2 string) bool {
	if tamanho1 == "M" {
		return true
	}

	if tamanho2 == "M" {
		return false
	}

	if len(tamanho1) < len(tamanho2) {
		return true
	}

	if len(tamanho1) > len(tamanho2) {
		return false
	}

	if tamanho1[len(tamanho1)-1] == 'S' && tamanho2[len(tamanho2)-1] == 'L' {
		return true
	}

	return false
}
