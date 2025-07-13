package main

import "fmt"

func somar(n1 int, n2 int) int { //parametros n1 e n2 int, e int após os parametros significa que vai retornar um valor int
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) { //da 2 retornos
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20) //chamando a função somar com os valores
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// resultadoSoma, _ := calculosMatematicos(10, 15) //aqui, ao colocar _, significa que não quero o valor da subtracao, apenas uma parte da função
	// fmt.Println(resultadoSoma, resultadoSubtracao)
}
