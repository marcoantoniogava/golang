package main

import "fmt"

type usuario struct { //cria o tipo usuario, que é uma struct
	nome string
	idade uint
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario //variavel "u", do tipo "usuario"
	u.nome = "Davi" //atribui valor a variavel "nome"
	u.idade = 21 //atribui valor a variavel "idade"
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Davi", 21, enderecoExemplo} //2° forma de atribuir
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"} //preenchendo apenas um campo
	fmt.Println(usuario3)
}
