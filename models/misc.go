package models

import (
	"github.com/Dasha-Kinsely/leaveswears/models/databases"
)

type Color struct {
	Name string `gorm:"column:name;primaryKey"`
	Supplies int32 `gorm:"column:supplies"`
}

type Material struct {
	Name string `gorm:"column:name;primaryKey"`
	Supplies int32 `gorm:"column:supplies"`
}

func SaveOneColor(data interface{}) error {
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}

func SaveOneMaterial(data interface{}) error {
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}

