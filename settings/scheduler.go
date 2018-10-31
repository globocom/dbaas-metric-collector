package settings

import (
	"strconv"

	"github.com/globocom/dbaas-metric-collector/util"
)

var LOADING_HOUR, _ = strconv.Atoi(util.GetEnv("LOADING_HOUR", "7"))
var LOADING_MINUTE, _ = strconv.Atoi(util.GetEnv("LOADING_MINUTE", "30"))
var LOADING_SECOND, _ = strconv.Atoi(util.GetEnv("LOADING_SECOND", "00"))
