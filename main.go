package main

import (
	route "devops-config-agent/src/route/api"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	corsInit := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "ACCEPT", "Access-Control-Allow-Origin"},
	})
	app.Use(corsInit)

	app = route.InitRoute(app)
	app.Run(iris.Addr(":1919"))
}
