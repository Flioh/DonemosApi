package controlador

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/flioh/DonemosApi/modelos"
	"github.com/gorilla/mux"
)

// Solicitud es el controlador de solicitudes, contiene referencia a la sesion de mongodb
// y contiene los handlers relacionados a las solicitudes.
type Solicitud struct {
	session *mgo.Session
}

func NewSolicitud(s *mgo.Session) *Solicitud {
	return &Solicitud{s}
}

func (c *Solicitud) collection() *mgo.Collection {
	return c.session.DB("donemos").C("solicitudes")
}

func (c Solicitud) SolicitudIndex(w http.ResponseWriter, r *http.Request) {
	var solicitudes modelo.Solicitudes

	c.collection().Find(nil).All(&solicitudes)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(solicitudes); err != nil {
		panic(err)
	}

}

func (c Solicitud) SolicitudShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["solicitudId"])

	fmt.Println("Buscando id: ", id)

	var solicitud modelo.Solicitud
	if err := c.collection().FindId(id).One(&solicitud); err != nil {
		// Handle not found
		w.WriteHeader(204)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(solicitud); err != nil {
		panic(err)
	}
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

	c.collection().Insert(s)

	sj, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", sj)
}
