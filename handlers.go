package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (c Controller) SolicitudIndex(w http.ResponseWriter, r *http.Request) {
	solicitudes := Solicitudes{
		*NewSolicitud("John Doe", "Iturraspe"),
		*NewSolicitud("Mary Jane", "Over the rainbow"),
	}

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
