package main

func main() {

	app := App{}
	app.Initialize()
	app.PopulateRoutes()
	app.run(":8080")

}
