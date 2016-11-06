package modelo

import "gopkg.in/mgo.v2/bson"

type Provincia struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Nombre string        `json:"nombre" bson:"nombre"`
}

type Provincias []Provincia

func NewProvincia(nombre string) *Provincia {
	return &Provincia{bson.NewObjectId(), nombre}
}

func (s Provincias) Len() int {
	return len(s)
}

func (s Provincias) Less(i, j int) bool {
	return s[i].Nombre < s[j].Nombre
}

func (s Provincias) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
