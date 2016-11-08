package modelo

import "gopkg.in/mgo.v2/bson"

type Banco struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	ProvinciaId bson.ObjectId `json:"provincia" bson:"provinciaId"`
	Lat         float64       `json:"lat" bson:"lat"`
	Lon         float64       `json:"lon" bson:"lon"`
	Ciudad      string        `json:"ciudad" bson:"ciudad"`
	Institución string        `json:"nombre" bson:"nombre"`
	Dirección   string        `json:"direccion" bson:"direccion"`
	Teléfono    string        `json:"telefono" bson:"telefono"`
	Horario     string        `json:"horario" bson:"horario"`
}

type Bancos []Banco

func NewBanco(provinciaId bson.ObjectId,
	lat, lon float64, ciudad, institución, dirección, teléfono, horario string) *Banco {
	return &Banco{
		bson.NewObjectId(),
		provinciaId,
		lat, lon,
		ciudad, institución, dirección, teléfono, horario,
	}
}

func (b *Banco) GetId() bson.ObjectId {
	return b.Id
}

func (b *Banco) SetIdHex(hex string) {
	b.Id = bson.ObjectIdHex(hex)
}

func (b *Banco) SetId(id bson.ObjectId) {
	b.Id = id
}
