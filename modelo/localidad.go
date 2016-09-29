package modelo

import "gopkg.in/mgo.v2/bson"

type Localidad struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	ProvinciaId bson.ObjectId `json:"provincia" bson:"provinciaId"`
	Nombre      string        `json:"nombre" bson:"nombre"`
}

type Localidades []Localidad

func NewLocalidad(nombre string, provinciaId bson.ObjectId, localidadId bson.ObjectId) *Localidad {
	return &Localidad{bson.NewObjectId(), provinciaId, nombre}
}
