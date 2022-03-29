package middlewares

import (
	"itsln.com/anonym/errors"
	"log"
	"net/http"
)

func errorHandling(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer recovery(w)
		next.ServeHTTP(w, r)
	})
}

func recovery(w http.ResponseWriter) {
	if e := recover(); e != nil {
		switch err := e.(type) {
		case errors.ServerError:
			http.Error(w, err.Error(), err.Status)
		case error:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("(errorHandling) Unhandled exception: %s", err.Error())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("(errorHandling) Waiting error but give: %s", err)
		}
	}
}
