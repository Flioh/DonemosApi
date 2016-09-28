package modelo

import "gopkg.in/mgo.v2/bson"

type IModelo interface {
	GetId() bson.ObjectId
	SetId(bson.ObjectId)
	SetIdHex(string)
}
