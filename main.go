package main

import (
	"zoo/services"
	"zoo/webserver"
)

func main() {
	zoo := services.NewZoo()
	webserver.NewWebServer(zoo)
	webserver.Run("3000")
}
