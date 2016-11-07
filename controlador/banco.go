package controlador

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/helper"
	"github.com/flioh/DonemosApi/modelo"
	"github.com/gorilla/mux"
)

type Banco struct {
	db *db.Database
}

func NewBanco(db *db.Database) *Banco {
	return &Banco{db}
}

func (c *Banco) BancoIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["provinciaId"]

	var bancos modelo.Bancos
	err := c.db.Colección().Find(bson.M{"provinciaId": bson.ObjectIdHex(id)}).All(&bancos)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(bancos); err != nil {
		panic(err)
	}

}

func (c *Banco) BancoDistancia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lat1, lat1err := strconv.ParseFloat(vars["lat1"], 32)
	lon1, lon1err := strconv.ParseFloat(vars["lon1"], 32)

	lat2, lat2err := strconv.ParseFloat(vars["lat2"], 32)
	lon2, lon2err := strconv.ParseFloat(vars["lon2"], 32)

	rango, rangoerr := strconv.ParseFloat(vars["rango"], 32)

	if lat1err != nil || lon1err != nil || lat2err != nil ||
		lon2err != nil || rangoerr != nil {

	} else {
		distancia := helper.ObtenerDistancia(lat1, lon1, lat2, lon2)

		if distancia <= rango {

		}
	}
}
