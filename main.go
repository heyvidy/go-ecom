package main

import (
	"database/sql"

	models "./models"
)

var db *sql.DB

func main() {
	db = models.InitializeDB(db)

	admin := models.User{
		Name:     "Chandra Vivek",
		Password: "xyzd",
		Email:    "vivs@speckbit.com",
		Phno:     "1234567890"}

	// admin.CreateUser(db)

	// models.DeleteUser(db, "vd")
	models.GetAllUsers(db)
	models.UpdateUser(db, "vivek", &admin)
	models.GetAllUsers(db)
}
