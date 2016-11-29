package controlador

import (
	"fmt"
	"net/http"

	"github.com/flioh/DonemosApi/db"

	mgo "gopkg.in/mgo.v2"
)

type Ping struct {
	sesión *mgo.Session
	db     *db.Database
}

func NewPing(s *mgo.Session, d *db.Database) *Ping {
	return &Ping{s, d}
}

func (c *Ping) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping success")
}

func (c *Ping) Health(w http.ResponseWriter, r *http.Request) {

	err := c.sesión.Ping()

	if err != nil {
		fmt.Fprintf(w, "Error pinging mgo server")
		return
	}

	c.db.Find(nil).Limit(1)

	fmt.Fprintf(w, "Health success")
}
