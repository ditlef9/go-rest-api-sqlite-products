// models/user.go
package models

import (
	"ekeberg.com/go-api-sql-gcp-products/db"
	"ekeberg.com/go-api-sql-gcp-products/utils"
	"errors"
)

type User struct {
	ID                 int64  `json:"id"`
	Email              string `json:"email" binding:"required"`
	Password           string `json:"password" binding:"required"`
	RegisteredDatetime string
	HumanOrService     string
}

func (u *User) SignUpUser() error {
	query := "INSERT INTO users(email, password, approved, human_or_service) VALUES (?, ?, 0, 'human')"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Hash the password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Execute the insert statement
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	// Get the last inserted ID and set it to u.ID
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId

	return nil
}

// modes/users.go::LoginUser()
func (u *User) LoginUser() error {
	query := "SELECT id, password, approved, human_or_service FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	var approved int

	// Scan both the password and the approval status from the row
	err := row.Scan(&u.ID, &retrievedPassword, &approved, &u.HumanOrService)
	if err != nil {
		return errors.New("user.go::LoginUser()::User not found or credentials invalid")
	}

	// Check if the provided password matches the hashed password in the database
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("user.go::LoginUser()::Invalid credentials: password does not match.")
	}

	// Check if the user is approved
	if approved == 0 {
		return errors.New("user.go::LoginUser()::User is not approved")
	}

	return nil
}