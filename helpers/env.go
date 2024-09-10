package helpers

import (
	"errors"
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

type prop struct {
	Default  string
	Required bool
}

var PROPS = map[string]prop{
	"MODE":             {"development", false},
	"PORT":             {"3000", false},
	"DATABASE":         {"tmp/main.development.db", false},
	"PUBLIC_FOLDER":    {"web/public", false},
	"GOOGLE_CLIENT_ID": {"", true},
	"CANONICAL_HOST":   {"", true},
}

func init() {
	godotenv.Load()
}

func Env(key string) string {
	props, ok := PROPS[key]
	if !ok {
		panic(errors.New("querying non-registered env variable."))
	}
	defined := OR(os.Getenv(key), props.Default)
	if Empty(defined) && props.Required {
		panic(errors.New("required env variable not defined: " + key))
	}
	return defined
}

func CurrentYear() string {
	return strconv.Itoa(time.Now().Year())
}
