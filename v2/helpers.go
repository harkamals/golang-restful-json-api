package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func json_encoder(w http.ResponseWriter, statusCode int, input interface{}) {

	fmt.Println("JSON:encoding..")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err:= encoder.Encode(input); err!=nil{
		panic(err)
	}

}