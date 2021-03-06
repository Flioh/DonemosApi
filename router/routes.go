package router

import "net/http"

type Ruta struct {
	Nombre  string
	Metodo  string
	Patron  string
	Seguro  bool
	Handler http.HandlerFunc
}

func (ruta *Ruta) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ruta.Handler(w, r)
}

type Rutas []Ruta

func GetRutas(r *Router) Rutas {
	return Rutas{
		Ruta{
			"Ping",
			"GET",
			"/ping",
			false,
			r.controladorPing.Ping,
		},
		Ruta{
			"MongoPing",
			"GET",
			"/mongoping",
			false,
			r.controladorPing.Health,
		},
		Ruta{
			"Index",
			"GET",
			"/solicitud",
			false,
			r.controladorSolicitud.SolicitudIndex,
		},
		Ruta{ // Solicitudes paginadas
			"Index",
			"GET",
			"/solicitud/{pag}",
			false,
			r.controladorSolicitud.SolicitudIndex,
		},
		Ruta{ // Filtro de solicitudes
			"SolicitudIndex",
			"GET",
			"/solicitud/filtrar/{provinciaId}/{localidadId}/{grupoId}/{factorId}",
			false,
			r.controladorSolicitud.SolicitudIndex,
		},
		Ruta{ // Solicitudes filtradas y paginadas
			"SolicitudIndex",
			"GET",
			"/solicitud/{pag}/filtrar/{provinciaId}/{localidadId}/{grupoId}/{factorId}",
			false,
			r.controladorSolicitud.SolicitudIndex,
		},
		Ruta{
			"SolicitudUsuario",
			"GET",
			"/solicitud/usuario/{usuarioId}",
			false,
			r.controladorSolicitud.SolicitudUsuario,
		},
		Ruta{
			"SolicitudCreate",
			"POST",
			"/solicitud",
			true,
			r.controladorSolicitud.SolicitudCreate,
		},
		Ruta{
			"SolicitudPreflightCheck",
			"OPTIONS",
			"/solicitud",
			false,
			r.controladorSolicitud.SolicitudPreflight,
		},
		Ruta{
			"SolicitudPreflightCheck",
			"OPTIONS",
			"/solicitud/{solicitudId}",
			false,
			r.controladorSolicitud.SolicitudPreflight,
		},
		Ruta{
			"SolicitudShow",
			"GET",
			"/solicitud/{solicitudId}",
			false,
			r.controladorSolicitud.SolicitudShow,
		},
		Ruta{
			"SolicitudUpdate",
			"PUT",
			"/solicitud/{solicitudId}",
			true,
			r.controladorSolicitud.SolicitudUpdate,
		},
		Ruta{
			"SolicitudDelete",
			"DELETE",
			"/solicitud/{solicitudId}",
			true,
			r.controladorSolicitud.SolicitudDelete,
		},

		Ruta{
			"ProvinciaIndex",
			"GET",
			"/provincia",
			false,
			r.controladorProvincia.ProvinciaIndex,
		},

		Ruta{
			"LocalidadIndex",
			"GET",
			"/localidad/{provinciaId}",
			false,
			r.controladorLocalidad.LocalidadIndex,
		},
		Ruta{
			"BancoProvincia",
			"GET",
			"/banco/{provinciaId}",
			false,
			r.controladorBanco.BancoIndex,
		},
		Ruta{
			"BancoDistancia",
			"GET",
			"/banco/{lat}/{lon}/{rango}",
			false,
			r.controladorBanco.BancoDistancia,
		},
	}
}
