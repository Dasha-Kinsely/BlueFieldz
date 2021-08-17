package loggers

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func SetUp(fileType string) error {
	var err error
	switch fileType {
	case "e":
		ErrorLogger()
	case "s":
		HTTPLogger()
	default:
		log.Printf("!!! Failed to Initialize Logger !!!\n")
		return err
	}
	return nil
}

func HTTPLogger() {
	fs, _ := os.Create("simple.log")
	gin.DefaultWriter = io.MultiWriter(fs, os.Stdout)
}

func ErrorLogger() {
	fe, _ := os.Create("error.log")
	gin.DefaultErrorWriter = io.MultiWriter(fe, os.Stderr)
}
