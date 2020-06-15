package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phno     string `json:"phno"`
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CreateUser(db *sql.DB, data *User) {
	sqlQuery := "INSERT user SET username = ?, password = ?, name = ?, phno = ?, email = ?, timestamp = ?"

	stmt, err := db.Prepare(sqlQuery)

	printError(err)

	_, err = stmt.Exec(data.Username, data.Password, data.Name, data.Phno, data.Email, time.Now())

	printError(err)

}

func GetAllUsers(db *sql.DB) {
	users, err := db.Query("SELECT username, password, email, phno FROM user")
	printError(err)

	for users.Next() {
		var user User

		err := users.Scan(&user.Username, &user.Password, &user.Email, &user.Phno)
		printError(err)

		fmt.Println(user)
	}
}

func DeleteUser(db *sql.DB, username string) {
	query := "DELETE FROM user WHERE username = ?"
	stmt, _ := db.Prepare(query)
	_, err := stmt.Exec(username)

	if err == nil {
		fmt.Println("Deleted User with username", username)

	}
}

func UpdateUser(db *sql.DB, username string, data *User) {
	userjson, err1 := json.Marshal(data)
	printError(err1)

	var usermap map[string]string

	json.Unmarshal([]byte(userjson), &usermap)

	query := "UPDATE user SET "
	count := 1
	for field, value := range usermap {
		if value != "" {
			if count < len(usermap) {

				query += fmt.Sprintf("%s='%s', ", field, value)
			} else {
				query += fmt.Sprintf("%s='%s' ", field, value)

			}
		}
		count += 1
	}

	query += fmt.Sprintf("WHERE username='%s'", username)
	fmt.Println(query)
	_, err2 := db.Query(query)
	printError(err2)
}
