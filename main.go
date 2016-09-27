package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("Iniciando servidor.")
	controller := NewController(getSession())
	router := NewRouter(controller)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
