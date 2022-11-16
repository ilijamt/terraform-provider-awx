package helpers

import "strings"

func TrimAwxString(in string) (out string) {
	return strings.Trim(strings.TrimSpace(in), "\n")
}
