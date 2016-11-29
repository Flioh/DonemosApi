package modelo

import (
	"encoding/json"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Geo struct {
	Kind  string    `json:"type" bson:"type"`
	Point []float64 `json:"coordinates" bson:"coordinates"`
}

type Banco struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	ProvinciaId bson.ObjectId `json:"provincia" bson:"provinciaId"`
	Loc         Geo           `json:"loc" bson:"loc"`
	Lat         float64       `json:"lat" bson:"lat"`
	Lon         float64       `json:"lon" bson:"lon"`
	Ciudad      string        `json:"ciudad" bson:"ciudad"`
	Institución string        `json:"nombre" bson:"nombre"`
	Dirección   string        `json:"direccion" bson:"direccion"`
	Teléfono    string        `json:"telefono" bson:"telefono"`
	Horario     string        `json:"horario" bson:"horario"`

	db *mgo.Database
}

type Bancos []Banco

func NewBanco(provinciaId bson.ObjectId,
	loc Geo, lat, lon float64, ciudad, institución, dirección, teléfono, horario string) *Banco {
	return &Banco{
		bson.NewObjectId(),
		provinciaId,
		loc, lat, lon,
		ciudad, institución, dirección, teléfono, horario, nil,
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

func (b Bancos) PrepararParaEncode(db *mgo.Database) Bancos {
	nuevas := make(Bancos, 0)
	for _, elem := range b {
		elem.db = db
		nuevas = append(nuevas, elem)
	}

	return nuevas
}

func (b *Banco) PrepararParaEncode(db *mgo.Database) {
	b.db = db
}

func (b *Banco) MarshalJSON() ([]byte, error) {
	type Alias Banco
	var provincia Provincia
	if b.db == nil {
		return nil, fmt.Errorf("No db reference: %v", b.db)
	}
	b.db.C("provincias").FindId(b.ProvinciaId).One(&provincia)
	return json.Marshal(&struct {
		Provincia Provincia `json:"provincia"`
		*Alias
	}{
		Alias:     (*Alias)(b),
		Provincia: provincia,
	})
}
