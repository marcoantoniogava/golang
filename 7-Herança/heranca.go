package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint
	altura uint
}

type estudante struct { //pega todos os campos que estão em "pessoa" e joga em estudante
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa {"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante {p1, "Engenharia", "USP"} //passando todos os campos de pessoa para estudante
	fmt.Println(e1)
	fmt.Println(e1.nome) //acessa apenas o nome
}
