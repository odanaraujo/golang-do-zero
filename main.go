package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	teste := fibonacci(posicao-2) + fibonacci(posicao-1)
	fmt.Println(teste)
	return teste
}

func main() {

	/**
	Função recursiva são funções que chamam ela mesma.
	*/

	fmt.Println(fibonacci(3))
}
