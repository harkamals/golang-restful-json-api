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

	// Define routes
	app.Routes = Routes{
		Route{
			"Index", "GET", "/", app.TOC,
		},
		Route{
			"TodoIndex", "GET", "/todos", todo_list,
		},
		Route{
			"TodoCreate", "POST", "/todos", TodoCreate,
		},
		Route{
			"TodoShow", "GET", "/todos/{todoId}", TodoShow,
		},
		Route{
			"Orders", "GET", "/orders", app.getOrders,
		},
		Route{
			"Order", "POST", "/order", app.createOrder,
		},
		Route{
			"Order", "GET", "/order/{id:[0-9]+}", app.getOrder,
		},
		Route{
			"Order", "PUT", "/order/{id:[0-9]+}", app.updateOrder,
		},
		Route{
			"Order", "DELETE", "/order/{id:[0-9]+}", app.deleteOrder,
		},
		//Route{
		//	"Logs", "GET", "/logs", http.FileServer(http.Dir("/")),
		//},
	}

	// 404
	app.Router.NotFoundHandler = http.HandlerFunc(not_found_404)

	// Logs
	app.Router.PathPrefix("/logs/").Handler(
		http.StripPrefix("/logs/", http.FileServer(http.Dir("."))))

	// Enumerate routes
	for _, route := range app.Routes {
		var handler http.Handler

		handler = route.HandlerFunc

		// handler = Logger(handler, route.Name)
		handler = Authenticator(handler, route.Name)

		app.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

}
