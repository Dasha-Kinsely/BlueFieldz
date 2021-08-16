package databases

import (
	"log"
	"os"
	"strconv"

	"github.com/Dasha-Kinsely/leaveswears/models/tests/pingTest"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Loading .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file...")
	}
	dsn := os.Getenv("MYSQL_DSN")
	disableDatetimePrecision, err := strconv.ParseBool(os.Getenv("MYSQL_DISABLEDATETIMEPRECISION"))
	dontSupportRenameIndex, err := strconv.ParseBool(os.Getenv("MYSQL_DONTSUPPORTRENAMEINDEX"))
	dontSupportRenameColumn, err := strconv.ParseBool(os.Getenv("MYSQL_DONTSUPPORTRENAMECOLUMN"))
	skipInitializeWithVersion, err := strconv.ParseBool(os.Getenv("MYSQL_SKIPINITIALIZEWITHVERSION"))
	// Open database connection
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  disableDatetimePrecision,
		DontSupportRenameIndex:    dontSupportRenameIndex,
		DontSupportRenameColumn:   dontSupportRenameColumn,
		SkipInitializeWithVersion: skipInitializeWithVersion,
	}), &gorm.Config{})

	if err != nil {
		panic("Error initializing MYSQL_DB...")
	}
	return db
}
