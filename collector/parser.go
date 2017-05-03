package collector

import (
	"time"
	"fmt"

	"github.com/otherpirate/dbaas-metric-collector/model"
)

func Extractor(databases []Database) {
	connection := model.GetConnection()
	defer connection.Session.Close()
	moment := time.Now()

	model.DatabaseCounterAdd(connection, moment, len(databases))
}