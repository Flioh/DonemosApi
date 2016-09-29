package router

import (
	"net/http"

	"github.com/flioh/DonemosApi/controlador"
	"github.com/flioh/DonemosApi/helper"
	"github.com/gorilla/mux"
)

func New(controladorSolicitud *controlador.Solicitud,
	controladorProvincia *controlador.Provincia,
	controladorLocalidad *controlador.Localidad) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, ruta := range GetRutas(controladorSolicitud, controladorProvincia, controladorLocalidad) {
		var handler http.Handler
		handler = helper.Logger(ruta.Handler, ruta.Nombre)

		router.
			Methods(ruta.Metodo).
			Name(ruta.Nombre).
			Path(ruta.Patron).
			Handler(handler)
	}

	return router
}
