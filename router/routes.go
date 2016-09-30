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

func GetRutas(cs *controlador.Solicitud, cp *controlador.Provincia, cl *controlador.Localidad) Rutas {
	return Rutas{
		Ruta{
			"Index",
			"GET",
			"/solicitud",
			cs.SolicitudIndex,
		},
		Ruta{
			"SolicitudIndex",
			"GET",
			"/solicitud/filtrar/{provinciaId}/{localidadId}/{grupoId}/{factorId}",
			cs.SolicitudIndex,
		},
		Ruta{
			"SolicitudCreate",
			"POST",
			"/solicitud",
			cs.SolicitudCreate,
		},
		Ruta{
			"SolicitudShow",
			"GET",
			"/solicitud/{solicitudId}",
			cs.SolicitudShow,
		},
		Ruta{
			"SolicitudUpdate",
			"PUT",
			"/solicitud/{solicitudId}",
			cs.SolicitudUpdate,
		},
		Ruta{
			"SolicitudDelete",
			"DELETE",
			"/solicitud/{solicitudId}",
			cs.SolicitudDelete,
		},

		Ruta{
			"ProvinciaIndex",
			"GET",
			"/provincia",
			cp.ProvinciaIndex,
		},

		Ruta{
			"LocalidadIndex",
			"GET",
			"/localidad/{provinciaId}",
			cl.LocalidadIndex,
		},
	}
}
