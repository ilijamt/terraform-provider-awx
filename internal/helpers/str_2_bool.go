package helpers

import (
	"strconv"
	"strings"
)

func Str2Bool(val string) bool {
	b, err := strconv.ParseBool(strings.ToLower(val))
	if err != nil {
		b = false
	}
	return b
}
