package middleware

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwtm "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func Headers(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	next(w, r)
}

func Logger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	comienzo := time.Now()

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(comienzo),
	)

	next(w, r)
}

func Jwt() *jwtm.JWTMiddleware {
	return jwtm.New(jwtm.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			decoded, err := base64.URLEncoding.DecodeString(os.Getenv("AUTH0_CLIENT_SECRET"))
			if err != nil {
				return nil, err
			}
			return decoded, nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, authErr string) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(401)

			err := map[string]interface{}{"error": authErr}
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
