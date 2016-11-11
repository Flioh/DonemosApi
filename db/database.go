package db

import (
	"fmt"

	"github.com/flioh/DonemosApi/modelo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	sesión          *mgo.Session
	nombreColección string
}

func NewDatabase(sesión *mgo.Session, n string) *Database {
	return &Database{sesión, n}
}

func (db *Database) GetMongoDB() *mgo.Database {
	return db.sesión.DB("donemos")
}

func (db *Database) Colección() *mgo.Collection {
	return db.sesión.DB("donemos").C(db.nombreColección)
}

func (db *Database) Find(query bson.M) *Query {

	return NewQuery(db.Colección().Find(query).Sort("-_id"))
}

func (db *Database) FindNear(lat, lon, rango float64) *Query {
	return NewQuery(db.Colección().Find(
		bson.M{
			"loc": bson.M{
				"$near": bson.M{
					"$geometry":    bson.M{"type": "Point", "coordinates": []float64{lon, lat}},
					"$minDistance": 0,
					"$maxDistance": rango,
				},
			},
		},
	))
}

func (db *Database) Create(m modelo.IModelo) error {
	id := bson.NewObjectId()
	fmt.Printf("GEN ID:%v\n", id)
	m.SetId(id)
	err := db.Colección().Insert(m)

	return err
}

func (db *Database) Read(hexId string) (m modelo.IModelo, err error) {
	if !bson.IsObjectIdHex(hexId) {
		return nil, fmt.Errorf("hexId invalido %v", hexId)
	}
	objectId := bson.ObjectIdHex(hexId)
	err = db.Colección().FindId(objectId).One(&m)

	return
}

func (db *Database) Update(hexId string, m modelo.IModelo) error {
	m.SetIdHex(hexId)
	return db.Colección().UpdateId(m.GetId(), m)
}

func (db *Database) Delete(hexId string) error {
	return db.Colección().RemoveId(bson.ObjectIdHex(hexId))
}
