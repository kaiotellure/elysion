package help

import (
	"os"
	"strings"

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
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func CN(args ...string) string {
	return strings.Join(args, " ")
}

func OR(dynamic, fallback string) string {
	if len(strings.ReplaceAll(dynamic, " ", "")) == 0 {
		return fallback
	}
	return dynamic
}
