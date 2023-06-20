package main

import (
	"fmt"
	"math"
	"time"
)

func calculadoraSimples() float64 {
	// Declarando variáveis
	var escolha int
	var resultado, x, y float64
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
		fmt.Println("(5) Voltar")
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

		case 2: // Subtração
			fmt.Println("Qual o primeiro número?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo número?")
			fmt.Scan(&y)

			resultado = x - y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)

		case 3: // Multiplicação
			fmt.Println("Qual o primeiro número?")
			fmt.Scan(&x)
			fmt.Println("Qual o segundo número?")
			fmt.Scan(&y)

			resultado = x * y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)

		case 4: // Divisão
			fmt.Println("Qual o dividendo?")
			fmt.Scan(&x)
			fmt.Println("Qual o divisor?")
			fmt.Scan(&y)

			resultado = x / y
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)
			fmt.Printf("O resultado é igual a %g \n", resultado)
		case 5:
			loope = false

		default:
			fmt.Println("Você não digitou um número válido.")
		}
	}
	return resultado
}

func calculadoraGeometria() float64 {
	//declarando as variaveis
	var loope bool
	var escolha int
	var raio, resultado, base, altura float64

	loope = true

	for loope {
		// Interface de escolha de cálculo
		fmt.Println("-- Calculadora--Geometrica--")
		fmt.Println()
		fmt.Println("(1) Circulo")
		fmt.Println("(2) Quadrado/Retangulo")
		fmt.Println("(3) Triangulo")
		fmt.Println("(4) Voltar")
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
				fmt.Printf("A área do quadrado é %.2f\n", resultado)

			} else {
				fmt.Printf("A área do retangulo é igual a %.2f\n", resultado)
			}

		case 3:
			fmt.Println("Area do triangulo = base * altura / 2")
			fmt.Println("")

			time.Sleep(1 * time.Second)

			fmt.Println("Sabendo disso:")
			fmt.Println("Qual é o valor da base em metros?")
			fmt.Scan(&base)
			fmt.Println("Qual o valor da altura em metros?")
			fmt.Scan(&altura)

			resultado = base * altura / 2

			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)

			fmt.Printf("A area do triangulo é igual a %.2f\n", resultado)

		case 4:
			loope = false
			fmt.Println("|-------------------------------------------|")
			fmt.Println("")

		default:
			fmt.Println("|-------------------------------------------|")
			fmt.Println("")
			fmt.Println("Voce digitou um numero invalido")
			fmt.Println("")

			loope = true

		}
	}
	return resultado
}

func calculadoraPorcentagem() float64 {
	//declarando variaveis
	var loope bool
	var escolha int
	var valor, porcentagem, aumento, reducao, resultado float64

	loope = true

	for loope {
		fmt.Println("|-------Calculadora Percentual--------|")
		fmt.Println("(1) Aumento percentual")
		fmt.Println("(2) Redução percentual")
		fmt.Println("(3) Voltar")
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			fmt.Println("Digite o valor inicial:")
			fmt.Scan(&valor)
			fmt.Println("Em quantos por cento queremos aumentar o valor informado?:")
			fmt.Scan(&porcentagem)

			aumento = (porcentagem / 100) * valor

			resultado = aumento + valor

			fmt.Println("")
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)

			fmt.Printf("O valor apos o aumento e de %.2f\n", resultado)

		case 2:
			fmt.Println("Digite o valor inicial:")
			fmt.Scan(&valor)
			fmt.Println("Em quantos por cento queremos reduzir o valor informado?:")
			fmt.Scan(&porcentagem)

			reducao = (porcentagem / 100) * valor

			resultado = valor - reducao

			fmt.Println("")
			fmt.Println("Calculando...")
			time.Sleep(1 * time.Second)

			fmt.Printf("O valor apos a reducao e de %.2f\n", resultado)
		case 3:
			fmt.Println("")
			fmt.Println("|-------------------------------------------|")
			fmt.Println("")
			loope = false
		default:
			fmt.Println("Voce não digitou um numero valido!")
			fmt.Println("")
			fmt.Println("Tente novamente")
			fmt.Println("")
			loope = true
		}
	}
	return resultado
}

func main() {
	var ligado bool
	ligado = true

	fmt.Println("-------------------------------------------")
	fmt.Println("------- Bem-vindo(a) à calculadora --------")
	fmt.Println("-------------------------------------------")
	fmt.Println("Qual assunto vamos abordar? ")
	for ligado {

		time.Sleep(1 * time.Second)

		// Declaração de variáveis
		var Materia int // qual calculadora o usuário vai escolher
		var loope bool
		loope = true
		for loope {
			// Recebendo o valor para Materia, para poder direcionar para a função desejada

			time.Sleep(1 * time.Second)
			fmt.Println("(1) Cálculadora simples")
			fmt.Println("(2) Geometria")
			fmt.Println("(3) Porcentagem")
			fmt.Println("(4) Sair")
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
			case 3:
				calculadoraPorcentagem()
			case 4:
				fmt.Println("")
				fmt.Println("Fechando progama..")
				time.Sleep(1 * time.Second)
				loope = false
				ligado = false
			default:
				fmt.Println("Voce não digitou um numero valido!")
				fmt.Println("")
				fmt.Println("Tente novamente")
				fmt.Println("")
				loope = true
			}
		}
	}
}
