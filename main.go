package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Existe uma diferência entre concorrência e paralelismo
		Paralelismo acontece quando duas ou mais tarefas são executadas ao mesmo tempo
		Já na concorrência, não necessariamente estão executando ao mesmo tempo.
		Caso tenha um processador com apenas um núcleo, ele também executa.
		Nesse caso, ele estaria revezando a execução
	*/

	go escrever("Dan") //executa a linha 18 mas não espera ela finalizar para seguir o programa
	escrever("Sabrina")

}

func escrever(texto string) {
	for {
		fmt.Println("Olá, ", texto)

		time.Sleep(time.Second * 2)
	}
}
