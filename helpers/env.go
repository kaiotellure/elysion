package helpers

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	MODE             = "MODE"
	PORT             = "PORT"
	DATABASE         = "DATABASE"
	PUBLIC_FOLDER    = "PUBLIC_FOLDER"
	GOOGLE_CLIENT_ID = "GOOGLE_CLIENT_ID"
	CANONICAL_HOST   = "CANONICAL_HOST"
)

var DEFAULTS_ENV = map[string]string{
	"DATABASE":      "tmp/main.development.db",
	"PUBLIC_FOLDER": "web/public",
	"PORT":          "3000",
	"MODE":          "development",
}

func init() {
	godotenv.Load()
}

func Env(key string) string {
	return OR(os.Getenv(key), DEFAULTS_ENV[key])
}

func CurrentYear() string {
	return strconv.Itoa(time.Now().Year())
}
