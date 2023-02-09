package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	//função anônima - vai esperar meio segundo para jogar o valor no canal
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	//função anônima - vai esperar 2 segundos para jogar o valor no canal
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(i, ": ", mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(i, ": ", mensagemCanal2)
		}
	}
	close(canal1)
	close(canal2)
}
