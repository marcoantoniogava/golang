package main

import "fmt"

func soma(numeros ...int) int { //... diz que a função vai receber de 1 a n ints
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int)  {
	for _, numero := range numeros {
		fmt.Printf(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}
