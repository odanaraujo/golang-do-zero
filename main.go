package main

import "fmt"

type pessoa struct {
	nome             string
	idade            uint8
	altura           float32
	dataDeNascimento string
}

type estudante struct {
	pessoa
	curso   string
	periodo uint8
	campus  string
}

func main() {
	fmt.Println("Herança")

	createPeople := pessoa{"Dan", 32, 1.94, "11/12/1989"}

	createStudent := estudante{createPeople, "Sistema da Informação", 4, "Federal"}

	fmt.Println(createStudent)
	fmt.Println(createStudent.altura)

}
