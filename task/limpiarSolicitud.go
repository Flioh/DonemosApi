package task

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/flioh/DonemosApi/db"
	"github.com/flioh/DonemosApi/modelo"
)

const day = time.Second * 4

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
	haceUnMes := ahora.AddDate(0, 0, -1)
	query := db.Find(bson.M{
		"fechaCreacion": bson.M{
			"$lt": haceUnMes,
		},
	})

	var solicitudesViejas modelo.Solicitudes

	query.All(&solicitudesViejas)
	fmt.Printf("\nsolicViejas:\n%+v\n\n", solicitudesViejas)
}
