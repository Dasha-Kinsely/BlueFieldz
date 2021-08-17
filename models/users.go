package models

import (
	"github.com/Dasha-Kinsely/leaveswears/databases"
	"gorm.io/gorm"
)

// User Model Section
type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique"`
	Email string `gorm:"column:email;unique"`
	PasswordHash string `gorm:"column:password;notNull"`
	IsAdmin bool `gorm:"column:admin"`
}

// D.A.O. plus Default Triggered Functions for User Model
func SaveOneUser(data interface{}) error {
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneUser(fetchCondition interface{}) (User, error) {
	db := databases.GetDB()
	var user User
	err := db.Where(fetchCondition).First(&user).Error
	return user, err
}
//--------------------------------------------------------

