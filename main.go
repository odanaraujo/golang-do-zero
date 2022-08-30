package main

import "fmt"

func main() {

	/**
	função que não tem nome
	*/

	func(texto string) {
		fmt.Println(texto)
	}("Dan")

	retorno := func(texto string, numero int) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, numero)
	}("Passando parâmetro", 10)

	fmt.Println(retorno)
}
