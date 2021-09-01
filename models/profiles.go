package models

import (
	"github.com/Dasha-Kinsely/leaveswears/models/databases"
	"gorm.io/gorm"
)

type GenderEnum int8
const (
	Male GenderEnum=iota+1
	Female
	Other
)
func (gender GenderEnum) ToString() string {
	return [...]string{"Male", "Female", "Other"}[gender-1]
}

type Profile struct {
	gorm.Model
	Name string `gorm:"index"`
	User User `gorm:"foreignKey:Username;references:Name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ShippingAddress string `gorm:shipping_address`
	Gender GenderEnum `gorm:"column:gender"`
	PreferredColor []*Color `gorm:"column:preferred_color"`
	PreferredMaterial []*Material `gorm:"column:preferred_material"`
	Phone string `gorm:"column:phone"`
	ArmLength string `gorm:"column:arm_length"`
	TorsoLength string `gorm:column:torso_length`
	ShoulderSize string `gorm:"column:shoulder_size"`
	WaistSize string `gorm:"column:waist_size"`
	Allergy string `gorm:"column:allergy"`
}

func SaveProfile(data interface{}) error {
	// This is often called during update, and will override any existing records identified by Username
	db := databases.GetDB()
	err := db.Save(data).Error
	return err
}

func GetProfileByUsername(u string) (Profile, error) {
	db := databases.GetDB()
	var profile Profile
	err := db.First(&profile, "name=?", u).Error
	return profile, err
}

func (p *Profile) UpdateProfile(data interface{}) error {
	db := databases.GetDB()
	err := db.Model(&p).Updates(data).Error
	return err
}