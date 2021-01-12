package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwd string) (string, error) {
	v, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func VerifyPassword(hashed, passwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwd))
	return err == nil
}
