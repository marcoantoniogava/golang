package main

import "fmt"

func main() {
	canal := make(chan string, 200) //especifica que o canal tem uma capacidade de 2 mensagens
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
