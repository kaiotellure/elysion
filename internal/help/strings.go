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

func JoinAnd(str, sep, joint, and string) string {
	if len(str) == 0 {
		return ""
	}
	list := make([]string, 0)
	cs := strings.Split(str, sep)
	for _, item := range cs[:len(cs)-1] {
		list = append(list, strings.TrimSpace(item))
	}
	return fmt.Sprintf("%s %s %s", strings.Join(list, joint), and, strings.TrimSpace(cs[len(cs)-1]))
}
