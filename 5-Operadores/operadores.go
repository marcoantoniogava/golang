package main

import "fmt"

func main() {
	//aritmeticos
	soma := 10 + 5
	subtracao := 10 - 5
	multiplicacao := 10 * 5
	divisao := 10 / 5
	restoDaDivisao := 10 % 5

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//fim dos aritmeticos

	//atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//fim dos operadores de atribuição

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//fim dos relacionais

	//operadores lógicos
	fmt.Println("--------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //e
	fmt.Println(verdadeiro || falso) //ou
	fmt.Println(!verdadeiro) //negação, inverte
	fmt.Println(!falso) //negação, inverte
	//fim dos operadores lógicos

	//operadores unários
	fmt.Println("--------------------")
	numero := 10
	numero++ //aumenta 1
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero-- //diminui 1
	numero -= 20 //numero = numero - 20

	numero *= 3 //numero = numero * 3
	numero /= 10 //numero = numero / 10
	numero %= 3 //numero = numero % 3

	fmt.Println(numero)
	//fim dos operadores unários
}
