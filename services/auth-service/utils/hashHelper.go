package utils

import "golang.org/x/crypto/bcrypt"

func HashStr(value string) (string, error) {
	b := []byte(value)
	hashedPass, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

func HashCompare(hashedPass string, pass string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPass), []byte(pass),
	)
	return err
}
