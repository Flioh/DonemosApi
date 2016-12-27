package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/bugsnag/bugsnag-go"

	"github.com/Flioh/DonemosApi/controlador"
	"github.com/Flioh/DonemosApi/db"
	"github.com/Flioh/DonemosApi/router"
)

func main() {
	fmt.Println("Iniciando servidor en puerto 8080")

	bugsnagApi := os.Getenv("BUGSNAG_API_KEY")

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       bugsnagApi,
		ReleaseStage: "development",
	})

	sesión := getSession()
	defer sesión.Close()
	dbSolicitudes := db.NewDatabase(sesión, "solicitudes")
	controladorPing := controlador.NewPing(sesión, dbSolicitudes)
	controladorSolicitudes := controlador.NewSolicitud(dbSolicitudes)
	controladorProvincias := controlador.NewProvincia(db.NewDatabase(sesión, "provincias"))
	controladorLocalidades := controlador.NewLocalidad(db.NewDatabase(sesión, "localidades"))
	controladorBancos := controlador.NewBanco(db.NewDatabase(sesión, "bancos"))
	router := router.NewRouter(
		controladorPing,
		controladorSolicitudes,
		controladorProvincias,
		controladorLocalidades,
		controladorBancos)

	//task.EmpezarTaskLimpieza(dbSolicitudes)
	log.Fatal(http.ListenAndServe(":8080", bugsnag.Handler(router)))
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
