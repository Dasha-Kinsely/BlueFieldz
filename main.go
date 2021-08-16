package main

import (
	"log"

	"github.com/Dasha-Kinsely/leaveswears/server/loggers"

	"github.com/Dasha-Kinsely/leaveswears/server/databases"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)
	log.Printf(">>> Currently in DebugMode...\n")
	// Create SimpleHTTP and Error Loggers
	eS := loggers.SetUp("s")
	if eS != nil {
		log.Printf(">>> Simple Loggers creation unsuccessful...\n")
	}
	eE := loggers.SetUp("e")
	if eE != nil {
		log.Printf(">>> Error Loggers creation unsuccessful...\n")
	}
	// Initialize Database
	db, err := databases.InitDB()
	if err != nil{
		log.Printf(">>> Failed to initialize Databases")
	} else {
		defer db.Close()
	}
}

func main() {
	r = gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	r.Run()
}
