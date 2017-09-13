// ************
// todo : https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// ************

package main_test

import (
	"testing"
	"."
	"os"
	"net/http"
	"net/http/httptest"
	"log"
	"github.com/spf13/viper"
)

var a main.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS orders
(
id SERIAL,
NAME TEXT NOT NULL,
price NUMERIC (10, 2) NOT NULL DEFAULT 0.00,
CONSTRAINT orders_pkey PRIMARY KEY (id)
)`

func TestMain(m *testing.M) {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	a = main.App{}
	a.Initialize(
		viper.GetString("testing.dbUser"),
		viper.GetString("testing.dbPass"),
		viper.GetString("testing.db"))

	a.PopulateRoutes()

	ensureTableExists()
	code := m.Run()
	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM orders")
	a.DB.Exec("ALTER SEQUENCE orders_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/orders", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}








func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetOrders(t *testing.T) {

	req, _ := http.NewRequest("GET", "/orders", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetTodos(t *testing.T) {

	req, _ := http.NewRequest("GET", "/todos", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

