package latest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"fmt"
	"github.com/gorilla/mux"
)

func not_found_404(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotFound, "not found")
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// todo: read from configuration
	http.Redirect(w, r, "https://127.0.0.1:8443"+r.RequestURI, http.StatusMovedPermanently)
}

func (app *App) TOC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome\n")
	for _, route := range app.Routes {
		fmt.Fprint(w, route.Method, "  ", route.Pattern, "\n")
	}
}

// ** POSTS **
func (app *App) getPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := getPosts(app.Gorm)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)

}

func (app *App) getPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	p := Post{Id: id}
	p.getPost(app.Gorm)

	respondWithJSON(w, http.StatusOK, p)

}

func (app *App) createPost(w http.ResponseWriter, r *http.Request) {

	var p Post
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	p.createPost(app.Gorm)
	respondWithJSON(w, http.StatusCreated, p)

}

func (app *App) updatePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid order id")
		return
	}

	var p Post
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	defer r.Body.Close()
	p.Id = id

	p.updatePost(app.Gorm)
	respondWithJSON(w, http.StatusOK, p)
}

func (app *App) deletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	p := Post{Id: id}
	p.deletePost(app.Gorm)

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})

}

// ** TODOS app **
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
