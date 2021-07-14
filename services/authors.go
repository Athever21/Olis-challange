package services

import (
	"olist/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Filter struct {
	Limit string `query:"limit"`
	Page  string `query:"page"`
	Name  string `query:"name"`
}

func GetAllAuthors(c *fiber.Ctx) error {
	filters := new(Filter)

	if err := c.QueryParser(filters); err != nil {
		return err
	}

	if filters.Page == "" {
		filters.Page = "0"
	}

	limit, err := strconv.Atoi(filters.Limit)

	if err != nil || limit > 100 {
		limit = 50
	}

	page, err := strconv.Atoi(filters.Page)

	if err != nil {
		page = 0
	}

	authors, err := models.GetAllAuthors(page, limit, filters.Name)

	if err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	return c.JSON(authors)
}
