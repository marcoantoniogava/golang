package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string] string { //dentro dos [] vai o tipo das chaves, fora vai o tipo dos valores
		"nome": "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"]) //assim acessa um campo

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") //apaga um campo
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
