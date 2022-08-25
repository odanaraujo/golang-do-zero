package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int8
	array1[0] = 1
	fmt.Println(array1)

	array2 := [2]string{"Dan", "Araújo"}
	fmt.Println(array2)

	/*
		Uma forma de deixar mais dinâmico, onde os três pontinhos,
		vão fixar o tamanho desse array, de acordo com a quantidade de valores informados
	*/
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	// SLICE

	slice := []string{"Dan", "Sabrina"}
	fmt.Println(slice)

	slice = append(slice, "Maju")
	fmt.Println(slice)

	//slice não é uma array
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice))

	slice2 := array3[1:3] //1 é inclusivo e 3 é exclusivo
	fmt.Println(slice2)

	array3[1] = 4

	fmt.Println(slice2)
}
