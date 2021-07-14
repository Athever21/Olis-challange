package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getDb() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/olist")
}

func closeDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
