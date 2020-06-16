package main

import (
	"database/sql"

	models "./models"
)

var db *sql.DB

func main() {
	db = models.InitializeDB(db)

	admin := models.User{
		Email: "spoofball3@speckbit.com"}

	models.CreateUser(db, &admin)
	// models.DeleteUser(db, "vd")
	// models.UpdateUser(db, "vd", &admin)
	models.GetAllUsers(db)
}
