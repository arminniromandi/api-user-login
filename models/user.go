package models

import (
	"errors"
	"go-project/database"
	"go-project/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email , password) VALUES (? , ?)"

	stmt, err := database.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := utils.Hashpassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}

func (u User) ValidateUserLogin() error {

	query := "SELECT user_id , password FROM user WHERE email = ?"

	row := database.Db.QueryRow(query, u.Email)

	var reterviedPass string

	err := row.Scan(&u.ID, &reterviedPass)

	if err != nil {
		return errors.New("Invalid Login")
	}

	checkpass := utils.ComparePasswordHashed(u.Password, reterviedPass)

	if !checkpass {
		return errors.New("Invalid Login")
	}

	return nil

}
