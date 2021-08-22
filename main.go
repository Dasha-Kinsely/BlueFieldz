package main

import (
	//"log"
	//"github.com/Dasha-Kinsely/leaveswears/helpers/loggers"
	"github.com/Dasha-Kinsely/leaveswears/models/databases"
	"github.com/Dasha-Kinsely/leaveswears/models/databases/migrations"
	"github.com/Dasha-Kinsely/leaveswears/routers"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	/* Create SimpleHTTP and Error Loggers
	eS := loggers.SetUp("s")
	if eS != nil {
		log.Printf(">>> Simple Loggers creation unsuccessful...\n")
	}
	eE := loggers.SetUp("e")
	if eE != nil {
		log.Printf(">>> Error Loggers creation unsuccessful...\n")
	}*/
	// Initialize Database
	databases.InitDB()
	db := databases.GetDB()
	migrations.FirstMigration(db)
}

func main() {
	// gin Defaults
	r = gin.Default()
	routers.InitializeRoutes(r)
	
	r.Run()
}
