package processes

import (
	"os"

	"github.com/joho/godotenv"
)

// Usage: strconv.ParseInt(processes.LoadEnv("getting env_var for the current step", "ENV_VARIABLE"))
func LoadEnv(stepLabel string, expectedString string) string {
	err := godotenv.Load()
	if err != nil {
		buildString := []string{"failed to load env file at ", stepLabel}
		panic(buildString)
	}
	return os.Getenv(expectedString)
}