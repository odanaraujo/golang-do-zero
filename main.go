package main

import "fmt"

func soma(numeros ...int) {

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(total)
}

func textoComNumeros(palavra string, numeros ...int) {

	for _, numero := range numeros {
		fmt.Println(palavra, numero)
	}
}

func main() {

	soma(4, 2, 3, 5, 6)
	textoComNumeros("Dan", 87, 90, 8)
}
