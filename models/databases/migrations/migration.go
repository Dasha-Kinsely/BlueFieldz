package migrations

import (
	"github.com/Dasha-Kinsely/leaveswears/models"
	"gorm.io/gorm"
)

func FirstMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}