package utils

import (
	m "smapurv1_api/models"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (hashed string, err error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword := string(hashedByte)

	return hashedPassword, nil
}

func CheckPasswordHash(hashedPassword, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func HidePasswordInJSON() {
	var u m.Users
	u.Password = ""
}
