package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal //Espera a mensagem ser escrita no canal
		if !aberto { //2° parametro verifica se o canal está aberto
			break
		}
		fmt.Println(mensagem)
	}

	// for mensagem := range canal { //2° maneira de receber mensagem do canal
	// 	fmt.Println(mensagem)
	// }

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) //depois de terminar o loop de 5 vezes, fecha o canal
}
