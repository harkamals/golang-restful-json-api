package main

import (
	"api/latest"
)

func main() {

	app := latest.App{}
	app.InitDB()
	app.InitRoutes()
	app.Run()

}
