package help

import (
	"strings"
)

func CN(args ...string) string {
	return strings.Join(args, " ")
}

func OR(dynamic, fallback string) string {
	if len(strings.ReplaceAll(dynamic, " ", "")) == 0 {
		return fallback
	}
	return dynamic
}
