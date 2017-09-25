package latest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (app *App) InitRoutes() *mux.Router {

	fmt.Println("Init Router..")

	// Router
	app.Router = mux.NewRouter().StrictSlash(true)

	// Define routes
	app.Routes = Routes{

		Route{
			"table-of-contents",
			"GET",
			"/",
			app.TOC},

		Route{
			"account-new",
			"POST",
			"/account",
			app.create_account},

		Route{
			"account-get",
			"GET",
			"/account/{id:[0-9]+}",
			app.get_account},

		Route{
			"account-get-all",
			"GET",
			"/accounts",
			app.get_accounts},

		Route{
			"post-new",
			"POST",
			"/post",
			app.createPost},

		Route{
			"post-get",
			"GET",
			"/post/{id:[0-9]+}",
			app.getPost},

		Route{
			"post-get-all",
			"GET",
			"/posts",
			app.getPosts},

		Route{
			"post-update",
			"PUT",
			"/post/{id:[0-9]+}",
			app.updatePost},

		Route{
			"post-delete",
			"DELETE",
			"/post/{id:[0-9]+}",
			app.deletePost},

		Route{"todo-new", "POST", "/todos", TodoCreate},
		Route{"todo-get-all", "GET", "/todos", todo_list},
		Route{"todo-get", "GET", "/todos/{todoId}", TodoShow},
	}

	// 404
	app.Router.NotFoundHandler = http.HandlerFunc(not_found_404)

	// Logs
	os.Mkdir("./logs", os.FileMode(755))
	app.Router.PathPrefix("/logs/").Handler(
		http.StripPrefix("/logs/", http.FileServer(http.Dir("./logs"))))

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

	return app.Router

}
