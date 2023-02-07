package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //adicionando duas goroutines que ele precisa esperar terminar

	//primeira goroutine
	go func() {
		escrever("Dan")
		waitGroup.Done() // -1 quando a função finaliza, ele termina e tira 1 goroutine do count 2.
	}()

	go func() {
		escrever("Sabrina")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() //fala para o programa para ele aguardar o waitgroup chegar em zero
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
