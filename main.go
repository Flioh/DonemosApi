package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"flioh/DonemosApi/controlador"
	"flioh/DonemosApi/db"
	"flioh/DonemosApi/router"
)

func main() {
	fmt.Println("Iniciando servidor en puerto 8080")
	sesión := getSession()
	controladorSolicitudes := controlador.NewSolicitud(db.NewDatabase(sesión, "solicitudes"))
	controladorProvincias := controlador.NewProvincia(db.NewDatabase(sesión, "provincias"))
	controladorLocalidades := controlador.NewLocalidad(db.NewDatabase(sesión, "localidades"))
	routerHttp := router.New(controladorSolicitudes, controladorProvincias, controladorLocalidades)
	log.Fatal(http.ListenAndServe(":8080", routerHttp))
}

func getSession() *mgo.Session {
	mongoUrl := os.Getenv("MONGO_URL")

	if mongoUrl == "" {
		// TODO:
		fmt.Println("Variable de entorno 'MONGO_URL' no proporcionada. Usando localhost")
		mongoUrl = "mongodb://localhost"
	}

	s, err := mgo.Dial(mongoUrl)

	if err != nil {
		panic(err)
	}
	return s
}
