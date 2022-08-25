package main

import "fmt"

func main() {

	/*
		make é uma função que serve para criar uma fatia
		ele recebe três parâmetros: 1 - tipo do slice 2 - o tamanho do slice 3 - Capacidade máxima do slice
		caso passe da capacidade, ele dobra a capacidade inicial. No exemplo abaixo, começa no 8 e pula para 16
	*/

	slice := make([]uint16, 2, 4)

	slice = append(slice, 10, 20, 30, 50, 100, 200, 300)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	slice2 := make([]uint16, 0)

	slice2 = append(slice2, 87, 90, 8)
	fmt.Println(slice2)
}
