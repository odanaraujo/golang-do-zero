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
		Dessa vez, no for infinito, ele executa as 5 impressões no método escrever
		porém, quando a goroutine finaliza, o canal não é fechado e ele fica esperando receber informações
		ocasionando o deadloock - problema que não é informado em tempo de compilação, apenas execução
	*/
	for {
		msg := <-canal
		fmt.Println(msg)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // tá enviando um valor para dentro do canal
		time.Sleep(time.Second)
	}
}
