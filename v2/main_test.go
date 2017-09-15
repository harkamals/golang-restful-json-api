// ************
// Inspired by : https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// ************

package main_test

import (
	"."
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
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

	if body := response.Body.String(); strings.TrimSpace(body) != "[]" {
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

func TestGetNonExistentProduct(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/order/999", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["text"] != "order not found" {
		t.Errorf("Expected the 'text' key of the response to be set to 'order not found'. Got '%s'", m["text"])
	}
}

func TestGetOrder(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/order/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO orders(name, price) VALUES($1, $2)", "Order "+strconv.Itoa(i), (i+1.0)*10)
	}
}
