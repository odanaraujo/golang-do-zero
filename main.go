package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	fmt.Println("===== com ponteiro =====")
	fmt.Println()

	novoNumero := 40
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
