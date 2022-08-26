package main

import "fmt"

func calculosMatematicos(valor1, valor2 int) (soma int, subtracao int) {
	soma = valor1 + valor2
	subtracao = valor1 - valor2
	return
}

func main() {

	fmt.Println(calculosMatematicos(4, 2))

}
