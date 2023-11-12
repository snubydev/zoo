package main

import (
	"zoo/services"
	"zoo/webserver"
)

func main() {
	println("hello")

	zoo := services.NewZoo()

	webserver.NewWebServer(zoo)
	webserver.Run("3000")
}
