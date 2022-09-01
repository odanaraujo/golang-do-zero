package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// pegar um ou mais canal e juntar em um só

	canal := multiplexar(escrever("Dan"), escrever("Sabrina"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case message := <-canalEntrada1:
				canalDeSaida <- message
			case message := <-canalEntrada2:
				canalDeSaida <- message
			}
		}
	}()
	return canalDeSaida
}
