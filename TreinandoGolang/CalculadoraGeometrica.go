package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("---Bem vindo(a) a calculadora geometrica---")
	fmt.Println("-------------------------------------------")

	time.Sleep(2 * time.Second)

	//Declaracao de variaveis
	var Materia int

	fmt.Println("Qual assunto vamos abordar? ")
	time.Sleep(2 * time.Second)
	fmt.Println("(1)Calculos padrões")
	fmt.Println("(2)Geometria")
	fmt.Println("(3)Coming soon")
	fmt.Println("")
	fmt.Scan(&Materia)

	time.Sleep(2 * time.Second)

	if Materia == 1 {
		fmt.Println("--Calculadora Simples--")
		fmt.Println()
		fmt.Println("(1)Adição")
		fmt.Println("(2)Multiplicação")
		fmt.Println("(3)Divisão")

	}
}
