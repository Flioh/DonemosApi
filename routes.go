package main

import "net/http"

type Ruta struct {
	Nombre  string
	Metodo  string
	Patron  string
	Handler http.HandlerFunc
}
type Rutas []Ruta

func GetRutas(controller *Controller) Rutas {
	return Rutas{
		Ruta{
			"Index",
			"GET",
			"/",
			Index,
		},
		Ruta{
			"SolicitudIndex",
			"GET",
			"/solicitud",
			controller.SolicitudIndex,
		},
		Ruta{
			"SolicitudShow",
			"GET",
			"/solicitud/{solicitudId}",
			SolicitudShow,
		},
	}
}
