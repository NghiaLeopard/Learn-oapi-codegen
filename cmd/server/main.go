package main

import (
	"log"
	"test-oapi/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := api.NewApi()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":8080"))
}
