package modelo

import "gopkg.in/mgo.v2/bson"

type Provincia struct {
	Id     bson.ObjectId `json:"id"`
	Nombre string        `json:"nombre"`
}

type Provincias []Provincia

func NewProvincia(nombre string) *Provincia {
	return &Provincia{bson.NewObjectId(), nombre}
}
