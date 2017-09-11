package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
	_"net/http"
	_"log"
)

func main() {
	fmt.Println("** Main **")

	app := App{}
	app.Initialize()
	app.run(":8080")

	// log.Fatal(http.ListenAndServe(":8080", a.Initialize()))

	//router := NewRouter()
	//log.Fatal(http.ListenAndServe(":8080", router))
}
