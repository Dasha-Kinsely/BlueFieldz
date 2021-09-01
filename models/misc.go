package models

import (
	"github.com/Dasha-Kinsely/leaveswears/models/databases"
)

type MiscBase struct {
	Name string `gorm:"column:name;primaryKey"`
	Supplies int32 `gorm:"column:supplies"`
}

type Color struct {
	MiscBase
	PriceImpact int32 `gorm:"column:price_impact"`
}

type Material struct {
	MiscBase
	PriceImpact int32 `gorm:"column:price_impact"`
}

type Inking struct {
	MiscBase
	PriceImpact int32 `gorm:"column:price_impact"`
}
/*func (m *Color) GetAllColor() error {
	db := databases.GetDB()

}*/

func GetOneColor(condition interface{}) (Color, error) {
	var color Color
	db := databases.GetDB()
	err := db.First(&color, condition).Error
	return color, err
}

func SaveOneColor(data interface{}) error {
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}

func DeleteOneColor(condition interface{}) error {
	db := databases.GetDB()
	err := db.Where(condition).Delete(Color{}).Error
	return err
}

/*func GetAllMaterial() error {

}*/

func GetOneMaterial(condition interface{}) (Material, error) {
	var material Material
	db := databases.GetDB()
	err := db.First(&material, condition).Error
	return material, err
}

func SaveOneMaterial(data interface{}) error {
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}

func DeleteOneMaterial(condition interface{}) error {
	db := databases.GetDB()
	err := db.Where(condition).Delete(Material{}).Error
	return err
}
