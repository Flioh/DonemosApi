package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/flioh/DonemosApi/modelos"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (c Controller) SolicitudIndex(w http.ResponseWriter, r *http.Request) {
	var solicitudes modelo.Solicitudes

	c.session.DB("donemos").C("solicitudes").Find(nil).All(&solicitudes)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(solicitudes); err != nil {
		panic(err)
	}

}

func SolicitudShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]
	fmt.Fprintln(w, "Solicitud Id:", id)
}

func (c Controller) SolicitudCreate(w http.ResponseWriter, r *http.Request) {
	s := modelo.Solicitud{}

	json.NewDecoder(r.Body).Decode(&s)

	s.SolicitudId = bson.NewObjectId()

	c.session.DB("donemos").C("solicitudes").Insert(s)

	sj, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", sj)
}
