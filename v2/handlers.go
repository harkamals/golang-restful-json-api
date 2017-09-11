package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called")
	fmt.Fprint(w, "Welcome\n")

	for _, route:= range routing_list{
		fmt.Fprint(w, route, "\n")
	}

}

func todo_list(w http.ResponseWriter, r *http.Request) {
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
