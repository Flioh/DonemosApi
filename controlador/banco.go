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

	rango, rangoerr := strconv.ParseFloat(vars["rango"], 32)

	response := make(map[string]interface{})
	if lat1err != nil || lon1err != nil || rangoerr != nil {
		response["error"] = "Invalid coordinates."
	} else {

		var bancos modelo.Bancos
		var bancosEnDistancia modelo.Bancos
		q := c.db.Find(nil)
		q.All(&bancos)

		for _, banco := range bancos {
			lat2 := banco.Lat
			lon2 := banco.Lon
			distancia := helper.ObtenerDistancia(lat1, lon1, lat2, lon2)
			if distancia <= rango {
				bancosEnDistancia = append(bancosEnDistancia, banco)
			}
		}

		response["bancos"] = bancosEnDistancia
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
