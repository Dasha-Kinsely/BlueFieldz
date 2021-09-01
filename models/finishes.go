package models

import (
	"gorm.io/gorm"
)

type Finish struct {
	gorm.Model
	ProductID Product `gorm:"column:product_id;"`
	ChargedPrice string `gorm:"column:charged_price"`
	TailoredFor User `gorm:"column:tailored_for"`
	CurrentStatus string `gorm:"column:current_status"`
	Inking Inking `gorm:"column:inking"`
	Color Color `gorm:"column:color"`
	Material Material `gorm:"column:material"`
}