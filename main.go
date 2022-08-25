package main

import "fmt"

func main() {

	numero := 10

	//if init - a variável outroNumero tá limitada ao escopo do if

	if outroNumero := numero; outroNumero > 5 {
		fmt.Println("Número é maior mesmo")
	} else {
		fmt.Println("Número náo é maior, não.")
	}

	if numero > 3 {
		fmt.Println("deu certo")
	}
}
