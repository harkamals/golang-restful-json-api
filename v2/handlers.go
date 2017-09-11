package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func json_encoder(w http.ResponseWriter, statusCode int, input interface{}) {

	fmt.Println("Encoding..")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err:= encoder.Encode(input); err!=nil{
		panic(err)
	}



}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	json_encoder(w, http.StatusOK, todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {

		json_encoder(w, http.StatusOK, todo)
		return

	}

	// If we didn't find it, 404
	json_encoder(w, http.StatusNotFound, jsonErr{Code: http.StatusNotFound, Text: "Not Found"})

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		json_encoder(w, 422, err)
	}

	t := RepoCreateTodo(todo)
	json_encoder(w, http.StatusCreated, t)

}
