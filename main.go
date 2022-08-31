package main

import "fmt"

func init() {
	fmt.Println("antes do main")
}

func main() {
	/*
		Mesmo não chamando dentro da main, ela é chamada.
		podemos ter um init por arquivo e não por pacote.
	*/
	fmt.Println("Dentrod o main")
}
