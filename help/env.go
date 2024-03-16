package help

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var Version string = "Indev"

const (
	MODE          = "MODE"
	PORT          = "PORT"
	DATABASE      = "DATABASE"
	PUBLIC_FOLDER = "PUBLIC_FOLDER"
)

func init() {
	godotenv.Load()
}

func Env(key, fallback string) string {
	return OR(os.Getenv(key), fallback)
}

func CurrentYear() string {
	return strconv.Itoa(time.Now().Year())
}
