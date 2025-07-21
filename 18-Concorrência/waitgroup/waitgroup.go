package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4) //passa a quantidade de goroutines que vão entrar na fila

	go func()  { //goroutine 1
		escrever("Olá Mundo!")
		waitGroup.Done() //quando a goroutine terminar, decrementa o valor da WaitGroup
	}()

	go func()  { //goroutine 2
		escrever("Programando em Go!")
		waitGroup.Done() //quando a goroutine terminar, decrementa o valor da WaitGroup
	}()

	go func()  { //goroutine 3
		escrever("Goroutine 3!")
		waitGroup.Done() //quando a goroutine terminar, decrementa o valor da WaitGroup
	}()
	
	go func()  { //goroutine 4
		escrever("Goroutine 4!")
		waitGroup.Done() //quando a goroutine terminar, decrementa o valor da WaitGroup
	}()
	
	waitGroup.Wait() //espera até que todas as goroutines terminem
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
	fmt.Println(texto)
	time.Sleep(time.Second)
	}
}
