package main

import (
	"fmt"
)

func main() {
	fmt.Println("** Main **")

	app := App{}
	app.Initialize()
	app.PopulateRoutes()
	app.run(":8080")

	// log.Fatal(http.ListenAndServe(":8080", a.Initialize()))

	//router := NewRouter()
	//log.Fatal(http.ListenAndServe(":8080", router))
}
