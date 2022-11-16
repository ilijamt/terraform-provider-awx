package helpers

import "strings"

func TrimString(space, newLine bool, in string) (out string) {
	out = in
	if space {
		out = strings.TrimSpace(out)
	}
	if newLine {
		out = strings.Trim(out, "\n")
	}
	return out
}
