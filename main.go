package main

import "fmt"

func main() {

	//chave e valor tem que ser do mesmo tipo
	usuario := map[string]string{
		"nome":      "Dan",
		"sobreNome": "Araújo",
		"idade":     "32",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"], usuario["sobreNome"], "Tem", usuario["idade"], "anos")

	//map aninhado

	usuario2 := map[string]map[string]string{
		"dadosPessoais": {
			"nome":      "Sabrina",
			"sobreNome": "Lopes",
			"idade":     "33",
		},
		"endereco": {
			"rua": "Rua do golang",
			"cep": "09494030",
		},
	}

	fmt.Println(usuario2["endereco"]["rua"])

	//deletando
	delete(usuario2, "endereco")
	fmt.Println(usuario2)

	//adicionando um novo map
	usuario2["cursos"] = map[string]string{
		"Sistema da informação": "completo",
		"Enfermagem":            "Cursando",
	}

	fmt.Println(usuario2)
}
