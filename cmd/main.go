package main

import "hermes/cmd/server"

func main() {

	app := server.NewServer(nil)
	app.Run()
	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// app.Listen(":3000")
}
