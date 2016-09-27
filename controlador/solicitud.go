package controlador

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/flioh/DonemosApi/modelos"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

// Solicitud es el controlador de solicitudes, contiene referencia a la sesion de mongodb
// y contiene los handlers relacionados a las solicitudes.
type Solicitud struct {
	session *mgo.Session
}

func NewSolicitud(s *mgo.Session) *Solicitud {
	return &Solicitud{s}
}

func (c *Solicitud) getDb() *mgo.Database {
	return c.session.DB("donemos")
}

func (c Solicitud) SolicitudIndex(w http.ResponseWriter, r *http.Request) {
	var solicitudes modelo.Solicitudes

	c.getDb().C("solicitudes").Find(nil).All(&solicitudes)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(solicitudes); err != nil {
		panic(err)
	}

}

func (c Solicitud) SolicitudShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]
	fmt.Fprintln(w, "Solicitud Id:", id)
}

func (c Solicitud) SolicitudCreate(w http.ResponseWriter, r *http.Request) {
	s := modelo.Solicitud{}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	if err := json.Unmarshal(body, &s); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
	}

	s.SolicitudId = bson.NewObjectId()

	c.getDb().C("solicitudes").Insert(s)

	sj, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", sj)
}
