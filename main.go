package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	default:
		return "Escolha um dia válido"
	}
}

func diaDaSemana2(dia int) string {
	switch {
	case dia == 1:
		return "Domingo"
	case dia == 2:
		return "Segunda"
	case dia == 3:
		return "Terça"
	case dia == 4:
		return "Quarta"
	case dia == 5:
		return "Quinta"
	case dia == 6:
		return "Sexta"
	case dia == 7:
		return "Sábado"
	default:
		return "Informe um dia da semana válido"
	}
}

func main() {

	//padrão
	fmt.Println(diaDaSemana(10))

	// outra forma de criar switch
	fmt.Println(diaDaSemana2(1))

}
