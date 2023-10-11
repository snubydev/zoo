package main

import "zoo/webserver"

func main() {
	println("hello")
	webserver.NewWebServer()
	webserver.Run("3000")
}
