package utilities

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Overload(".env.dev", ".env")
}

func Env(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
