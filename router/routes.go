package router

import (
	"net/http"

	"github.com/flioh/DonemosApi/controlador"
)

type Ruta struct {
	Nombre  string
	Metodo  string
	Patron  string
	Handler http.HandlerFunc
}
type Rutas []Ruta

func GetRutas(solicitudController *controlador.Solicitud) Rutas {
	return Rutas{
		Ruta{
			"Index",
			"GET",
			"/",
			solicitudController.SolicitudIndex,
		},
		Ruta{
			"SolicitudIndex",
			"GET",
			"/solicitud",
			solicitudController.SolicitudIndex,
		},
		Ruta{
			"SolicitudCreate",
			"POST",
			"/solicitud",
			solicitudController.SolicitudCreate,
		},
		Ruta{
			"SolicitudShow",
			"GET",
			"/solicitud/{solicitudId}",
			solicitudController.SolicitudShow,
		},
		Ruta{
			"SolicitudUpdate",
			"PUT",
			"/solicitud/{solicitudId}",
			solicitudController.SolicitudUpdate,
		},
	}
}
