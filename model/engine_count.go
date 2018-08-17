package model

import (
	"time"
)

type EngineCount struct {
	Name  string
	Count int
}

type EngineMoment struct {
	Moment  time.Time
	Engines []EngineCount
}

func EngineCounterAdd(connection Connection, moment time.Time, engines map[string]int) {
	enginesCount := []EngineCount{}
	for name, count := range engines {
		enginesCount = append(enginesCount, EngineCount{name, count})
	}

	collection := connection.Database.C("EngineCounterMoment")
	err := collection.Insert(&EngineMoment{moment, enginesCount})
	if err != nil {
		panic(err)
	}
}

func EngineCounterGet(connection Connection, dateFrom string, dateTo string) []EngineMoment {
	counters := []EngineMoment{}
	collection := connection.Database.C("EngineCounterMoment")

	filter := DateTimeFilter(dateFrom, dateTo)

	err := collection.Find(filter).All(&counters)
	if err != nil {
		panic(err)
	}

	return counters
}
