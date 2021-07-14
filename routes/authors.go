package routes

import (
	"olist/services"

	"github.com/gofiber/fiber/v2"
)

func AuthorsRouter(app *fiber.App) {
	app.Get("/api/authors", services.GetAllAuthors)
}
