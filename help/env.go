package help

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	MODE          = "MODE"
	PORT          = "PORT"
	DATABASE      = "DATABASE"
	PUBLIC_FOLDER = "PUBLIC_FOLDER"
	UPLOAD_FOLDER = "UPLOAD_FOLDER"
)

func init() {
	godotenv.Load()
}

func Env(key, fallback string) string {
	return OR(os.Getenv(key), fallback)
}
