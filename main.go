package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. O resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno tá aprovado")
	media := (n1 + n2) / 2
	if media > 7 {
		return true
	}

	return false
}

func main() {

	/**
	defer - Adiar a execução até o último momento possível.
	*/

	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(8, 8))
}
