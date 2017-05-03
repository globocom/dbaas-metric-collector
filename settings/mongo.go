package settings

import (
	"github.com/otherpirate/dbaas-metric-collector/util"
)

var MONGODB_ENDPOINT = util.GetEnv("DBAAS_MONGODB_ENDPOINT", "mongodb://127.0.0.1:27017/dbaas_metric_collector")
