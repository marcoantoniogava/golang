package main

import "fmt"

func main() {
	var numero int = 1000000000
	fmt.Println(numero)

	var numero2 uint = 10000 //uint apenas inteiros positivos
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	//fim numeros inteiros

	//numeros reais

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//fim numeros reais

	//strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A' //retorna o numero do caracter na tabela asc
	fmt.Println(char)

	//fim strings

	var texto string //todo tipo de dado tem seu valor inicial, por exemplo: para string = vazio (valor em branco), int = 0, bool = false, erros = nil (nulo)
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
}
