package main

import "fmt"

var n int

func init() {
	fmt.Println("antes do main")
	n = 10
}

func main() {
	/*
		Mesmo não chamando dentro da main, ela é chamada.
		podemos ter um init por arquivo e não por pacote.
		Podemos iniciar uma variável dentro da init
	*/
	fmt.Println(n)
	fmt.Println("Dentrod o main")
}
