package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeDB(db *sql.DB) *sql.DB {
	dbConnection, err := sql.Open("mysql", "root:password@(localhost:3306)/ecom")

	if err != nil {
		fmt.Println(err)
	}

	db = dbConnection

	return db

}
