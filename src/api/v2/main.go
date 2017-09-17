package v2

func main() {

	app := App{}
	app.Initialize(
		"postgres",
		"postgres",
		"postgres")

	app.run(":8080")

}
