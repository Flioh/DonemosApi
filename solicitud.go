package main

import "time"

type Solicitud struct {
	SolicitudId     int       `json:"solicitudID"`
	UsuarioId       int       `json:"usuarioID"`
	Fecha           time.Time `json:"fechaCreacion"`
	Vigente         bool      `json:"estaVigente"`
	CantidadDadores int       `json:"cantidadDadores"`
	NombrePaciente  string    `json:"nombrePaciente"`
	Institucion     string    `json:"institucion"`
	Direccion       string    `json:"direccion"`
	HoraDesde       string    `json:"horaDesde"`
	HoraHasta       string    `json:"horaHasta"`
	Adicionales     string    `json:"datosAdicionales"`
	//Provincia Provincia
	//Ciudad Ciudad
	//GrupoSanguineo GrupoSanguineo
	//Factor Factor
}
type Solicitudes []Solicitud

func NewSolicitud(nombre, institucion string) *Solicitud {
	s := new(Solicitud)
	s.NombrePaciente = nombre
	s.Institucion = institucion
	return s
}
