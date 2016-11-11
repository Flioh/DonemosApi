package modelo

import (
	"encoding/json"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Solicitud struct {
	SolicitudId     bson.ObjectId   `json:"solicitudID" bson:"_id"`
	UsuarioId       int             `json:"usuarioID" bson:"usuarioID"`
	Fecha           time.Time       `json:"fechaCreacion" bson:"fecha"`
	Vigente         bool            `json:"estaVigente" bson:"estaVigente"`
	CantidadDadores int             `json:"cantidadDadores" bson:"cantidadDadores"`
	NombrePaciente  string          `json:"nombrePaciente" bson:"nombrePaciente"`
	DNIPaciente     int             `json:"pacienteDNI" bson:"pacienteDNI"`
	Institucion     string          `json:"institucion" bson:"institucion"`
	Direccion       string          `json:"direccion" bson:"direccion"`
	HoraDesde       string          `json:"horaDesde" bson:"horaDesde"`
	HoraHasta       string          `json:"horaHasta" bson:"horaHasta"`
	Adicionales     string          `json:"datosAdicionales" bson:"datosAdicionales"`
	Grupo           GrupoSanguineo  `json:"grupoSanguineo" bson:"grupoSanguineo"`
	Factor          FactorSanguineo `json:"factorSanguineo" bson:"factorSanguineo"`
	ProvinciaId     bson.ObjectId   `json:"provinciaId" bson:"provinciaId"`
	CiudadId        bson.ObjectId   `json:"localidadId" bson:"localidadId"`

	db *mgo.Database
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

func (s Solicitudes) PrepararParaEncode(db *mgo.Database) Solicitudes {
	nuevas := make(Solicitudes, 0)
	for _, elem := range s {
		elem.db = db
		nuevas = append(nuevas, elem)
	}
	return nuevas
}

func (s *Solicitud) PrepararParaEncode(db *mgo.Database) {
	s.db = db
}

func (s *Solicitud) MarshalJSON() ([]byte, error) {
	type Alias Solicitud
	var ciudad Localidad
	var provincia Provincia
	if s.db == nil {
		return nil, fmt.Errorf("No db reference: %v", s.db)
	}
	s.db.C("provincias").FindId(s.ProvinciaId).One(&provincia)
	s.db.C("localidades").FindId(s.CiudadId).One(&ciudad)
	return json.Marshal(&struct {
		Provincia Provincia `json:"provincia"`
		Localidad Localidad `json:"ciudad"`
		*Alias
	}{
		Alias:     (*Alias)(s),
		Localidad: ciudad,
		Provincia: provincia,
	})
}

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
