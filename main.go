package main

import "fmt"

func main() {

	fmt.Println("Ponteiros")
	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel1++

	fmt.Println("===== Sem ponteiro ====")
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	/*
		A variável 2 continua com o mesmo valor 10, pois, ele é uma cópia da variável 1,
		mas não tá no mesmo endereço de memória onde essa variável tá salta
	*/

	var variavelComPonteiro1 uint8
	var variavelComPonteiro2 *uint8

	fmt.Println(variavelComPonteiro1) //0
	fmt.Println(variavelComPonteiro2) // nil

	variavelComPonteiro1 = 10
	variavelComPonteiro2 = &variavelComPonteiro1

	fmt.Println(variavelComPonteiro1) //10
	fmt.Println(variavelComPonteiro2) //0x1400001c092 -> valor da memória da variável

	fmt.Println(variavelComPonteiro1)
	fmt.Println(*variavelComPonteiro2) // desreferenciação -> valor vai sair 10

	variavelComPonteiro1 = 100

	fmt.Println(variavelComPonteiro1)
	fmt.Println(*variavelComPonteiro2) // ambos saem 100, pois, estão no mesmo endereço de memória.
}
