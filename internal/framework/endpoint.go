package framework

import (
	"fmt"
	p "path"
)

func CleanEndpoint(base string) string {
	return p.Clean(base) + "/"
}

func EndpointWithID(base string, id any) string {
	return p.Clean(fmt.Sprintf("%s/%v", base, id)) + "/"
}
