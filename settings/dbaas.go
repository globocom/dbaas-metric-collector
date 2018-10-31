package settings

import (
	"github.com/globocom/dbaas-metric-collector/util"
)

var DBAAS_ENDPOINT = util.GetEnv("DBAAS_ENDPOINT", "http://127.0.0.1:8000")
var DBAAS_USER = util.GetEnv("DBAAS_USER", "admin")
var DBAAS_PASSWORD = util.GetEnv("DBAAS_PASSWORD", "admin_pwd")
