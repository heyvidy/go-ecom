package models

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	ID        int    `json:"idcustomer"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phno      string `json:"phno"`
	Timestamp string `json:"timestamp"`
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CreateUser(db *sql.DB, data *User) {
	sqlQuery := "INSERT user SET username = ?, password = ?, name = ?, phno = ?, email = ?, timestamp = ?"

	stmt, err := db.Prepare(sqlQuery)
	err = data.HashPassword()

	_, err = stmt.Exec(data.Username, data.Password, data.Name, data.Phno, data.Email, time.Now())
	stmt.Close()
	printError(err)

}

func GetAllUsers(db *sql.DB) {
	users, err := db.Query("SELECT * FROM user")
	printError(err)

	for users.Next() {
		var user User

		err := users.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.Phno, &user.Timestamp)
		printError(err)

		fmt.Println(user)
	}
}

func DeleteUser(db *sql.DB, username string) {
	query := "DELETE FROM user WHERE username = ?"
	stmt, _ := db.Prepare(query)
	_, err := stmt.Exec(username)
	stmt.Close()

	if err == nil {
		fmt.Println("Deleted User with username", username)

	}
}

func UpdateUser(db *sql.DB, username string, data *User) {

	data_json, _ := json.Marshal(data)

	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	var user map[string]interface{}

	json.Unmarshal([]byte(data_json), &user)

	query := "UPDATE user SET "
	count := 1

	for field, val := range user {
		if val == "" {
			delete(user, field)
		}
	}

	for field, value := range user {

		if field != "" {
			if count < len(user) {
				query += fmt.Sprint(strings.ToLower(field), "=", "'", value, `',`)
			} else {
				query += fmt.Sprint(strings.ToLower(field), "=", "'", value, `' `)

			}
		}
		count += 1

		fmt.Println(count)
	}

	query += fmt.Sprintf("WHERE username='%s'", username)

	fmt.Println(query)
	_, err2 := db.Query(query)
	printError(err2)
}

// HashPassword : hashing the password using PBKDF2_SHA256
func (user *User) HashPassword() error {
	randByte := make([]byte, 8)

	_, err := rand.Read(randByte)
	if err != nil {
		return err
	}

	base64RandByte := base64.StdEncoding.EncodeToString(randByte)
	salt := []byte(base64RandByte)

	iter := 100000

	dk := pbkdf2.Key([]byte(user.Password), salt, iter, 32, sha256.New)

	hashedPW := "pbkdf2_sha256$100000$" + string(salt) + "$" + base64.StdEncoding.EncodeToString(dk)

	user.Password = hashedPW

	return nil
}

// ComparePassword : compare the password
func (user *User) ComparePassword(password string) bool {
	splitted := strings.Split(user.Password, "$")

	salt := []byte(splitted[2])

	// saved password iteration value should be converted to int
	iter, _ := strconv.Atoi(splitted[1])

	dk := pbkdf2.Key([]byte(password), salt, iter, 32, sha256.New)

	hashedPW := "pbkdf2_sha256$100000$" + splitted[2] + "$" + base64.StdEncoding.EncodeToString(dk)

	if subtle.ConstantTimeCompare([]byte(user.Password), []byte(hashedPW)) == 0 {
		return false
	}

	return true
}
