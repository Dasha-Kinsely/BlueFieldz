package migrations

import (
	"github.com/Dasha-Kinsely/leaveswears/models"
	"gorm.io/gorm"
)

/*func LoadAdmins(){
	a1 := processes.EnvString("loading admin 1 id", "ADMIN_1")
	a1p := processes.EnvString("loading admin 1 pass", "ADMIN_1_PASSWORD")
	db := databases.GetDB()
	buildCondition := []byte("{\"ID\":")
	condition := processes.CustomTypeStruct()
	db.Save()
}*/

func FirstMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})
	//db.AutoMigrate(&models.Profile{})
}