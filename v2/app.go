package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"log"
	"fmt"
	"database/sql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(dbUser, dbPass, db string) *mux.Router {

	fmt.Println("App initializing..")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, db)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		app.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return app.Router
}

func (app *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

// List of /paths
var routing_list []string

func (app *App) PopulateRoutes() {
	for _, r := range routes {
		routing_list = append(routing_list, r.Pattern)
	}
}

//func NewRouter() *mux.Router {
//
//	router := mux.NewRouter().StrictSlash(true)
//	for _, route := range routes {
//		var handler http.Handler
//
//		handler = route.HandlerFunc
//		handler = Logger(handler, route.Name)
//
//		router.
//		Methods(route.Method).
//			Path(route.Pattern).
//			Name(route.Name).
//			Handler(handler)
//
//	}
//
//	return router
//}
