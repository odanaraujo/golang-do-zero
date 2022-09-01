package main

import (
	"fmt"
	"time"
)

func main() {

	//a func main não necessita criar uma goroutines

	canal := escrever("Dan") //canal que só recebe

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) //valor que tá chegando no canal
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
