package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando servidor.")
	controller := NewController()
	router := NewRouter(controller)
	log.Fatal(http.ListenAndServe(":8080", router))
}
