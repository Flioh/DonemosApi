package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/flioh/DonemosApi/controlador"
	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/router"
)

func main() {
	fmt.Println("Iniciando servidor en puerto 8080")
	sesión := getSession()
	controller := controlador.NewSolicitud(db.NewDatabase(sesión, "solicitudes"))
	router := router.New(controller)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
