package models

import "gorm.io/gorm"

type ColorEnum int
const (
	Black ColorEnum = iota+1
	Grey
	White
	Green
	Blue
	Red
	Orange
	Yellow
	Purple
)
func (color ColorEnum) ToString() string {
	return [...]string{"Black", "Grey", "White", "Green", "Blue", "Red", "Orange", "Yellow", "Purple"}[color-1]
}

type MaterialEnum int
const (
	Leather MaterialEnum=iota+1
	Cotton
	Nylon
	Polyester
	Silk
	Wool
)
func (material MaterialEnum) ToString() string {
	return [...]string{"Leather", "Cotton", "Nylon", "Polyester", "Silk", "Wool"}[material-1]
}

type Profile struct {
	gorm.Model
	BelongsTo User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PreferredColor ColorEnum `gorm:"column:preferred_color"`
	PreferredMaterial MaterialEnum `gorm:"column:preferred_material"`
	ArmLength string `gorm:"column:arm_length"`
	TorsoLength string `gorm:column:torso_length`
	ShoulderSize string `gorm:"column:shoulder_size"`
	WaistSize string `gorm:"column:waist_size"`
	Allergy string `gorm:"column:allergy"`
}

