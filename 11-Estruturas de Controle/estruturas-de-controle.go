package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("É menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { //declarando uma variável dentro do if
		fmt.Println("É maior que 0")
	} else if outroNumero < 10 {
		fmt.Println("É menor que 10")
	} else {
		fmt.Println("É igual a 10")
	}

	//fmt.Println(outroNumero) // não pode ser acessado fora do if

}
