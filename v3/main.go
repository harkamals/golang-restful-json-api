package main

func main() {

	app := App{}
	app.Initialize(
		"postgres",
		"postgres",
		"postgres")

	app.run(":8080")

}
