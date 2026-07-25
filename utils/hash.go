package utils

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string) (string, error) {

	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(result), err
}

func ComparePasswordHashed(password string, hashedPass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}
