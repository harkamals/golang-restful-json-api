package latest

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Gorm   *gorm.DB
	Routes []Route
	Post   Post
}

func (app *App) InitDb(dbHost, dbPort, dbUser, dbPass, db string) {

	fmt.Println("Init Db..")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, db)

	var err error
	app.Gorm, err = gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	app.Gorm.LogMode(true)
	app.Gorm.AutoMigrate(&Post{}, &Comment{})

}

func (app *App) Run(addr string) {

	fmt.Println("Running..")

	defer app.Gorm.Close()

	handler := handlers.CombinedLoggingHandler(os.Stdout, app.Router)
	go http.ListenAndServeTLS(addr, "/Users/hk/Documents/code/go/certs/cert.pem", "/Users/hk/Documents/code/go/certs/key.pem", handler)

	// Redirect to https
	http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))

}
