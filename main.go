package main

import (
	"project2/config"
	"project2/middlewares"
	"project2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	// start on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
