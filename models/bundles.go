package models

import (
	"gorm.io/gorm"
)

type Bundle struct {
	gorm.Model
	Name string `gorm:"column:name;primaryKey"`
	Associated []Product `gorm:"column:name;"` // Needs further implementation
}
