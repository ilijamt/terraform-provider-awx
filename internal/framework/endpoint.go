package framework

import (
	"fmt"
	p "path"
)

// CleanEndpoint returns a cleaned endpoint path with a trailing slash.
func CleanEndpoint(base string) string {
	return p.Clean(base) + "/"
}

// EndpointWithID builds a cleaned endpoint path for a specific resource ID.
func EndpointWithID(base string, id any) string {
	return p.Clean(fmt.Sprintf("%s/%v", base, id)) + "/"
}
