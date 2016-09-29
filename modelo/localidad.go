package modelo

import "gopkg.in/mgo.v2/bson"

type Localidad struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Nombre string        `json:"nombre" bson:"nombre"`
}

type Localidades []Localidad

func NewLocalidad(nombre string, localidadId bson.ObjectId) *Localidad {
	return &Localidad{bson.NewObjectId(), nombre}
}
