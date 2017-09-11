package main

import (
	"net/http"
	"fmt"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index",
		"GET", "/", Index,},

	Route{"TodoIndex",
		"GET", "/todos", todo_list,},

	Route{"TodoCreate",
		"POST", "/todos", TodoCreate,},

	Route{"TodoShow",
		"GET", "/todos/{todoId}", TodoShow,},

	Route{"Orders",
		"GET", "/orders", orders_list,},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called")
	fmt.Fprint(w, "Welcome !!")

}
