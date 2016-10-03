package db

import mgo "gopkg.in/mgo.v2"

type Query struct {
	*mgo.Query
}

func NewQuery(q *mgo.Query) *Query {
	return &Query{q}
}

func (q *Query) Paginar(página int) *Query {
	por_página := 2
	q.Skip((página - 1) * por_página).Limit(por_página)
	return q
}
