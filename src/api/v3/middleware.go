package v3

import (
	"log"
	"net/http"
	"time"
)

//if !isValidAPIKey(r.URL.Query().Get("key")) {
//respondErr(w, r, http.StatusUnauthorized, "invalid API key")
//return
//}

func Authenticator(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// println("Authenticator calling")
			if 1 == 1 {
				inner.ServeHTTP(w, r)
			} else {
				respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			}
		})
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %-s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
