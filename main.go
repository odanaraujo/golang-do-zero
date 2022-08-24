package main

import "fmt"

type pessoa struct {
	nome             string
	idade            uint8
	altura           float32
	dataDeNascimento string
}

func main() {
	fmt.Println("Struct")

	// primeira forma de declarar uma struct

	createPeople := pessoa{"Dan", 32, 1.94, "11/12/1989"}

	fmt.Println(createPeople)

	// segunda forma de declarar uma struct

	var createPeopleTwo pessoa
	createPeopleTwo.nome = "Sabrina"
	createPeopleTwo.idade = 33
	createPeopleTwo.altura = 1.74
	createPeopleTwo.dataDeNascimento = "13/10/1988"

	fmt.Println(createPeopleTwo)

	// terceira forma de declarar uma struct

	createPeopleThree := pessoa{nome: "Maria Julia", idade: 1, altura: 1.10, dataDeNascimento: "27/10/2021"}
	fmt.Println(createPeopleThree)
}
