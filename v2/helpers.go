package main

import (
	"net/http"
	"encoding/json"
)

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

