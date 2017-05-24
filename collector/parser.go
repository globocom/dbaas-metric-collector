package collector

import (
	"time"

	"github.com/otherpirate/dbaas-metric-collector/model"
)

func Extractor(databases []Database) {
	connection := model.GetConnection()
	defer connection.Session.Close()

	teams := make(map[string]int)
	for _, database := range databases {
		teams[database.Team] += 1
	}

	moment := time.Now()
	model.DatabaseCounterAdd(connection, moment, len(databases))
	model.TeamCounterAdd(connection, moment, teams)
}
