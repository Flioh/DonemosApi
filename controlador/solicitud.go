package controlador

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/modelo"
	"github.com/gorilla/mux"
)

// Solicitud es el controlador de solicitudes, contiene referencia a la sesion de mongodb
// y contiene los handlers relacionados a las solicitudes.
type Solicitud struct {
	db db.IColección
}

func NewSolicitud(db db.IColección) *Solicitud {
	return &Solicitud{db}
}

func (c *Solicitud) SolicitudIndex(w http.ResponseWriter, r *http.Request) {

	solicitudes, _ := c.db.Todos()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(solicitudes.(modelo.Solicitudes)); err != nil {
		panic(err)
	}

}

func (c *Solicitud) SolicitudShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]

	fmt.Println("Buscando id: ", id)

	solicitud, err := c.db.Read(id)
	if err != nil {
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

func (c *Solicitud) SolicitudCreate(w http.ResponseWriter, r *http.Request) {
	s := new(modelo.Solicitud)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	if err := json.Unmarshal(body, &s); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
	}

	c.db.Create(s)

	sj, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", sj)
}

func (c *Solicitud) SolicitudUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]

	nuevaSolicitud := new(modelo.Solicitud)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &nuevaSolicitud); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
	}

	err = c.db.Update(id, nuevaSolicitud)
	if err != nil {
		w.WriteHeader(404) //TODO: Probablemente el código apropiado sea otro.
		return
	}
	r.Body.Close()

	sj, _ := json.Marshal(nuevaSolicitud)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", sj)
}

func (c *Solicitud) SolicitudDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]

	err := c.db.Delete(id)
	if err != nil {
		w.WriteHeader(204)
		return
	}

	w.WriteHeader(202)
}
