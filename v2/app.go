package main

import (
	"net/http"

	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
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
