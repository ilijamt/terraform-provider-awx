package helpers

import (
	"os"
)

func GetFirstSetEnvVar(envVars ...string) string {
	for _, envVar := range envVars {
		if val := os.Getenv(envVar); val != "" {
			return val
		}
	}
	return ""
}
