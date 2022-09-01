package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type carro struct {
	Marca         string `json:"marca"`
	AnoFabricacao uint16 `json:"anoFabricacao"`
	Modelo        string `json:"modelo"`
}

func main() {

	celta := carro{"Chebrollet", 2010, "2 portas"}

	infoCarroJSON, erro := json.Marshal(celta)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(infoCarroJSON)
	fmt.Println(bytes.NewBuffer(infoCarroJSON))

	gol := map[string]string{
		"marca":         "volkswagen",
		"anoFabricacao": "2020",
		"modelo":        "4 portas",
	}

	infoCarroJSON2, erro := json.Marshal(gol)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(infoCarroJSON2)
	fmt.Println(bytes.NewBuffer(infoCarroJSON2))

}
