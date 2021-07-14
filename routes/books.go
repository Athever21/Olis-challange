package routes

import (
	"olist/services"

	"github.com/gofiber/fiber/v2"
)

func BooksRouter(app *fiber.App) {
	app.Get("/api/books", services.GetAllBooks)
	app.Post("/api/books", services.CreateBook)
	app.Get("/api/books/:id", services.GetBook)
	app.Delete("/api/books/:id", services.DeleteBook)
	app.Put("/api/books/:id", services.ChangeBook)
}
