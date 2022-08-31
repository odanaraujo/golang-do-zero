package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	/**
	funções que referenciam variáveis que estão fora do seu corpo
	*/

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure() //variavel do tipo funcao

	funcaoNova() //a referência vai ser dentro da função closure(texto dentro dela será impresso)

}
