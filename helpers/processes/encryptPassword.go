package processes

import (
	"golang.org/x/crypto/bcrypt"
)
func EncryptPassword(password string) string{
	bytePassword := []byte(password)
	hashed, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashed)
}