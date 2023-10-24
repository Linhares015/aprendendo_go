package main

import (
	"fmt"
	"os"
)

func adicionar(x, y float64) float64 {
	return x + y
}

func subtrair(x, y float64) float64 {
	return x - y
}

func multiplicar(x, y float64) float64 {
	return x * y
}

func dividir(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Erro: Divisão por zero!")
		os.Exit(1)
	}
	return x / y
}

func main() {
	var num1, num2 float64
	var operacao string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Digite a operação: ")
	fmt.Scan(&operacao)

	switch operacao {
	case "+":
		fmt.Printf("%f + %f = %f", num1, num2, adicionar(num1, num2))
	case "-":
		fmt.Printf("%f - %f = %f", num1, num2, subtrair(num1, num2))
	case "*":
		fmt.Printf("%f * %f = %f", num1, num2, multiplicar(num1, num2))
	case "/":
		fmt.Printf("%f / %f = %f", num1, num2, dividir(num1, num2))
	default:
		fmt.Println("Operação inválida!")
	}

}
