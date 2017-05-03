package settings

import (
	"os"
)

var DBAAS_ENDPOINT = GetEnv("DBAAS_ENDPOINT", "http://127.0.0.1:8000")
var DBAAS_USER = GetEnv("DBAAS_USER", "admin")
var DBAAS_PASSWORD = GetEnv("DBAAS_PASSWORD", "admin_pwd")

func GetEnv(variable string, default_value string) string {
	value := os.Getenv(variable)
	if len(value) == 0 {
		value = default_value
	}
	return value
}
