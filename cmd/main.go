package main

import "hermes/cmd/server"

func main() {

	app := server.NewServer(nil)
	app.Run()
}
