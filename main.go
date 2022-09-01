package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		maneira mais utilizada para sincronizar as goroutines (fazendo-as rodar ao mesmo tempo)
		chanel - pois é um canal de comunicação. Ele pode enviar ou receber dados.

	*/

	canal := make(chan string)

	go escrever("Dan", canal) //chama o método goroutines

	/*
		msg := canal: espera receber um valor.
		Entra na função escrever, no primeiro loop recebe a mensagem que o canal envia
		programa finaliza e só imprime uma vez
	*/
	msg := <-canal
	fmt.Println(msg)
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // tá enviando um valor para dentro do canal
		time.Sleep(time.Second)
	}
}
