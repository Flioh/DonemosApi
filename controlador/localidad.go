package controlador

import (
	"encoding/json"
	"net/http"

	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/modelo"
)

// Localidad es el controlador de localidades, contiene referencia a la sesion de mongodb
// y contiene los handlers relacionados a las solicitudes.
type Localidad struct {
	db *db.Database
}

func NewLocalidad(db *db.Database) *Localidad {
	return &Localidad{db}
}

func (c *Localidad) LocalidadIndex(w http.ResponseWriter, r *http.Request) {

	var provincias modelo.Localidades
	c.db.Todos().All(&provincias)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(provincias); err != nil {
		panic(err)
	}

}
