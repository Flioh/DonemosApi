package modelo

import "gopkg.in/mgo.v2/bson"

type Banco struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	ProvinciaId bson.ObjectId `json:"provincia" bson:"provinciaId"`
	CiudadId    bson.ObjectId `json:"ciudad" bson:"ciudadId"`
	Lat         float32       `json:"lat" bson:"lat"`
	Lon         float32       `json:"lon" bson:"lon"`
	Institución string        `json:"nombre" bson:"nombre"`
	Dirección   string        `json:"direccion" bson:"direccion"`
	Teléfono    string        `json:"telefono" bson:"telefono"`
	Horario     string        `json:"horario" bson:"horario"`
}

type Bancos []Banco

func NewBanco(provinciaId, ciudadId bson.ObjectId,
	lat, lon float32, institución, dirección, teléfono, horario string) *Banco {
	return &Banco{
		bson.NewObjectId(),
		provinciaId,
		ciudadId,
		lat, lon,
		institución, dirección, teléfono, horario,
	}
}
