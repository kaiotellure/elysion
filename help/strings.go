package help

import (
	"fmt"
	"strings"
)

func OR(dynamic, fallback string) string {
	if len(strings.ReplaceAll(dynamic, " ", "")) == 0 {
		return fallback
	}
	return dynamic
}

func Tenary[T comparable](condition bool, truthy, falthy T) T {
	if condition {
		return truthy
	}
	return falthy
}

// works like an or operator, return the first non-zero value it finds, if none, returns the zero-value
func NZ[T comparable](vals ...T) T {
	for _, v := range vals {
		if v != *new(T) {
			return v
		}
	}
	return *new(T)
}

func JoinAnd(str, sep, joint, and string) string {
	if len(str) == 0 {
		return ""
	}
	list := make([]string, 0)
	cs := strings.Split(str, sep)
	if len(cs) == 1 {
		return cs[0]
	}
	for _, item := range cs[:len(cs)-1] {
		list = append(list, strings.TrimSpace(item))
	}
	return fmt.Sprintf("%s %s %s", strings.Join(list, joint), and, strings.TrimSpace(cs[len(cs)-1]))
}
