package main

import (
	"log"
	"olist/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	routes.AuthorsRouter(app)
	routes.BooksRouter(app)

	log.Fatal(app.Listen(":3000"))
}
