package tests

import (
	"gorm.io/gorm"
)

type Ping struct {
	gorm.Model
	Key   uint
	Value string
}
