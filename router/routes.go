package router

import (
	"net/http"

	"github.com/flioh/DonemosApi/controlador"
)

type Ruta struct {
	Nombre  string
	Metodo  string
	Patron  string
	Seguro  bool
	Handler http.HandlerFunc
}
type Rutas []Ruta

func GetRutas(cs *controlador.Solicitud, cp *controlador.Provincia, cl *controlador.Localidad) Rutas {
	return Rutas{
		Ruta{
			"Ping",
			"GET",
			"/ping",
			true,
			cs.Ping,
		},
		Ruta{
			"Index",
			"GET",
			"/solicitud",
			false,
			cs.SolicitudIndex,
		},
		Ruta{ // Solicitudes paginadas
			"Index",
			"GET",
			"/solicitud/{pag}",
			false,
			cs.SolicitudIndex,
		},
		Ruta{ // Filtro de solicitudes
			"SolicitudIndex",
			"GET",
			"/solicitud/filtrar/{provinciaId}/{localidadId}/{grupoId}/{factorId}",
			false,
			cs.SolicitudIndex,
		},
		Ruta{ // Solicitudes filtradas y paginadas
			"SolicitudIndex",
			"GET",
			"/solicitud/{pag}/filtrar/{provinciaId}/{localidadId}/{grupoId}/{factorId}",
			false,
			cs.SolicitudIndex,
		},
		Ruta{
			"SolicitudCreate",
			"POST",
			"/solicitud",
			true,
			cs.SolicitudCreate,
		},
		Ruta{
			"SolicitudPreflightCheck",
			"OPTIONS",
			"/solicitud",
			false,
			cs.SolicitudPreflight,
		},
		Ruta{
			"SolicitudPreflightCheck",
			"OPTIONS",
			"/solicitud/{solicitudId}",
			false,
			cs.SolicitudPreflight,
		},
		Ruta{
			"SolicitudShow",
			"GET",
			"/solicitud/{solicitudId}",
			false,
			cs.SolicitudShow,
		},
		Ruta{
			"SolicitudUpdate",
			"PUT",
			"/solicitud/{solicitudId}",
			true,
			cs.SolicitudUpdate,
		},
		Ruta{
			"SolicitudDelete",
			"DELETE",
			"/solicitud/{solicitudId}",
			true,
			cs.SolicitudDelete,
		},

		Ruta{
			"ProvinciaIndex",
			"GET",
			"/provincia",
			false,
			cp.ProvinciaIndex,
		},

		Ruta{
			"LocalidadIndex",
			"GET",
			"/localidad/{provinciaId}",
			false,
			cl.LocalidadIndex,
		},
	}
}
