package model

import (
	"time"
)

type ProjectCount struct {
	Name  string
	Count int
}

type ProjectMoment struct {
	Moment   time.Time
	Projects []ProjectCount
}

func ProjectCounterAdd(connection Connection, moment time.Time, projects map[string]int) {
	projectsCount := []ProjectCount{}
	for name, count := range projects {
		projectsCount = append(projectsCount, ProjectCount{name, count})
	}

	collection := connection.Database.C("ProjectCounterMoment")
	err := collection.Insert(&ProjectMoment{moment, projectsCount})
	if err != nil {
		panic(err)
	}
}

func ProjectCounterGetLatest(connection Connection, dateFrom string, dateTo string) ProjectMoment {
	counters := ProjectMoment{}
	collection := connection.Database.C("ProjectCounterMoment")

	filter := DateTimeFilter(dateFrom, dateTo)
	
	err := collection.Find(filter).Limit(1).Sort("-$natural").One(&counters)

	if err != nil {
		if err.Error() == "not found"{
			return ProjectMoment{}
		}
		panic(err)
	}
	return counters
}
