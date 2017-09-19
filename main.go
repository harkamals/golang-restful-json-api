package main

import (
	"api/latest"
)

func main() {

	app := latest.App{}
	app.Initialize(
		"postgres",
		"postgres",
		"postgres")

	app.Run(":8080")

}
