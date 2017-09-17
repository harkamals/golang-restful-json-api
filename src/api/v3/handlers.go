package v3

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
)

func not_found_404(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotFound, "not found")

}

func (app *App) TOC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome\n")
	for _, route := range app.Routes {
		fmt.Fprint(w, route.Pattern, "\n")
	}
}

func (app *App) getOrders(w http.ResponseWriter, r *http.Request) {

	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}

	if start < 0 {
		start = 0
	}

	orders, err := getOrders(app.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, orders)

}

func (app *App) getOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid order id")
		return
	}

	o := Order{Id: id}

	if err := o.getOrder(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "order not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, o)

}

func (app *App) createOrder(w http.ResponseWriter, r *http.Request) {

	var o Order
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := o.createOrder(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, o)

}

func (app *App) updateOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid order id")
		return
	}

	var o Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	defer r.Body.Close()
	o.Id = id

	if err := o.updateOrder(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, o)

}

func (app *App) deleteOrder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	o := Order{Id: id}
	if err := o.deleteOrder(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}

// TODOS app
func todo_list(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		// error
		respondWithJSON(w, http.StatusInternalServerError, jsonErr{Code: http.StatusInternalServerError, Text: "??"})
		return
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {

		respondWithJSON(w, http.StatusOK, todo)
		return

	}

	// If we didn't find it, 404
	respondWithJSON(w, http.StatusNotFound, jsonErr{Code: http.StatusNotFound, Text: "Not Found"})

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
		respondWithJSON(w, 422, err)
	}

	t := RepoCreateTodo(todo)
	respondWithJSON(w, http.StatusCreated, t)

}
