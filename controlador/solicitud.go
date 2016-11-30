package controlador

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/Flioh/DonemosApi/db"
	"github.com/Flioh/DonemosApi/helper"
	"github.com/Flioh/DonemosApi/modelo"
	"github.com/gorilla/mux"
)

// Solicitud es el controlador de solicitudes, contiene referencia a la sesion de mongodb
// y contiene los handlers relacionados a las solicitudes.
type Solicitud struct {
	db *db.Database
}

func NewSolicitud(db *db.Database) *Solicitud {
	return &Solicitud{db}
}

func (c *Solicitud) SolicitudIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	provinciaId := vars["provinciaId"]
	localidadId := vars["localidadId"]
	grupoId := vars["grupoId"]
	factorId := vars["factorId"]
	página, errorPágina := strconv.Atoi(vars["pag"])

	var solicitudes modelo.Solicitudes = make(modelo.Solicitudes, 20)

	var query bson.M = make(bson.M)
	if provinciaId != "" && provinciaId != "null" {
		if !helper.IdValido(w, provinciaId) {
			return
		}
		query["provinciaId"] = bson.ObjectIdHex(provinciaId)
	}
	if localidadId != "" && localidadId != "null" {
		if !helper.IdValido(w, localidadId) {
			return
		}
		query["localidadId"] = bson.ObjectIdHex(localidadId)
	}
	if (grupoId != "" && grupoId != "null") && (factorId != "" && factorId != "null") {
		query["tiposSanguineos.grupoSanguineo"], _ = strconv.Atoi(grupoId)
		query["tiposSanguineos.factorSanguineo"], _ = strconv.Atoi(factorId)
	}

	q := c.db.Find(query)

	if errorPágina == nil {
		q.Paginar(página)
	}
	q.All(&solicitudes)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	encodables := solicitudes.PrepararParaEncode(c.db.GetMongoDB())

	if err := json.NewEncoder(w).Encode(encodables); err != nil {
		panic(err)
	}

}

func (c *Solicitud) SolicitudUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["usuarioId"]

	var query = bson.M{"usuarioID": id}
	var solicitudes modelo.Solicitudes = make(modelo.Solicitudes, 5)

	q := c.db.Find(query)
	q.All(&solicitudes)

	encodables := solicitudes.PrepararParaEncode(c.db.GetMongoDB())

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(encodables); err != nil {
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(solicitud); err != nil {
		panic(err)
	}
}

func (c *Solicitud) SolicitudCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE METHOD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	s := new(modelo.Solicitud)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	if err := json.Unmarshal(body, s); err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("error: ", err)
		w.WriteHeader(422)
		return
	}

	if err := c.db.Create(s); err != nil {
		fmt.Println("error: ", err)
	}

	//sj, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	s.PrepararParaEncode(c.db.GetMongoDB())
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

func (c *Solicitud) SolicitudUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	nuevaSolicitud.PrepararParaEncode(c.db.GetMongoDB())
	if err := json.NewEncoder(w).Encode(nuevaSolicitud); err != nil {
		panic(err)
	}
}

func (c *Solicitud) SolicitudDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["solicitudId"]
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := c.db.Delete(id)
	if err != nil {
		w.WriteHeader(204)
		return
	}

	w.WriteHeader(202)
}

func (c *Solicitud) SolicitudPreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(http.StatusOK)
}
