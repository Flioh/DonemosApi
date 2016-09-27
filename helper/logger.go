package helper

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, nombre string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comienzo := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			nombre,
			time.Since(comienzo),
		)
	})
}
