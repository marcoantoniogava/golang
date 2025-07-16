package main

import "fmt"

func recuperarExecucao()  {
	if r := recover(); r != nil { //recover recupera a execução do programa
		fmt.Println("Recuperando execução...")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A Média é exatamente 6") //caso a condição seja verdadeira, o programa em panico e para.
}

func main() {
	fmt.Println(alunoEstaAprovado(6,7))
	fmt.Println("Pós execução!")
}
