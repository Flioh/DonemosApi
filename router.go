package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(controller *Controller) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, ruta := range GetRutas(controller) {
		var handler http.Handler
		handler = Logger(ruta.Handler, ruta.Nombre)

		router.
			Methods(ruta.Metodo).
			Name(ruta.Nombre).
			Path(ruta.Patron).
			Handler(handler)
	}

	return router
}
