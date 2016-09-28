package db

import (
	"github.com/flioh/DonemosApi/modelo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Solicitudes struct {
	sesión *mgo.Session
}

func NewSolicitudes(sesión *mgo.Session) *Solicitudes {
	return &Solicitudes{sesión}
}

func (s *Solicitudes) colección() *mgo.Collection {
	return s.sesión.DB("donemos").C("solicitudes")
}

func (s *Solicitudes) Todos() (modelo.Solicitudes, error) {
	var solicitudes modelo.Solicitudes

	//solicitudes := c.db.solicitudes.findAll()
	//solicitudes := c.collection.findAll()
	err := s.colección().Find(nil).All(&solicitudes)

	if err != nil {
		return nil, err
	}
	return solicitudes, nil
}

func (s *Solicitudes) Create(solicitud modelo.Solicitud) error {
	solicitud.SolicitudId = bson.NewObjectId()
	err := s.colección().Insert(solicitud)

	return err
}

func (s *Solicitudes) Read(hexId string) (solicitud modelo.Solicitud, err error) {
	objectId := bson.ObjectIdHex(hexId)
	err = s.colección().FindId(objectId).One(&solicitud)

	return
}

func (s *Solicitudes) Update(hexId string, solicitud modelo.Solicitud) error {
	id := bson.ObjectIdHex(hexId)
	solicitud.SolicitudId = id
	return s.colección().UpdateId(bson.ObjectIdHex(hexId), solicitud)
}
