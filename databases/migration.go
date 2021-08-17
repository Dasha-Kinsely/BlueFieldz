package databases

import "github.com/Dasha-Kinsely/leaveswears/models"

func FirstMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}