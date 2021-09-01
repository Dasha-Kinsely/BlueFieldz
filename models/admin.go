package models

type Admin struct {
	AdminID string `gorm:column:id;primaryKey`
	Password string `gorm:"column:password"`
}

func FindOneAdmin(id string) {

}