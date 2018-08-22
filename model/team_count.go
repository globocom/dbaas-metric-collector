package model

import (
	"time"
)

type TeamCount struct {
	Name  string
	Count int
}

type TeamMoment struct {
	Moment time.Time
	Teams  []TeamCount
}

func TeamCounterAdd(connection Connection, moment time.Time, teams map[string]int) {
	teamsCount := []TeamCount{}
	for name, count := range teams {
		teamsCount = append(teamsCount, TeamCount{name, count})
	}

	collection := connection.Database.C("TeamCounterMoment")
	err := collection.Insert(&TeamMoment{moment, teamsCount})
	if err != nil {
		panic(err)
	}
}

func TeamCounterGetLatest(connection Connection, dateFrom time.Time, dateTo time.Time) TeamMoment {
	counters := TeamMoment{}
	collection := connection.Database.C("TeamCounterMoment")
	filter := DateTimeFilter(dateFrom, dateTo)
	
	err := collection.Find(filter).Limit(1).Sort("-$natural").One(&counters)

	if err != nil {
		if err.Error() == "not found"{
			return TeamMoment{}
		}
		panic(err)
	}
	
	return counters
}
