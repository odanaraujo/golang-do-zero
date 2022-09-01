package main

import "fmt"

func main() {

	tarefas := make(chan int, 45)    //crio um canal de tarefas
	resultados := make(chan int, 45) //crio um canal com o resultado

	go worker(tarefas, resultados) //crio uma goroutines

	//faço um for alimentando as tarefas
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas) //fechando o canal para não receber mais dados

	// faço um for para printar o resultado
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

// <- antes do chan, recebe os dados. Depois do chan, é um canal que só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
