package main

import (
	"log"

	"loggers"

	"databases"

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
	databases.SetUp()
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
