package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Price string `gorm:"column:price"`
	Photo *string `gorm:"column:photo"`
	AvailableInking []Inking `gorm:"column:available_ink"`
	AvailableColor []Color `gorm:"column:available_color"`
	AvailableMaterial []Material `gorm:"column:available_material"`
}