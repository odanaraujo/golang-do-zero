package main

import (
	"fmt"
	"godozero/endereco"
)

func main() {

	result := endereco.TipoDeEndereco("RUA nossa senhora")
	fmt.Println(result)
}
