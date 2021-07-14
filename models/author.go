package models

import (
	"fmt"
)

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllAuthors(page, limit int, name string) ([]Author, error) {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return nil, err
	}

	results, err := db.Query(fmt.Sprintf("SELECT * FROM authors WHERE Name LIKE '%%%s%%' LIMIT %d OFFSET %d", name, limit, page*limit))
	if err != nil {
		return nil, err
	}

	authors := []Author{}

	for results.Next() {
		var author Author
		err = results.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func SaveAuthor(name string) (*Author, error) {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return nil, err
	}

	_, err = db.Query(fmt.Sprintf("INSERT INTO authors(name) VALUES (\"%s\")", name))

	if err != nil {
		return nil, err
	}

	author, err := GetAllAuthors(0, 1, name)

	if err != nil {
		return nil, err
	}

	return &author[0], nil
}

func SaveAuthors(values string) error {
	db, err := getDb()
	defer closeDb(db)

	if err != nil {
		return err
	}

	_, err = db.Query(fmt.Sprintf("INSERT INTO authors(name) VALUES %s", values))

	if err != nil {
		return err
	}

	return nil
}
