package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") //goroutine, faz a função escrever rodar em paralelo, sem bloquear o programa
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
	fmt.Println(texto)
	time.Sleep(time.Second)
	}
}
