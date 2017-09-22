// ************
// Inspired by : https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// ************

package main

import (
	"api/latest"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var app latest.App

func TestMain(m *testing.M) {

	app = latest.App{}
	app.Initialize(
		"localhost",
		"5409",
		"postgres",
		"postgres",
		"postgres_test")

	ensureTableExists()
	code := m.Run()
	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	app.Gorm.Create(app.Post)
}

func clearTable() {
	app.Gorm.Delete(app.Post)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// Model: TODOs
func Test_GetTodos(t *testing.T) {
	req, _ := http.NewRequest("GET", "/todos", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

// Model: POST
func Test_EmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/posts", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); strings.TrimSpace(body) != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func Test_GetNonExistentProduct(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/post/999", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["text"] != "order not found" {
		t.Errorf("Expected the 'text' key of the response to be set to 'order not found'. Got '%s'", m["text"])
	}
}

func Test_GetPosts(t *testing.T) {
	req, _ := http.NewRequest("GET", "/posts", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func Test_GetPost(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/post/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		// app.DB.Exec("INSERT INTO orders(name, price) VALUES($1, $2)", "Order "+strconv.Itoa(i), (i+1.0)*10)
	}
}

func Test_Createorder(t *testing.T) {
	clearTable()
	payload := []byte(`{"name": "test order", "price": 11.22 }`)

	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test order" {
		t.Errorf("expected name to be 'test order'. Got %v", m["name"])
	}

	if m["price"] != 11.22 {
		t.Errorf("expected price to be '11.22'. Got %v", m["price"])
	}

	// m[string]interface{} converts int to float
	if m["id"] != 1.0 {
		t.Errorf("expected id to be '1.0'. Got %v", m["id"])
	}

}

func Test_UpdateOrder(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/order/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var originalOrder map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalOrder)

	payload := []byte(`{"name": "updated order", "price": 11.22 }`)

	req, _ = http.NewRequest("PUT", "/order/1", bytes.NewBuffer(payload))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalOrder["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalOrder["id"], m["id"])
	}

	if m["name"] == originalOrder["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalOrder["name"], m["name"], m["name"])
	}

	if m["price"] == originalOrder["price"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", originalOrder["price"], m["price"], m["price"])
	}
}

func Test_DeletePost(t *testing.T) {

	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/post/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/post/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/post/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

}
