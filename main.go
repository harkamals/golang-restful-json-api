package main

import (
	"api/latest"
	"encoding/json"
	"os"
)

type config struct {
	http_port  string
	https_port string
	cert       string
	key        string
}

func main() {

	var c config

	//data, err := ioutil.ReadFile("config.json")
	//if err != nil {
	//	println(err.Error())
	//}
	//
	//print(string(data))

	configFile, _ := os.Open("config.json")
	jsonparser := json.NewDecoder(configFile)
	jsonparser.Decode(&c)

	println(jsonparser)

	println(c.http_port)
	println(c.https_port)
	println(c.cert)
	println(c.key)

	print("----")
	return

	app := latest.App{}
	app.InitDB(
		"localhost",
		"5409",
		"postgres",
		"postgres",
		"postgres")

	app.InitRoutes()
	app.Run(":8443")

}
