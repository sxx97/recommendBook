package main

import (
	"github.com/kataras/iris/v12"
	"readBook/internal/routes"
)
func main() {
	app := iris.New()
	routes.Route(app)
	routes.RouteError(app)
	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("server terminated")
		})
	})
	configFile := iris.WithConfiguration(iris.YAML("./config/main.yaml"))
	app.Run(iris.Addr(":9090"), configFile)
}

