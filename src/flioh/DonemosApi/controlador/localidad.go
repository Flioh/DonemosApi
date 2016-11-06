package controlador

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"flioh/DonemosApi/db"
	"flioh/DonemosApi/modelo"
	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	id := vars["provinciaId"]

	var localidades modelo.Localidades
	err := c.db.Colección().Find(bson.M{"provinciaId": bson.ObjectIdHex(id)}).All(&localidades)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(localidades); err != nil {
		panic(err)
	}

}
