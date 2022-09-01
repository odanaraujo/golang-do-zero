package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type carro struct {
	Marca         string `json:"marca"`
	AnoFabricacao string `json:"anoFabricacao"`
	Modelo        string `json:"modelo"`
}

func main() {

	golJSON := `{"marca":"volkswagen", "anoFabricacao":"2020", "modelo":"4 portas"}`

	var gol carro

	if erro := json.Unmarshal([]byte(golJSON), &gol); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(gol)

	celtaJSON := `{"marca":"chevrolet", "anoFabricacao":"2010", "modelo":"4 portas"}`
	celta := make(map[string]string)

	if erro := json.Unmarshal([]byte(celtaJSON), &celta); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(celta)

}
