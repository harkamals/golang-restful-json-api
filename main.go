package main

import (
	"api/latest"
)

func main() {

	app := latest.App{}
	app.InitDb(
		"localhost",
		"5409",
		"postgres",
		"postgres",
		"postgres")

	app.InitRoutes()
	app.Run(":8443")

}
