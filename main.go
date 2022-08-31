package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	/*
		podemos passar, como parâmetro da função genêrica, qualquer tipo
		int, string, bool, etc.
		o fmt.println, por exemplo, funciona como genêrico
	*/

	generica("teste")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     20,
	}

	fmt.Println(mapa)
}
