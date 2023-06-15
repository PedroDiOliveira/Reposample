package main

import (
	"fmt"
	"math"
	"time"
)

func calculadoraSimples() float64 {
	// Declarando variáveis
	var escolha int
	var resultado float64
	var x float64  // primeiro valor
	var y float64  // segundo valor
	var loope bool // cria um loop que caso o usuário não digite um número válido, ele retorna ao início
	loope = true

	for loope {
		// Interface de escolha de cálculo
		fmt.Println("------- Calculadora Simples-------")
		fmt.Println()
		fmt.Println("(1) Adição")
		fmt.Println("(2) Subtração")
		fmt.Println("(3) Multiplicação")
		fmt.Println("(4) Divisão")
		fmt.Scan(&escolha)
		loope = true
		fmt.Println("")
		fmt.Println("|-------------------------------------------|")
		fmt.Println("")

		switch escolha {
		case 1: // Soma
			fmt.Println("Qual o primeiro número?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo número?")
			fmt.Scan(&y)

			resultado = x + y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)
			loope = false

		case 2: // Subtração
			fmt.Println("Qual o primeiro número?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo número?")
			fmt.Scan(&y)

			resultado = x - y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)
			loope = false

		case 3: // Multiplicação
			fmt.Println("Qual o primeiro número?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo número?")
			fmt.Scan(&y)

			resultado = x * y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)
			loope = false

		case 4: // Divisão
			fmt.Println("Qual o dividendo?")
			fmt.Scan(&x)
			fmt.Println("Qual o divisor?")
			fmt.Scan(&y)

			resultado = x / y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)
			loope = false

		default:
			fmt.Println("Você não digitou um número válido.")
		}
	}
	return resultado
}

func calculadoraGeometria() float64 {
	var loope bool
	var escolha int
	loope = true
	var raio float64
	var resultado float64
	var base float64
	var altura float64

	for loope {
		// Interface de escolha de cálculo
		fmt.Println("-- Calculadora--Simples--")
		fmt.Println()
		fmt.Println("(1) Circulo")
		fmt.Println("(2) Quadrado/Retangulo")
		fmt.Println("(3) Triangulo")
		fmt.Println("(4) Cilindro")
		fmt.Scan(&escolha)
		switch escolha {
		case 1:
			fmt.Println("Podemos calcular a area de um circulo multiplicando PI com o raio ao quadrado!")
			fmt.Println("")
			time.Sleep(3 * time.Second)
			fmt.Println("OBS:Caso não tenha a medida do raio, basta pegar o diametro ou seja a medida de ponta a ponta e dividir por 2.")
			fmt.Println("")
			fmt.Println("|-------------------------------------------|")
			time.Sleep(3 * time.Second)
			fmt.Println("Sabendo disso:")
			fmt.Println("Qual é o valor do raio em metros?")
			fmt.Scan(&raio)

			resultado = math.Pi * (raio * raio)

			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)

			fmt.Printf("A área do circulo e igual a %.2f metros quadrados\n", resultado)
			loope = false
		case 2:
			fmt.Println("")
			fmt.Println("Calculamos multiplicando a base pela altura")
			fmt.Println("")

			time.Sleep(3 * time.Second)

			fmt.Println("Sabendo disso:")
			fmt.Println("Qual é o valor da base em metros?")
			fmt.Scan(&base)
			fmt.Println("Qual o valor da altura em metros?")
			fmt.Scan(&altura)

			resultado = base * altura

			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)

			if base == altura {

			}

			fmt.Printf("A área do quadrado é igual a %.2f metros quadrados\n", resultado)
			loope = false

		}
	}
	return resultado
}

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("--------- Bem-vindo(a) à calculadora ------")
	fmt.Println("-------------------------------------------")

	time.Sleep(1 * time.Second)

	// Declaração de variáveis
	var Materia int // qual calculadora o usuário vai escolher

	// Recebendo o valor para Materia, para poder direcionar para a função desejada
	fmt.Println("Qual assunto vamos abordar? ")
	time.Sleep(1 * time.Second)
	fmt.Println("(1) Cálculadora simples")
	fmt.Println("(2) Geometria")
	fmt.Println("(3) Coming soon")
	fmt.Println("")
	fmt.Scan(&Materia)
	fmt.Println("")
	fmt.Println("|-------------------------------------------|")
	fmt.Println("")

	switch Materia {
	case 1:
		calculadoraSimples()

	case 2:
		calculadoraGeometria()
	}
	fmt.Println("")
	fmt.Println("Fechando progama..")
	time.Sleep(1 * time.Second)
}
