package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!, This is Go-Fiber Tutorial Series")
	})

	app.Listen(":8080")

}
