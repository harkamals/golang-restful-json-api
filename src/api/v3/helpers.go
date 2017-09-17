package v3

import (
	"encoding/json"
	"net/http"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func respondWithJSON(w http.ResponseWriter, statusCode int, input interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(input); err != nil {
		panic(err)
	}

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, jsonErr{Code: code, Text: message})
}
