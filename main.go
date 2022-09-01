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
		no fim da execução do for dentro do método, eu fecho o canal
		no for abaixo (infinito) eu passo um segundo parâmetro que é o aberto
		verifico se ele ainda se encontra em aberto, caso não eu paro a execução do for
	*/
	// for {
	// 	msg, aberto := <-canal
	// 	fmt.Println(msg)
	// 	if !aberto {
	// 		break
	// 	}
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // tá enviando um valor para dentro do canal
		time.Sleep(time.Second)
	}
	close(canal)
}
