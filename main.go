package main

import "api/v3"

func main() {

	app := v3.App{}
	app.Initialize(
		"postgres",
		"postgres",
		"postgres")

	app.Run(":8080")

}
