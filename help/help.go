package help

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	PORT          = "PORT"
	DATABASE      = "DATABASE"
	PUBLIC_FOLDER = "PUBLIC_FOLDER"
	UPLOAD_FOLDER = "UPLOAD_FOLDER"
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
