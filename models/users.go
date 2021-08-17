package models

import (
	"gorm.io/gorm"
	/*"golang.org/x/crypto/bcrypt"
	"errors"*/

	"github.com/Dasha-Kinsely/leaveswears/databases"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique_index"`
	Email string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;notNull"`
	IsAdmin bool `gorm:"column:admin"`
}

func SaveOneUser(data interface{}) error {
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}


