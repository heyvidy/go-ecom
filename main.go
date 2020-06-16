package main

import (
	"database/sql"

	models "./models"
)

var db *sql.DB

func main() {
	db = models.InitializeDB(db)

	admin := models.User{

		Password: "amazon_vd_123",
		Email:    "vd3@speckbit.com",
		Phno:     "123"}

	// models.CreateUser(db, &admin)
	// models.DeleteUser(db, "vd")
	models.UpdateUser(db, "vd", &admin)
	models.GetAllUsers(db)
}
