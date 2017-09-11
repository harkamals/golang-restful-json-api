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
)

var a main.App

func TestMain(m *testing.M) {

	a = main.App{}
	a.Initialize()
	a.PopulateRoutes()

	code := m.Run()
	os.Exit(code)

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