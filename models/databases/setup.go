package databases

import (
	//"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// This is only necessary for the first initialization on server start up
func InitDB() {
	// Loading .env and parse variables required
	err := godotenv.Load()
	if err != nil {
		// log.Println("Error loading .env file...")
		panic("Error loading .env file...")
	}
	dsn := os.Getenv("MYSQL_DSN")
	disableDatetimePrecision, err := strconv.ParseBool(os.Getenv("MYSQL_DISABLE_DATETIME_PRECISION"))
	dontSupportRenameIndex, err := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_INDEX"))
	dontSupportRenameColumn, err := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_COLUMN"))
	skipInitializeWithVersion, err := strconv.ParseBool(os.Getenv("MYSQL_SKIP_INITIALIZE_WITH_VERSION"))
	// Database connection configuration but returns nothing
	db, dbErr := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  disableDatetimePrecision,
		DontSupportRenameIndex:    dontSupportRenameIndex,
		DontSupportRenameColumn:   dontSupportRenameColumn,
		SkipInitializeWithVersion: skipInitializeWithVersion,
	}), &gorm.Config{})
	if dbErr != nil {
		// log.Println("Error initializing MYSQL_DB...")
		panic("Error initializing MYSQL_DB...")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
