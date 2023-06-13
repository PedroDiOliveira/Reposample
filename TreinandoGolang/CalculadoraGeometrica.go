package main

import (
	"fmt"
	"time"
)

func calculadoraSimples() float64 {
	//declarando variaveis
	var escolha int
	var resultado float64
	var x float64  //primeiro valor
	var y float64  //segundo valor
	var loope bool //cria um loope que caso o usuario nao digite um numero valido, ele retorna ao incio
	loope = true

	for loope {
		//interface de escolha de calculo
		fmt.Println("--Calculadora Simples--")
		fmt.Println()
		fmt.Println("(1)Adição")
		fmt.Println("(2)Subtração")
		fmt.Println("(3)Multiplicação")
		fmt.Println("(4)Divisão")
		fmt.Scan(&escolha)
		loope = true

		switch escolha {
		case 1: //soma
			fmt.Println("Qual o primeiro numero?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo numero?")
			fmt.Scan(&y)

			resultado = x + y
			fmt.Println("Calculando..")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g ", resultado)
			loope = false
		case 2: //subtracao
			fmt.Println("Qual o primeiro numero?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo numero?")
			fmt.Scan(&y)

			resultado = x - y
			fmt.Println("Calculando..")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g ", resultado)
			loope = false
		case 3: //multiplicacao
			fmt.Println("Qual o primeiro numero?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo numero?")
			fmt.Scan(&y)

			resultado = x * y
			fmt.Println("Calculando..")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g ", resultado)
			loope = false
		case 4: //
			fmt.Println("Qual o dividendo?")
			fmt.Scan(&x)
			fmt.Println("Qual o divisor?")
			fmt.Scan(&y)

			resultado = x / y
			fmt.Println("Calculando..")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g ", resultado)
			loope = false
		default:
			fmt.Println("Voce nao digitou um numero valido.")
		}
	}
	return resultado
}

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("---------Bem vindo(a) a calculadora--------")
	fmt.Println("-------------------------------------------")

	time.Sleep(1 * time.Second)

	//Declaracao de variaveis
	var Materia int //qual calculadora o usuario vai escolher

	//recebendo o valora para Materia, para poder direcionar para a funcao desejada
	fmt.Println("Qual assunto vamos abordar? ")
	time.Sleep(2 * time.Second)
	fmt.Println("(1)Calculos padrões")
	fmt.Println("(2)Geometria")
	fmt.Println("(3)Coming soon")
	fmt.Println("")
	fmt.Scan(&Materia)

	time.Sleep(1 * time.Second) // 2 segundos de delay para melhorar a interface do progama, deixando as informacoes mais organizadas

	switch Materia {
	case 1:
		calculadoraSimples()
	}
}
