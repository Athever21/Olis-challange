package services

import (
	"fmt"
	"olist/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Filters struct {
	Name             string `query:"name"`
	Publication_year string `query:"publication_year"`
	Edition          string `query:"edition"`
	Author           string `query:"author"`
}

type Body struct {
	Name             string `json:"name" xml:"name" form:"name"`
	Edition          string `json:"edition" xml:"edition" form:"edition"`
	Publication_year string `json:"publication_year" xml:"publication_year" form:"publication_year"`
	Author           string `json:"author" xml:"author" form:"author"`
}

func GetAllBooks(c *fiber.Ctx) error {
	filters := new(Filters)

	if err := c.QueryParser(filters); err != nil {
		return err
	}

	books, err := models.GetAllBooks(filters.Name, filters.Author, filters.Edition, filters.Publication_year)

	if err != nil {
		c.Status(400).SendString("Bad Request")
	}

	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(Body)

	if err := c.BodyParser(book); err != nil {
		return err
	}

	if book.Name == "" || book.Edition == "" || book.Publication_year == "" || book.Author == "" {
		return c.Status(400).SendString("Bad Request")
	}

	author, err := models.GetAllAuthors(0, 1, book.Author)

	if err != nil || len(author) != 1 {
		newAuthor, err := models.SaveAuthor(book.Author)

		if err != nil {
			return c.Status(500).SendString("Something went wrong")
		}
		err = models.SaveBook(book.Name, book.Edition, book.Publication_year, strconv.Itoa(newAuthor.ID))

		if err != nil {
			return c.Status(500).SendString("Something went wrong")
		}

		return c.JSON(&fiber.Map{"success": true})
	} else {
		err = models.SaveBook(book.Name, book.Edition, book.Publication_year, strconv.Itoa(author[0].ID))

		if err != nil {
			fmt.Println(err.Error())
			return c.Status(500).SendString("Something went wrong")
		}

		return c.JSON(&fiber.Map{"success": true})
	}

}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := models.GetBook(id)

	if err != nil {
		return c.Status(500).SendString("Something went wrong")
	}

	if book.ID == 0 {
		return c.Status(404).SendString("Book not found")
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	err := models.DeleteBook(id)

	if err != nil {
		return c.Status(500).SendString("Something went wrong")
	}

	return c.JSON(&fiber.Map{"success": true})
}

func ChangeBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := models.GetBook(id)

	if err != nil || book.ID == 0 {
		return c.Status(404).SendString("Book not found")
	}

	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	err = models.UpdateBook(id, body.Name, body.Edition, body.Publication_year)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).SendString("Something went wrong")
	}

	return c.JSON(&fiber.Map{"success": true})
}
