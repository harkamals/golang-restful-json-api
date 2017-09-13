package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func json_encoder(w http.ResponseWriter, statusCode int, input interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(input); err != nil {
		panic(err)
	}

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	json_encoder(w, code, jsonErr{Code: code, Text: message})
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
