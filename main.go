package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cao := cachorro{"Bob", "Pastor Alemão", 4}

	resultJSON, err := json.Marshal(cao)

	if err != nil {
		log.Fatal("Ocorreu um erro na linha 21")
	}
	fmt.Println(bytes.NewBuffer(resultJSON))

	c2 := map[string]string{
		"nome": "Amora",
		"raca": "Poodler",
	}

	cachorro2EmJSON, err := json.Marshal(c2)

	if err != nil {
		log.Fatal("Ocorreu um erro na linha 34")
	}

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
