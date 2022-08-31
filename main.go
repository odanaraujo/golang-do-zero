package main

import (
	"fmt"
	"time"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
	time.Sleep(time.Second * 3)
	fmt.Println("Salvo com sucesso!")
}

func (u usuario) update() {
	fmt.Printf("Atualizando os dados do usuário %s no banco de dados \n", u.nome)
	time.Sleep(time.Second * 4)
	fmt.Println("Dados atualizados com sucesso")
}

func (u usuario) maiorDeIdade() bool {
	return u.idade > 18
}

func (u *usuario) felizAniversario() {
	u.idade++
}

func main() {

	/*
		Tecnicamente, método não é igual a uma função
		ela pode tá associada a uma struct, interface, etc.
	*/

	usuario := usuario{"Dan", 32}
	fmt.Println(usuario)
	usuario.salvar()
	usuario.update()
	fmt.Println(usuario.maiorDeIdade())
	usuario.felizAniversario()
	fmt.Println(usuario.idade)
}
