package router

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

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

	agregarRutas(router, controladorSolicitud, controladorProvincia, controladorLocalidad)
	agregarRutasV1(router.PathPrefix("/v1").Subrouter(), controladorSolicitud, controladorProvincia, controladorLocalidad)

	return router
}

func agregarRutasV1(router *mux.Router, controladorSolicitud *controlador.Solicitud,
	controladorProvincia *controlador.Provincia,
	controladorLocalidad *controlador.Localidad) {
	agregarRutas(router, controladorSolicitud, controladorProvincia, controladorLocalidad)
}

func agregarRutas(router *mux.Router, controladorSolicitud *controlador.Solicitud,
	controladorProvincia *controlador.Provincia,
	controladorLocalidad *controlador.Localidad) {
	//router := mux.NewRouter().StrictSlash(true)

	for _, ruta := range GetRutas(controladorSolicitud, controladorProvincia, controladorLocalidad) {
		var loggingHandler http.Handler
		loggingHandler = helper.Logger(ruta.Handler, ruta.Nombre)

		rutaMux := router.
			Methods(ruta.Metodo).
			Name(ruta.Nombre).
			Path(ruta.Patron)

		if ruta.Patron == "POST" || ruta.Patron == "PUT" || ruta.Patron == "OPTIONS" {
			setHeaders(rutaMux)
		}

		if ruta.Seguro {
			setJwtMiddleware(rutaMux, loggingHandler)
		} else {
			rutaMux.Handler(loggingHandler)
		}
	}
}

func setHeaders(rutaMux *mux.Route) {
	rutaMux.Headers("Content-Type", "application/json")
	rutaMux.Headers("Access-Control-Allow-Origin", "*")
	rutaMux.Headers("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE, OPTIONS")
	rutaMux.Headers("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

func setJwtMiddleware(rutaMux *mux.Route, wrap http.Handler) {
	jwtMiddleware := jwtm.New(jwtm.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			decoded, err := base64.URLEncoding.DecodeString(os.Getenv("AUTH0_CLIENT_SECRET"))
			if err != nil {
				return nil, err
			}
			return decoded, nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, authErr string) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(401)
			err := map[string]interface{}{"error": authErr}
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	authMiddleware := negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(wrap),
	)
	rutaMux.Handler(authMiddleware)
}
