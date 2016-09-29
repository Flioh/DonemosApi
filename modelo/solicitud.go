package modelo

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Solicitud struct {
	SolicitudId     bson.ObjectId   `json:"solicitudID" bson:"_id"`
	UsuarioId       int             `json:"usuarioID" bson:"usuarioID"`
	Fecha           time.Time       `json:"fechaCreacion" bson:"fecha"`
	Vigente         bool            `json:"estaVigente" bson:"estaVigente"`
	CantidadDadores int             `json:"cantidadDadores" bson:"cantidadDadores"`
	NombrePaciente  string          `json:"nombrePaciente" bson:"nombrePaciente"`
	Institucion     string          `json:"institucion" bson:"institucion"`
	Direccion       string          `json:"direccion" bson:"direccion"`
	HoraDesde       string          `json:"horaDesde" bson:"horaDesde"`
	HoraHasta       string          `json:"horaHasta" bson:"horaHasta"`
	Adicionales     string          `json:"datosAdicionales" bson:"datosAdicionales"`
	Grupo           GrupoSanguineo  `json:"grupoSanguineo" bson:"grupoSanguineo"`
	Factor          FactorSanguineo `json:"factorSanguineo" bson:"factorSanguineo"`
	ProvinciaId     bson.ObjectId   `bson: "provinciaId"`
	CiudadId        bson.ObjectId   `bson:"provinciaId"`
}
type Solicitudes []Solicitud

func NewSolicitud(nombre, institucion string) *Solicitud {
	s := new(Solicitud)
	s.NombrePaciente = nombre
	s.Institucion = institucion
	return s
}

func (s *Solicitud) GetId() bson.ObjectId {
	return s.SolicitudId
}

func (s *Solicitud) SetIdHex(hex string) {
	s.SolicitudId = bson.ObjectIdHex(hex)
}

func (s *Solicitud) SetId(id bson.ObjectId) {
	s.SolicitudId = id
}

// func (s *Solicitud) MarshalJSON() ([]byte, error) {
// 	type Alias Solicitud
// 	return json.Marshal(&struct{
// 		Provincia bson.ObjectId `json:""`
// 	})
// }

func (s *Solicitud) UnmarshalJSON(data []byte) error {
	type Alias Solicitud
	aux := &struct {
		Provincia Provincia `json:"provincia"`
		Localidad Localidad `json:"ciudad"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.ProvinciaId = aux.Provincia.Id
	s.CiudadId = aux.Localidad.Id
	return nil
}
