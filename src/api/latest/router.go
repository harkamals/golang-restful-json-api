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

		// ** ORDERS
		Route{
			"order-new",
			"POST",
			"/order",
			app.create_order},

		Route{
			"order-get",
			"GET",
			"/order/{id:[0-9]+}",
			app.get_order},

		Route{
			"order-get-all",
			"GET",
			"/orders",
			app.get_orders},

		Route{
			"order-status-update",
			"PUT",
			"/order",
			app.update_order},

		Route{
			"order-delete",
			"DELETE",
			"/order",
			app.delete_order},

		// ** ACCOUNTS
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
			"account-get-all-emails",
			"GET",
			"/emails",
			app.GetEmails},

		Route{
			"email-get",
			"GET",
			"/email/{id:[0-9]+}",
			app.GetEmail},

		Route{
			"account-get-next-available-email-for-account-creation",
			"GET",
			"/email/next",
			app.GetNextEmail},

		// ** POSTS
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

		// TO-DO
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
