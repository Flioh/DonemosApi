package controlador

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/Flioh/DonemosApi/db"
	"github.com/Flioh/DonemosApi/helper"
	"github.com/Flioh/DonemosApi/modelo"
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

	if !helper.IdValido(w, id) {
		return
	}

	var bancos modelo.Bancos
	err := c.db.Colección().Find(bson.M{"provinciaId": bson.ObjectIdHex(id)}).Sort("ciudad").All(&bancos)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	encodables := bancos.PrepararParaEncode(c.db.GetMongoDB())

	if err := json.NewEncoder(w).Encode(encodables); err != nil {
		panic(err)
	}

}

func (c *Banco) BancoDistancia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lat, laterr := strconv.ParseFloat(vars["lat"], 32)
	lon, lonerr := strconv.ParseFloat(vars["lon"], 32)

	rango, rangoerr := strconv.ParseFloat(vars["rango"], 32)

	response := make(map[string]interface{})
	if laterr != nil || lonerr != nil || rangoerr != nil {
		response["error"] = "Coordenadas invalidas"
	} else {

		var bancos modelo.Bancos
		q := c.db.FindNear(lat, lon, rango).Limit(10)
		q.All(&bancos)
		encodables := bancos.PrepararParaEncode(c.db.GetMongoDB())
		response["bancos"] = encodables
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
