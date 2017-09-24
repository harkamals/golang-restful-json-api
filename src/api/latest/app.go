package latest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Db     *gorm.DB
	Routes []Route
	Config config
}

func (app *App) InitDB() {

	fmt.Println("Init Db..")

	// Read config.json
	data, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(data, &app.Config)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		app.Config.Db.Host, app.Config.Db.Port, app.Config.Db.User, app.Config.Db.Pass, app.Config.Db.Name)

	var err error
	app.Db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	app.Db.LogMode(true)
	app.Db.AutoMigrate(
		&Post{},
		&Comment{},
		&Orders{},
		&Accounts{})

}

func (app *App) Run() {

	fmt.Println("Running..")

	defer app.Db.Close()

	handler := handlers.CombinedLoggingHandler(os.Stdout, app.Router)
	go http.ListenAndServeTLS(app.Config.Https, app.Config.Cert, app.Config.Key, handler)

	// Redirect to https
	http.ListenAndServe(app.Config.Http, http.HandlerFunc(app.redirectToHttps))

}
