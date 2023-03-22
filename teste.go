package main

import "fmt"

func main() {
	var n1 int
	var n2 float64
	fmt.Print("Qual é o primeiro numero?")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo numero?")
	fmt.Scan(&n2)
	resultado := int(float64(n1) + n2)
	fmt.Print(resultado)

}
