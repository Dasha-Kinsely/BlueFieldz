package processes

import (
	"golang.org/x/crypto/bcrypt"
)

func DecryptPassword(formPassword string, actualPassword string) error{
	bytePassword := []byte(formPassword)
	byteHashed := []byte(actualPassword)
	return bcrypt.CompareHashAndPassword(byteHashed, bytePassword)
}