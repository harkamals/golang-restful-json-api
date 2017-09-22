package latest

import (
	"encoding/json"
	"net/http"
)

// Settings: Config.json
type config struct {
	Host    string   `json:"host"`
	Http    string   `json:"http_port"`
	Https   string   `json:"https_port"`
	Cert    string   `json:"cert"`
	Key     string   `json:"key"`
	Db      dbconfig `json:"default"`
	Db_test dbconfig `json:"test"`
}

type dbconfig struct {
	Host string `json:"dbhost"`
	Port int    `json:"dbport"`
	User string `json:"dbuser"`
	Pass string `json:"dbpass"`
	Name string `json:"dbname"`
}

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
