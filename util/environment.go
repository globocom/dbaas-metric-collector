package util

import (
	"os"
)

func GetEnv(variable string, default_value string) string {
	value := os.Getenv(variable)
	if len(value) == 0 {
		value = default_value
	}
	return value
}
