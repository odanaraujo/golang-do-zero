package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func calcularNota(notaPrimeiroSemestre, notaSegundoSemestre float32) bool {

	defer recuperarExecucao()

	media := (notaPrimeiroSemestre + notaSegundoSemestre) / 2

	if media >= 7 {
		return true
	}

	if media < 6 {
		return false
	}
	panic("A média não pode ser igual a 6")

}

func main() {

	fmt.Println(calcularNota(6, 6))
	fmt.Println("Pós execução da função")
}
