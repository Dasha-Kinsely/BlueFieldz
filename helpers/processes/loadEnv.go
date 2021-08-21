package processes

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnv(stepLabel string, expectedString string) string {
	err := godotenv.Load()
	if err != nil {
		buildString := []string{"failed to load env file at ", stepLabel}
		panic(buildString)
	}
	return os.Getenv(expectedString)
}