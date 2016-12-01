package router

import (
	"net/http"

	"github.com/Flioh/DonemosApi/controlador"
	"github.com/Flioh/DonemosApi/middleware"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Router struct {
	*mux.Router

	controladorPing      *controlador.Ping
	controladorSolicitud *controlador.Solicitud
	controladorProvincia *controlador.Provincia
	controladorLocalidad *controlador.Localidad
	controladorBanco     *controlador.Banco
}

func (r Router) SubRouterPrefijo(prefijo string) *Router {
	r.Router = r.PathPrefix(prefijo).Subrouter()
	return &r
}

func NewRouter(
	controladorPing *controlador.Ping,
	controladorSolicitud *controlador.Solicitud,
	controladorProvincia *controlador.Provincia,
	controladorLocalidad *controlador.Localidad,
	controladorBanco *controlador.Banco) *Router {
	muxRouter := mux.NewRouter().StrictSlash(true)

	router := &Router{
		muxRouter,

		controladorPing,
		controladorSolicitud,
		controladorProvincia,
		controladorLocalidad,
		controladorBanco,
	}

	agregarRutas(router)
	agregarRutasV1(router.SubRouterPrefijo("/v1"))

	return router
}

func agregarRutasV1(router *Router) {
	agregarRutas(router)
}

func agregarRutas(router *Router) {
	for _, ruta := range GetRutas(router) {
		// var loggingHandler http.Handler
		// loggingHandler = helper.Logger(ruta.Handler, ruta.Nombre)

		rutaMux := router.
			Methods(ruta.Metodo).
			Name(ruta.Nombre).
			Path(ruta.Patron)

		finalHandler := negroni.New(
			negroni.HandlerFunc(middleware.Headers),
			negroni.HandlerFunc(middleware.Logger),
			negroni.Wrap(ruta.Handler),
		)
		if ruta.Seguro {
			setJwtMiddleware(rutaMux, finalHandler)
		} else {
			rutaMux.Handler(finalHandler)
		}
	}
}

func setJwtMiddleware(rutaMux *mux.Route, wrap http.Handler) {
	jwtMiddleware := middleware.Jwt()
	authMiddleware := negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(wrap),
	)
	rutaMux.Handler(authMiddleware)
}
