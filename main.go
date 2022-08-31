package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	/**
	Função recursiva são funções que chamam ela mesma.
	*/

	fmt.Println(fibonacci(3))
}
