package processes

import (
	"os"

	"github.com/joho/godotenv"
)

// Usage: strconv.ParseInt(processes.EnvString("getting env_var for the current step", "ENV_VARIABLE"))
// Used for loading single strings
func EnvString(stepLabel string, expectedString string) string {
	err := godotenv.Load()
	if err != nil {
		buildString := []string{"failed to load env file at ", stepLabel}
		panic(buildString)
	}
	return os.Getenv(expectedString)
}

func LoadEnv() {
	// Loading .env and parse variables required
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file...")
	}
	return
}