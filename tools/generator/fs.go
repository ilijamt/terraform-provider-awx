package generator

import (
	"embed"
)

//go:embed templates
var resources embed.FS

func GetFileAsByte(path string) ([]byte, error) {
	return resources.ReadFile(path)
}

func Fs() embed.FS {
	return resources
}
