package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handlers "krejzac.cz/krejzac-api/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("mnam krejzac miluju krezaceie joooo!")
	})
	app.Get("/pfp/:userName", handlers.GetPfp)

	log.Fatal(app.Listen(":42069"))
}
