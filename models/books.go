package models

import (
	"fmt"
	"strings"
)

type Book struct {
	ID               int    `json:"ID"`
	Name             string `json:"name"`
	Edition          int    `json:"edition"`
	Publication_year int    `json:"publication_year"`
	Author           string "json:`author`"
}

func GetBook(id string) (*Book, error) {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return nil, err
	}

	result, err := db.Query(fmt.Sprintf("SELECT books.Id,books.name,books.publication_year,books.edition,authors.Name FROM books INNER JOIN authors ON books.author = authors.ID WHERE books.ID = %s", id))

	if err != nil {
		return nil, err
	}

	var book Book

	if result.Next() {
		err = result.Scan(&book.ID, &book.Name, &book.Publication_year, &book.Edition, &book.Author)

		if err != nil {
			return nil, err
		}
	}

	return &book, nil
}

func GetAllBooks(name, author, edition, publication_year string) ([]Book, error) {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return nil, err
	}

	results, err := db.Query(fmt.Sprintf("SELECT books.Id,books.name,books.publication_year,books.edition,authors.Name FROM books INNER JOIN authors ON books.author = authors.ID WHERE books.name LIKE '%%%s%%' AND authors.name LIKE '%%%s%%' AND edition LIKE '%%%s%%' AND publication_year LIKE '%%%s%%'", name, author, edition, publication_year))

	if err != nil {
		return nil, err
	}

	books := []Book{}

	for results.Next() {
		var book Book
		err := results.Scan(&book.ID, &book.Name, &book.Publication_year, &book.Edition, &book.Author)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func SaveBook(name, edition, publication_year, author string) error {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return err
	}

	_, err = db.Query(fmt.Sprintf("INSERT INTO books(name,edition,publication_year,author) VALUES ('%s',%s,%s,%s)", name, edition, publication_year, author))

	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(id string) error {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return err
	}

	_, err = db.Query(fmt.Sprintf("DELETE FROM books WHERE ID = %s", id))

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(id, name, edition, publication_year string) error {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return err
	}
	updateString := make([]string, 0, 3)

	if name != "" {
		updateString = append(updateString, fmt.Sprintf("name = '%s'", name))
	}
	if edition != "" {
		updateString = append(updateString, fmt.Sprintf("edition = %s", edition))
	}
	if publication_year != "" {
		updateString = append(updateString, fmt.Sprintf("publication_year = %s", publication_year))
	}

	_, err = db.Query(fmt.Sprintf("UPDATE books SET %s WHERE ID = %s", strings.Join(updateString, (", ")), id))

	if err != nil {
		return err
	}

	return nil
}
