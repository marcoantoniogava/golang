package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5] string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...] int{1,2,3,4,5} //ao usar ... o array se adapta à quantidade de elementos
	fmt.Println(array3)

	slice := [] int{10, 11, 12, 13, 14, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18) //append pega o slice e adiciona um novo elemento, retornando o novo slice
	fmt.Println(slice)

	slice2 := array2[1:3] //fatiando o array, o array fica inalterado mas é criado um slice com os elementos de 1 a 3 do array
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("-------------------")
	slice3 := make([]float32, 10, 11) //make cria um slice com a quantidade de elementos e o tamanho máximo
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5) //slice na capacidade maxima
	slice3 = append(slice3, 6) //estoura o limite, cria um novo slice com o dobro da capacidade

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5) //se não especificar a capacidade, a capacidade será igual ao tamanho
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
