package collector

import (
	"time"

	"github.com/globocom/dbaas-metric-collector/model"
)

func Extractor(databases []Database) {
	connection := model.GetConnection()
	defer connection.Session.Close()

	teams := make(map[string]int)
	projects := make(map[string]int)
	environments := make(map[string]int)
	engines := make(map[string]int)
	for _, database := range databases {
		teams[database.Team] += 1
		projects[database.Project] += 1
		environments[database.Environment] += 1
		engines[database.Engine] += 1
	}

	moment := time.Now()
	model.DatabaseCounterAdd(connection, moment, len(databases))
	model.TeamCounterAdd(connection, moment, teams)
	model.ProjectCounterAdd(connection, moment, projects)
	model.EnvironmentCounterAdd(connection, moment, environments)
	model.EngineCounterAdd(connection, moment, engines)
}
