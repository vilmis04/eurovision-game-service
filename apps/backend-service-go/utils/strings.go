package utils

import "strings"

func Capitalize(value string) string {
	return strings.ToUpper(string(value[0])) + value[1:]
}
