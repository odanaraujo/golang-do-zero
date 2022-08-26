package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	//opção de for
	for i < 10 {
		i++
		time.Sleep(2 * time.Second)
		fmt.Println("Incrementando i")
	}

	//for padrão
	for j := 10; j > 0; j-- {
		time.Sleep(2 * time.Second)
		fmt.Println("Decrementando J")
	}

	//for em uma array

	nomes := [3]string{"Dan", "Sabrina", "Maju"}

	for indice, nome := range nomes {
		time.Sleep(time.Second)
		fmt.Println(indice, nome)
	}

	//for em uma string
	for indice, letra := range "PALAVRA" {
		time.Sleep(time.Second)
		fmt.Println(indice, string(letra)) // colocar string(letra) caso contrário, traz o código da tabela asc
	}

	//for em um map

	usuario := map[string]string{
		"nome":      "Dan",
		"sobreNome": "Araújo",
		"idade":     "32",
		"altura":    "184",
	}

	slice := make([]string, 0)
	for _, infos := range usuario {
		slice = append(slice, infos)
	}
	fmt.Println(slice[0])
}
