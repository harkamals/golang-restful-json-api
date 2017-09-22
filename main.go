package main

import (
	"api/latest"
)

func main() {

	app := latest.App{}
	app.Initialize(
		"localhost",
		"5409",
		"postgres",
		"postgres",
		"postgres")

	app.Run(":8443")

}
