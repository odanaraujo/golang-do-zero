package main

import "fmt"

func main() {

	/*
		maneira mais utilizada para sincronizar as goroutines (fazendo-as rodar ao mesmo tempo)
		chanel - pois é um canal de comunicação. Ele pode enviar ou receber dados.

	*/

	//canal com buffer só bloqueia quando atinge a capacidade total dele
	canal := make(chan string, 2)

	canal <- "Dan"
	canal <- "Programando em go"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
