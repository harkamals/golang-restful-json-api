package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (app *App) initializeRoutes() {

	app.Routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			app.TOC,
		},
		Route{
			"TodoIndex",
			"GET",
			"/todos",
			todo_list,
		},
		Route{
			"TodoCreate",
			"POST",
			"/todos",
			TodoCreate,
		},
		Route{
			"TodoShow",
			"GET",
			"/todos/{todoId}",
			TodoShow,
		},
		Route{
			"Order",
			"GET",
			"/order/{id:[0-9]+}",
			app.getOrder,
		},
	}

	for _, route := range app.Routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		app.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

}
