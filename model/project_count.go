package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
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

func ProjectCounterGetLatest(connection Connection) ProjectMoment {
	counters := ProjectMoment{}
	collection := connection.Database.C("ProjectCounterMoment")
	err := collection.Find(bson.M{}).Limit(1).Sort("-$natural").One(&counters)
	if err != nil {
		panic(err)
	}
	return counters
}
