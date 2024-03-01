package help

import (
	"fmt"
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

func CommaList(str string) string {
	if len(str) == 0 {
		return ""
	}
	res := make([]string, 0)
	cs := strings.Split(str, ",")
	for _, arg := range cs[:len(cs)-1] {
		res = append(res, strings.TrimSpace(arg))
	}
	return fmt.Sprintf("%s & %s", strings.Join(res, ", "), cs[len(cs)-1])
}
