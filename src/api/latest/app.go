package latest

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Db     *gorm.DB
	Routes []Route
}

func (app *App) InitDB(dbHost, dbPort, dbUser, dbPass, db string) {

	fmt.Println("Init Db..")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, db)

	var err error
	app.Db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	app.Db.LogMode(true)
	app.Db.AutoMigrate(&Post{}, &Comment{})

}

func (app *App) Run(addr string) {

	fmt.Println("Running..")

	defer app.Db.Close()

	handler := handlers.CombinedLoggingHandler(os.Stdout, app.Router)
	go http.ListenAndServeTLS(addr, "/Users/hk/Documents/code/go/certs/cert.pem", "/Users/hk/Documents/code/go/certs/key.pem", handler)

	// Redirect to https
	http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))

}
