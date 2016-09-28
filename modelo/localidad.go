package modelo

import "gopkg.in/mgo.v2/bson"

type Localidad struct {
	Id          bson.ObjectId `json:"id"`
	ProvinciaId bson.ObjectId `json:"provinciaId"`
	Nombre      string        `json:"nombre"`
}

func NewLocalidad(nombre string, provinciaId bson.ObjectId) *Localidad {
	return &Localidad{bson.NewObjectId(), provinciaId, nombre}
}
