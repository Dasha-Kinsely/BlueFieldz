package models

type Admin struct {
	ID string `gorm:column:id`
	Password string `gorm:"column:password"`
}