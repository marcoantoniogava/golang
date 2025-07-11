package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" //1° maneira de declarar uma variável
	variavel2 := "Variável 2" //2° maneira de declarar uma variável

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( //declarando duas ou mais variaáveis de uma vez
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6" //declarando duas ou mais variaáveis de uma vez
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1" //declarando variável com const
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 //invertendo o valor das variáveis 5 e 6
	fmt.Println(variavel5, variavel6)
}
