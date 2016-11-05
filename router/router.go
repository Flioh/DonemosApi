package router

import (
	"net/http"
	"fmt"
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

		rutaMux := router.
			Methods(ruta.Metodo).
			Name(ruta.Nombre).
			Path(ruta.Patron).
			Handler(handler)

		if ruta.Patron == "POST" || ruta.Patron == "PUT" || ruta.Patron == "OPTIONS" {
			rutaMux.Headers("Content-Type", "application/json")
			rutaMux.Headers("Access-Control-Allow-Origin", "*")
			rutaMux.Headers("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE, OPTIONS")
			rutaMux.Headers("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		}
	}

	return router
}





/*package router

import (
	"net/http"
	"fmt"
	jwtm "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/flioh/DonemosApi/controlador"
	"github.com/flioh/DonemosApi/helper"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func New(controladorSolicitud *controlador.Solicitud,
	controladorProvincia *controlador.Provincia,
	controladorLocalidad *controlador.Localidad) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, ruta := range GetRutas(controladorSolicitud, controladorProvincia, controladorLocalidad) {
		var loggingHandler http.Handler
		loggingHandler = helper.Logger(ruta.Handler, ruta.Nombre)

		rutaMux := router.
			Methods(ruta.Metodo).
			Name(ruta.Nombre).
			Path(ruta.Patron)

		if ruta.Patron == "POST" || ruta.Patron == "PUT" || ruta.Patron == "OPTIONS" {
			fmt.Println(ruta.Patron)
			rutaMux.Headers("Content-Type", "application/json")
			rutaMux.Headers("Access-Control-Allow-Origin", "*")
			rutaMux.Headers("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE, OPTIONS")
			rutaMux.Headers("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		}

		/*if ruta.Seguro {
			jwtMiddleware := jwtm.New(jwtm.Options{
				ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
					return []byte("Secret"), nil
				},
				SigningMethod: jwt.SigningMethodHS256,
			})
			authMiddleware := negroni.New(
				negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
				negroni.Wrap(loggingHandler),
			)
			rutaMux.Handler(authMiddleware)
		} else {
			rutaMux.Handler(loggingHandler)
		}
	}

	return router
}
*/