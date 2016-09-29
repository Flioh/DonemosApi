package controlador

import (
	"encoding/json"
	"net/http"

	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/modelo"
)

// Provincia es el controlador de provincias, contiene referencia a la sesion de mongodb
// y contiene los handlers relacionados a las solicitudes.
type Provincia struct {
	db *db.Database
}

func NewProvincia(db *db.Database) *Provincia {
	return &Provincia{db}
}

func (c *Provincia) ProvinciaIndex(w http.ResponseWriter, r *http.Request) {

	var provincias modelo.Provincias
	c.db.Todos(100).All(&provincias)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(provincias); err != nil {
		panic(err)
	}

}
