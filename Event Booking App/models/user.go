package models

import (
	"errors"
	"event-booking-app/db"
	"event-booking-app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPaasword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userID, err := res.LastInsertId()

	u.ID = userID

	return err

}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email= ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
