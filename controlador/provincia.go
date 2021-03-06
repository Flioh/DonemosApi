package controlador

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/Flioh/DonemosApi/db"
	"github.com/Flioh/DonemosApi/modelo"
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
	c.db.Find(nil).All(&provincias)

	sort.Sort(provincias)

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(provincias); err != nil {
		panic(err)
	}

}
