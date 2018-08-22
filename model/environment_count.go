package model

import (
	"time"
)

type EnvironmentCount struct {
	Name  string
	Count int
}

type EnvironmentMoment struct {
	Moment       time.Time
	Environments []EnvironmentCount
}

func EnvironmentCounterAdd(connection Connection, moment time.Time, environments map[string]int) {
	environmentsCount := []EnvironmentCount{}
	for name, count := range environments {
		environmentsCount = append(environmentsCount, EnvironmentCount{name, count})
	}

	collection := connection.Database.C("EnvironmentCounterMoment")
	err := collection.Insert(&EnvironmentMoment{moment, environmentsCount})
	if err != nil {
		panic(err)
	}
}

func EnvironmentCounterGetLatest(connection Connection, dateFrom string, dateTo string) EnvironmentMoment {
	counters := EnvironmentMoment{}
	collection := connection.Database.C("EnvironmentCounterMoment")
 
	filter := DateTimeFilter(dateFrom, dateTo)

	err := collection.Find(filter).Limit(1).Sort("-$natural").One(&counters)

	if err != nil {
		if err.Error() == "not found"{
			return EnvironmentMoment{}
		}
		panic(err)
	}
	return counters
}
