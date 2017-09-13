package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"fmt"
	"database/sql"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Routes []Route
}

func (app *App) Initialize(dbUser, dbPass, db string) *mux.Router {

	fmt.Println("App is initializing..")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, db)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.initializeRoutes()

	return app.Router
}

func (app *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
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
	json_encoder(w, http.StatusOK, o)

}
