package models

import (
	"github.com/Dasha-Kinsely/leaveswears/models/databases"
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
func SaveOneUser(data interface{}, u string) error {
	db := databases.GetDB()
	if err := db.Save(data).Error; err != nil {
		return err
	}
	// Saving into Database will automatically generate a profile with only gorm.Model and Username
	linkedProfile := Profile{Name:u}
	if err := db.Table("profiles").Create(&linkedProfile).Error; err != nil {
		return err
	}
	return nil
}

func FindOneUser(fetchCondition interface{}) (User, error) {
	db := databases.GetDB()
	var user User
	err := db.First(&user, fetchCondition).Error
	return user, err
}

func (m *User) UpdateOneUser(data interface{}) error {
	db := databases.GetDB()
	err := db.Model(&m).Updates(data).Error
	return err
}
//--------------------------------------------------------

