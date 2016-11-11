package task

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/modelo"
)

const day = time.Hour * 24

//const day = time.Second * 4

func EmpezarTaskLimpieza(db *db.Database) {
	ticker := time.NewTicker(day)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				limpiar(db)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func limpiar(db *db.Database) {
	ahora := time.Now()
	haceUnMes := ahora.AddDate(0, -1, 0)
	query := db.Find(bson.M{
		"fecha": bson.M{
			"$lt": haceUnMes,
		},
	})

	var solicitudesViejas modelo.Solicitudes

	query.All(&solicitudesViejas)

	for _, solicitud := range solicitudesViejas {
		solicitud.Grupo = *modelo.NewGrupoSanguineo(-1, "Eliminado")
		solicitud.Factor = *modelo.NewFactorSanguineo(-1, "Eliminado")
		db.Update(solicitud.GetId().Hex(), &solicitud)
	}
}
