package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//CRUD - CREATE, READ, UPDATE, DELETE

	router := mux.NewRouter()

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
	
}
