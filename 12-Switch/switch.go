package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default: //default é usado quando não há caso para o valor (else)
		return "Dia da semana inválido"
	}
}

func diaDaSemana2(numero int) string { //2° maneira de usar o switch
	var diaDaSemana string
	
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		//fallthrough //se o caso acima for verdadeiro, ele vai para o caso abaixo
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Dia da semana inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5)
	fmt.Println(dia)

	fmt.Println("--------------")

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}
