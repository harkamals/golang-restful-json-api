package latest

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

func (app *App) Run(addr string) {

	handler := handlers.CombinedLoggingHandler(os.Stdout, app.Router)
	go http.ListenAndServeTLS(addr, "/Users/hk/Documents/code/go/certs/cert.pem", "/Users/hk/Documents/code/go/certs/key.pem", handler)

	// Redirect to https
	http.ListenAndServe(":8081", http.HandlerFunc(redirectToHttps))

}
