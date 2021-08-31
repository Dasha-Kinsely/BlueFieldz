package databases

import (
	//"log"
	"os"
	"strconv"

	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// This is only necessary for the first initialization on server start up
func InitDB() {
	processes.LoadEnv()
	dsn := os.Getenv("MYSQL_DSN")
	disableDatetimePrecision, err := strconv.ParseBool(os.Getenv("MYSQL_DISABLE_DATETIME_PRECISION"))
	dontSupportRenameIndex, err := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_INDEX"))
	dontSupportRenameColumn, err := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_COLUMN"))
	skipInitializeWithVersion, err := strconv.ParseBool(os.Getenv("MYSQL_SKIP_INITIALIZE_WITH_VERSION"))
	if err != nil {
		panic(err)
	}
	// Database connection configuration but returns nothing
	db, dbErr := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  disableDatetimePrecision,
		DontSupportRenameIndex:    dontSupportRenameIndex,
		DontSupportRenameColumn:   dontSupportRenameColumn,
		SkipInitializeWithVersion: skipInitializeWithVersion,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if dbErr != nil {
		// log.Println("Error initializing MYSQL_DB...")
		panic("Error initializing MYSQL_DB...")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
